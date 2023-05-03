docker run --rm --tty --interactive --volume=$(pwd):/app --workdir=/app ubuntu:18.04 /bin/bash
apt-get update && apt-get install -y hugo make && apt install git
git clone https://github.com/PhYdrogen/holbertonschool-validation
cd holbertonschool-validation/
make build
exit