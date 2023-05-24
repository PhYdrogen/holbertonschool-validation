FROM golang:1.15.8-alpine

SHELL ["/bin/bash", "-o", "pipefail", "-c"]
#shellcheck disable=SC1091
#shellcheck source=/root/install.sh

COPY --chmod=751 install.sh /root/install.sh

WORKDIR /root

RUN install.sh | bash && . /root/.nvm/nvm.sh

WORKDIR /

RUN apt-get update && apt-get install -y --no-install-recommends hugo=0.54.0 wget=1.20.3 zip=3.0 npm=9.6.4\
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*\
  &&\
  nvm install --lts &&\
  npm install -g npm@9.6.7 &&\
  npm install -g markdownlint-cli@0.34.0 &&\
  npm install -g markdown-link-check@3.11.2

ENV PATH="${PATH}:/usr/local/go/bin"
