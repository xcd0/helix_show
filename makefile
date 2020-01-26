all: build


build:
	go build

run: build
	./helix_show ./yj3/yj3.json
	mv ./yj3/yj3.h ./yj3/5.h

qmk:
	bash qmk_build.sh

