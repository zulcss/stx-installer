FROM debian:testing

RUN apt-get update -y && \
    apt-get install -y parted systemd udev vim \
    golang cobra-cli build-essential ca-certificates


