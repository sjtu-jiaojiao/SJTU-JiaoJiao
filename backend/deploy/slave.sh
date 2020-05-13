#!/bin/bash
color_green="\033[32m"
color_red="\033[31m\033[01m"
color_yellow="\033[33m\033[01m"
color_end="\033[0m"

function error() {
    echo -e $2 "${color_red}[ ERROR ] $1${color_end}"
}

function warning() {
    echo -e $2 "${color_yellow}[ WARNING ] $1${color_end}"
}

function info() {
    echo -e $2 "${color_green}[ INFO ]${color_end} $1"
}

function confirm() {
    while true; do
        echo -n "$1"
        read -n1 tmp
        echo
        if [ "$tmp" = "Y" ]; then
            return 0
        elif [ "$tmp" = "N" ]; then
            return 1
        fi
    done
}

info "Updating system..."
sudo yum update -y

info "Installing docker..."
sudo yum install -y yum-utils
sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
sudo yum install docker-ce docker-ce-cli containerd.io -y
sudo systemctl enable docker
sudo systemctl start docker

info "Initializing swarm..."
info "Input master token: " -n
read token
info "Input master ip/port: " -n
read ipport
sudo docker swarm join --token $token $ipport

info "Install consul?(Y/N) " -n
confirm
if [ $? -eq 0 ]; then
    VER="1.7.1"
    info "Installing consul..."
    sudo yum install -y wget unzip
    wget https://releases.hashicorp.com/consul/${VER}/consul_${VER}_linux_amd64.zip
    unzip consul_${VER}_linux_amd64.zip
    rm consul_${VER}_linux_amd64.zip
fi
info "Slave node OK!"
