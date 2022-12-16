cwd  := $(shell pwd)

all: clean build

image:
	docker image rm -f stx-installer
	docker build -t stx-installer .
	docker run -i -t -v $(cwd):/work --privileged -v /dev:/dev --workdir /work stx-installer

build:
	go build -o stx-installer main.go

clean:
	rm -f stx-installer
