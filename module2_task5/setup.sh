apt-get update && apt-get install -y hugo make git wget
wget https://github.com/gohugoio/hugo/releases/download/v0.111.3/hugo_0.111.3_Linux-64bit.tar.gz
tar -xf hugo_0.111.3_Linux-64bit.tar.gz
mv ./hugo /usr/bin/
wget https://go.dev/dl/go1.20.4.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
cd /app
git clone https://github.com/PhYdrogen/holbertonschool-validation
cd holbertonschool-validation/module1_task4/
make build
mv ./dist ../../
exit