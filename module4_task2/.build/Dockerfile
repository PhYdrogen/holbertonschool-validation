FROM golang:1.15.8-alpine

WORKDIR /root

SHELL ["/bin/sh", "-o", "pipefail", "-c"]


WORKDIR /
#spellcheck disable=DL3018
RUN apk update && apk add --no-cache nodejs=~14.20 hugo=~0.79.1 wget=~1.21 zip=~3.0 npm=~14.20\
  && rm -rf /var/cache/apk/*\
  &&\
  npm install -g npm@9.6.7 &&\
  npm install -g markdownlint-cli@0.34.0 &&\
  npm install -g markdown-link-check@3.11.2

ENV PATH="${PATH}:/usr/local/go/bin"
