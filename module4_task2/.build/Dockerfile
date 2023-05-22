FROM golang:1.15.8-buster

SHELL ["/bin/bash", "--login", "-c"]

RUN apt-get update && apt-get install -y --no-install-recommends hugo wget zip npm\
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash &&\
  . /root/.nvm/nvm.sh &&\
  nvm install --lts &&\
  npm install -g npm@9.6.7 &&\
  npm install -g markdownlint-cli@0.34.0 &&\
  npm install -g markdown-link-check@3.11.2

ENV PATH="${PATH}:/usr/local/go/bin"