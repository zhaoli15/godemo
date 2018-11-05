.PHONY: check build clean test all

check:
	[[ -n "${GOROOT}" ]] || (echo "GOROOT is not defined!" && exit 1)
	[[ -n "${GOPATH}" ]] || (echo "GOPATH is not defined!" && exit 1)

build:
	@echo "building via go install"
	go install -x

clean:
	@echo "cleaning bin dir"
	rm -rf ${GOPATH}/bin/godemo

all: check clean build 
