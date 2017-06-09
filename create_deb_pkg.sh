#!/bin/bash
FILE="$1"
ARCH="$2"
VERSION="$3"

CONTROL="build/deb/${FILE}_${VERSION}-${ARCH}/DEBIAN/control"
BINDIR="build/deb/${FILE}_${VERSION}-${ARCH}/usr/local/bin"

mkdir -p build/deb/${FILE}_${VERSION}-${ARCH}/usr/local/bin
mkdir -p build/deb/${FILE}_${VERSION}-${ARCH}/DEBIAN

cp $FILE $BINDIR

echo "Package: kmsdecryptenv" > $CONTROL
echo "Version: 1.0.0-1" >> $CONTROL
echo "Section: base" >> $CONTROL
echo "Priority: optional" >> $CONTROL
echo "Architecture: $ARCH" >> $CONTROL
echo "Maintainer: Anders Wallin <anders.wallin@danielwellington.com>" >> $CONTROL
echo "Description: KMS Decrypt ENV" >> $CONTROL
echo " Decrypts all or specific ENV variables using AWS KMS." >> $CONTROL
echo " It run all calls concurrently so decrypting 100s or keys" >> $CONTROL
echo " usally takes less than a couple of 100 ms." >> $CONTROL

docker run -v $(pwd)/build:/build -w /build ubuntu:16.04 dpkg-deb --build /build/deb/${FILE}_${VERSION}-${ARCH}
mv build/deb/${FILE}_${VERSION}-${ARCH}.deb build/${FILE}_${VERSION}-${ARCH}.deb

rm -rf build/deb