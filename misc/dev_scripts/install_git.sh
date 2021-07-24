#! /bin/sh

# see https://git-scm.com/book/en/v2/Getting-Started-Installing-Git

yum -y install dh-autoreconf curl-devel expat-devel gettext-devel \
  openssl-devel perl-devel zlib-devel

VERSION=2.8.0
PACKAGE=git-$VERSION.tar.gz
wget https://mirrors.edge.kernel.org/pub/software/scm/git/$PACKAGE
tar -xzvf $PACKAGE
cd git-$VERSION
make configure
./configure --prefix=/usr
make && make install