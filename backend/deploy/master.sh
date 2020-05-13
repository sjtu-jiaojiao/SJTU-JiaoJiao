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
info "Input your IP: " -n
read ip
sudo docker swarm init --advertise-addr $ip
warning "Copy and input token in slaves!"
confirm "Press Y if slaves are ready."
if [ $? -eq 1 ]; then
    exit
fi
sudo docker node ls
info "Master node OK!"
