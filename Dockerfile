FROM debian:testing

RUN apt-get update -y && \
    apt-get install -y parted systemd udev vim \
    golang build-essential ca-certificates git
RUN mkdir -p /oem/images
