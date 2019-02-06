package frontend

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/digitalrebar/provision/backend"
	"github.com/digitalrebar/provision/models"
	"github.com/galthaus/gzip"
	"github.com/gin-gonic/gin"
)

func (fe *Frontend) getEndpoint(c *gin.Context, id string) *backend.RawModel {
	ots := fe.dt.GetObjectTypes()
	found := false
	for _, ot := range ots {
		if ot == "endpoints" {
			found = true
			break
		}
	}
	if !found {
		return nil
	}
	rt := fe.rt(c, "endpoints")
	var res *backend.RawModel
	rt.Do(func(d backend.Stores) {
		if u := rt.Find("endpoints", id); u != nil {
			res = u.(*backend.RawModel)
		}
	})
	return res
}

func (fe *Frontend) getEndpointUrl(c *gin.Context, id string) (string, bool) {
	if res := fe.getEndpoint(c, id); res != nil {
		return res.GetStringField("URL")
	}
	return "", false
}

func (fe *Frontend) forwardRequest(c *gin.Context, id, rest string, newBody interface{}) bool {
	if ep, ok := fe.getEndpointUrl(c, id); ok {
		// parse the url
		turl, _ := url.Parse(ep)

		// create the reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(turl)
		proxy.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // client uses self-signed cert
		}

		// Update the headers to allow for SSL redirection
		req := c.Request
		req.URL.Host = turl.Host
		req.URL.Scheme = turl.Scheme
		if rest != "" {
			req.URL.Path = rest
		}
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Host = turl.Host

		if newBody != nil {
			jstr, jerr := json.Marshal(newBody)
			if jerr != nil {
				err := &models.Error{
					Model:    "endpoints", // XXX: This is not right.
					Key:      id,
					Code:     http.StatusBadRequest,
					Type:     c.Request.Method,
					Messages: []string{jerr.Error()},
				}
				c.AbortWithStatusJSON(err.Code, err)
				return true
			}
			if req.Body != nil {
				req.Body.Close()
			}
			req.Body = ioutil.NopCloser(bytes.NewReader([]byte(jstr)))
			req.ContentLength = int64(len([]byte(jstr)))
		}

		// Remove Writer Headers - let the caller put them in.
		c.Writer.Header().Del("Access-Control-Allow-Credentials")
		c.Writer.Header().Del("Access-Control-Expose-Headers")
		c.Writer.Header().Del("Access-Control-Allow-Origin")

		if gzw, ok := c.Writer.(*gzip.GzipWriter); ok {
			gzw.SetSkipCompression(c)
		}

		// Note that ServeHttp is non blocking and uses a go routine under the hood
		proxy.ServeHTTP(c.Writer, req)
		return true
	}
	return false
}

//
// This returns two bools.  If either is true, the call should just return.
// First bool is forwarded or not.
// Second bool is true if the object is not yours and an error was returned.
//
func (fe *Frontend) processRequestWithForwarding(c *gin.Context, obj interface{}, newBody interface{}) (bool, bool) {
	// Should we proxy this - manager plugin sends with true.
	if v, ok := c.GetQuery("noproxy"); ok && v == "true" {
		return false, false
	}
	mobj, _ := obj.(models.Model)
	if owned, ook := mobj.(models.Owner); ook {
		owner := owned.GetEndpoint()
		if owner == "" {
			return false, false // This is mine, don't forward.
		}
		for _, id := range fe.DrpIds {
			if owner == id {
				return false, false // This is mine, don't forward.
			}
		}
		if fok := fe.forwardRequest(c, owned.GetEndpoint(), "", newBody); fok {
			return true, true // forwarded and not mine
		}
		err := &models.Error{
			Model:    mobj.Prefix(),
			Key:      mobj.Key(),
			Code:     http.StatusNotFound,
			Type:     c.Request.Method,
			Messages: []string{"Not Found"},
		}
		c.AbortWithStatusJSON(err.Code, err)
		return false, true // This is not forward, but is also not mine.
	}
	return false, false // Not forwarded and mine (likely)
}
