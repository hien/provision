#!/bin/sh
set -e
(cd docker && docker build -t ipxebuild .)
[ -d ipxe ] || git clone git://git.ipxe.org/ipxe.git
# Work around a bug in building host utils in ipxe
sed -i -e '/^HOST_CFLAGS/ s/\$(WORKAROUND_CFLAGS)//' ipxe/src/Makefile.housekeeping
if [[ ! -f $1 ]]; then
    echo "First argument must be the path to an .ipxe script"
    echo "which will be embedded into the generated ISO image"
    echo "as the default script to run."
    exit 1
fi
docker run \
       -v $PWD/ipxe:/src/ipxe:rw \
       -v "$1:/src/embed.ipxe:ro" \
       ipxebuild /bin/build_ipxe.sh "bios_iso"
echo "Finished ISO is at ipxe/bin/ipxe.iso"
