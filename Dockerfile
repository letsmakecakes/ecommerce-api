FROM ubuntu:latest
LABEL authors="adwai"

ENTRYPOINT ["top", "-b"]