include config.mk

SRCS := $(wildcard cmd/**/*.go)
BINS := $(SRCS:cmd/%/main.go=bin/%.exe)

all: ${BINS}

bin/%.exe: cmd/%/main.go
	${OPTIONS} go build -o $@ $<
	
clean: 
	rm -rf ${BINS}

.PHONY: all clean
