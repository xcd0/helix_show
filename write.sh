#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0); pwd)
build=yj
cd ../qmk_firmware
make helix:$build:avrdude

