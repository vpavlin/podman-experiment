build-runtime:
	cd container-runtime &&\
	podman build -t container-runtime .

build-transformer:
	cd transformer &&\
	podman build -t transformer .

build: build-runtime build-transformer

run:
	go run -tags "exclude_graphdriver_devicemapper exclude_graphdriver_btrfs" main.go