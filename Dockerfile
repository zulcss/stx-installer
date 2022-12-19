FROM debian:testing

RUN apt-get update -y && \
    apt-get install -y parted systemd udev vim \
    golang build-essential ca-certificates fdisk
RUN mkdir -p /oem/images /etc/stx/
ADD config /etc/stx
