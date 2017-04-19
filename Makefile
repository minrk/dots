PREFIX ?= /usr/local

NAME = dots

$(NAME): *.go
	go build

install: dots
	cp $< $(PREFIX)/bin/dots

dist: dist/dots-darwin dist/dots-linux dist/dots-windows.exe

dist/dots-%: dots.go
	env GOOS=$(basename $*) GOARCH=amd64 go build dots.go
	test -d $(@D) || mkdir -p $(@D)
	mv dots$(suffix $@) $@

clean:
	-rm -rf dist
	-rm -f dots

.PHONY: clean dist

