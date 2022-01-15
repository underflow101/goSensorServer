#! /bin/bash

# Env Build Script
# ================================================

function check_arch() {
    local aarch=$(getconf LONG_BIT)
    echo "${aarch}"

    if [ "${aarch}" = "64" ]; then
        echo "ok"
        echo "Proceed env setting." 
    else
        echo "Arch is 32-bit."
        echo "Can't install MongoDB 4.x to 32-bit architecture."
        echo "Stop procedure."
        echo "exit"
        exit 0
    fi
}

function install_mongodb() {
    echo "Install MongoDB..."

    mkdir ~/db

    wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add -
    echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.2.list
    sudo apt-get update
    sudo apt-get install -y mongodb-org

    echo "Done."
}

function install() {
    echo "Installation start."

    install_mongodb

    echo "Done."
}

check_arch
install

exit 0