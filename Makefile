default:
	$(MAKE) deps
	$(MAKE) all
deps:
	go get -v github.com/Masterminds/glide
	glide install
test:
	$(MAKE) deps
	go test ./netstat -v
check:
	$(MAKE) test
all:
	bash -c "./scripts/build.sh $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))"
