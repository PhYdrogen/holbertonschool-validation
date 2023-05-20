FROM golang:latest

RUN apt-get update && apt-get install -y hugo make git wget zip nodejs npm &&\
  wget https://github.com/gohugoio/hugo/releases/download/v0.111.3/hugo_0.111.3_Linux-64bit.tar.gz &&\
  tar -xf hugo_0.111.3_Linux-64bit.tar.gz &&\
  mv ./hugo /usr/bin/ &&\
  rm hugo_0.111.3_Linux-64bit.tar.gz &&\
  npm install -g markdownlint-cli &&\
  npm install -g markdown-link-check