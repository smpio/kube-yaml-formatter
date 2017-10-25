build:
	./build.sh

install: build
	cp kube-yaml-cleaner /usr/local/bin/
