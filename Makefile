appname := kmsdecryptenv
version := 1.0.1
debversion := $(version)-2

sources := $(wildcard *.go)

build = GOOS=$(1) GOARCH=$(2) go build -o build/$(appname)
deb = ./create_deb_pkg.sh $(appname) $(2) $(debversion)
tar = cd build && tar -cvzf $(1)_$(2).tar.gz $(appname) && rm $(appname)
zip = cd build && zip $(1)_$(2).zip $(appname) && rm $(appname)

.PHONY: all windows darwin linux clean

all: windows darwin linux

clean:
	rm -rf build/

# Linux
linux: linux_arm linux_arm64 linux_386 linux_amd64

linux_386: $(sources)
	$(call build,linux,386)
	$(call deb,linux,386)
	$(call tar,linux,386)

linux_amd64: $(sources)
	$(call build,linux,amd64)
	$(call deb,linux,amd64)
	$(call tar,linux,amd64)

linux_arm: $(sources)
	$(call build,linux,arm)
	$(call deb,linux,arm)
	$(call tar,linux,arm)

linux_arm64: $(sources)
	$(call build,linux,arm64)
	$(call deb,linux,arm64)
	$(call tar,linux,arm64)

# Darwin
darwin: darwin_amd64 darwin_arm64

darwin_amd64: $(sources)
	$(call build,darwin,amd64)
	$(call tar,darwin,amd64)

darwin_arm64: $(sources)
	$(call build,darwin,arm64)
	$(call tar,darwin,arm64)

# Windows
windows: windows_386 windows_amd64

windows_386: $(sources)
	$(call build,windows,386)
	$(call zip,windows,386)

windows_amd64: $(sources)
	$(call build,windows,amd64)
	$(call zip,windows,amd64)