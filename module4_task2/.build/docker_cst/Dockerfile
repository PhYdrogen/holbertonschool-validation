FROM ubuntu:18.04

WORKDIR /root

ADD container-structure-test-linux-amd64 /usr/local/bin/container-structure-test
ADD cst.yml /root/cst.yml

RUN chmod u+x /usr/local/bin/container-structure-test
CMD container-structure-test test --image awesome:build --config cst.yml