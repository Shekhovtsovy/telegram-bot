FROM golang:1.16-alpine

ENV PACKAGES="bash \
    bash-completion \
    git \
    nano \
    vim \
    curl \
    net-tools \
    coreutils \
    sudo \
    build-base \
    openssh \
    "

RUN apk add ${PACKAGES}

ENTRYPOINT [ "tail", "-f", "/dev/null" ]
