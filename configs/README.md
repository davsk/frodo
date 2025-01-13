# `/configs`

Configuration file templates or default configs.

Put your `confd` or `consul-template` template files here.

Installation instructions for consul-template
``` zsh
cd ~/Documents/Workspace/github.com/hashicorp

brew tap hashicorp/tap
brew install hashicorp/tap/nomad
brew install hashicorp/tap/vault

# Install Consul
git clone https://github.com/hashicorp/consul.git

cd consul
  make dev
cd -

consul version

# Install consul-template
git clone https://github.com/hashicorp/consul-template.git

cd consul-template
  make dev
  make test
cd -  
```
To start hashicorp/tap/nomad now and restart at login:
  brew services start hashicorp/tap/nomad
Or, if you don't want/need a background service you can just run:
  /home/linuxbrew/.linuxbrew/opt/nomad/bin/nomad agent -dev
