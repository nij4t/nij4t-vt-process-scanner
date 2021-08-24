include config.mk

all: vt.exe

vt.exe: 
	${OPTIONS} go build -o bin/$@ ./cmd/main.go$<
	
clean: 
	rm -rf bin/vt.exe

.PHONY: all clean
