#!/bin/bash

SCRIPT_DIR=$(cd $(dirname $0); pwd)
cd $SCRIPT_DIR

dir=`find . -maxdepth 1 -type d -not -name '.*'`
#echo $dir

for i in $dir; do
	tmp=${i#./}
	rm -rf ../qmk_firmware/keyboards/helix/rev2/keymaps/$tmp*
	cp -rf  $SCRIPT_DIR/$tmp ../qmk_firmware/keyboards/helix/rev2/keymaps/$tmp
done

build=yj

cd ../qmk_firmware

make helix:$build

cp *.hex ../helix_show

make helix:$build:avrdude
make helix:$build:avrdude

