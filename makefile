all: build


build:
	go build

run: build
	./helix_show ./sample_y_5.h
