FROM golang:latest

SHELL ["/bin/bash", "--login", "-c"]

RUN apt-get update && apt-get install -y hugo wget zip npm
ENV NVM_DIR /root/.nvm
RUN wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.1/install.sh | bash &&\
  . $NVM_DIR/nvm.sh &&\
  nvm install --lts &&\
  npm install -g npm &&\
  npm install -g markdownlint-cli &&\
  npm install -g markdown-link-check

ENV PATH="${PATH}:/usr/local/go/bin"