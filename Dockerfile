FROM golang:1.20-rc-buster

RUN useradd -m -s /bin/bash -u 1000 user

USER user