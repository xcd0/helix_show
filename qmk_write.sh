#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0); pwd)
cd $SCRIPT_DIR

build=yj

cd ../qmk_firmware

make helix:$build:avrdude

