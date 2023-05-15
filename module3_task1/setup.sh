sudo apt-get update && sudo apt-get install -y hugo make git wget
wget https://github.com/gohugoio/hugo/releases/download/v0.111.3/hugo_0.111.3_Linux-64bit.tar.gz
tar -xf hugo_0.111.3_Linux-64bit.tar.gz
sudo mv ./hugo /usr/bin/
sudo rm hugo_0.111.3_Linux-64bit.tar.gz
sudo npm install -g markdownlint-cli
sudo npm install -g markdown-link-check