DST=./release
BIN=helix_show
GOARCH=amd64

all: build

build:
	go build

run: build
	./helix_show ./yj5/yj5.json
	mv ./yj5/yj5.h ./yj5/5.h

qmk:
	bash qmk_build.sh

release:
	rm -rf $(DST) && mkdir -p $(DST)
	# for windows
	GOARCH=$(GOARCH) GOOS=windows go build -o $(DST)/$(BIN)_windows.exe $(FLAGS_WIN)
	cd $(DST) && mv $(BIN)_windows.exe $(BIN).exe && zip $(BIN)_binary_$(GOARCH)_windows.zip $(BIN).exe && mv $(BIN).exe $(BIN)_windows.exe && rm $(BIN)_windows.exe
	# for mac
	GOARCH=$(GOARCH) GOOS=darwin go build -o $(DST)/$(BIN)_macOS $(FLAGS)
	cd $(DST) && mv $(BIN)_macOS $(BIN) && zip $(BIN)_binary_$(GOARCH)_macOS.zip $(BIN) && mv $(BIN) $(BIN)_macOS && rm $(BIN)_macOS
	# for linux
	GOARCH=$(GOARCH) GOOS=linux go build -o $(DST)/$(BIN)_linux $(FLAGS)
	cd $(DST) && mv $(BIN)_linux $(BIN) && zip $(BIN)_binary_$(GOARCH)_linux.zip $(BIN) && mv $(BIN) $(BIN)_linux && rm $(BIN)_linux
	# for freeBSD
	GOARCH=$(GOARCH) GOOS=freebsd go build -o $(DST)/$(BIN)_freeBSD $(FLAGS)
	cd $(DST) && mv $(BIN)_freeBSD $(BIN) && zip $(BIN)_binary_$(GOARCH)_freeBSD.zip $(BIN) && mv $(BIN) $(BIN)_freeBSD && rm $(BIN)_freeBSD
