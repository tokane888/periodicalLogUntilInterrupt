#!/bin/bash

set -eux

deb_build() {
  local ver=$1

  cp ../main.go periodical-log-$ver/
  cd periodical-log-$ver
  dh_make -s --email deb_test@gmail.com --createorig -y
  # debuild -us -uc
  cd ..
}

deb_build 1.0
deb_build 1.1
