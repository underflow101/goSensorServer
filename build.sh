#! /bin/bash
set -e

# NOTE
# build script
#		for linux            	: ./build.sh linux
#		for Raspberry Pi 64-bit : ./build.sh
# ==========================================

export SERVER_NAME="goSensorServer"
export SERVER_VERSION="v0.1.0"
export CURRENT_DIR=$(pwd)
export SRC_DIR=$(pwd)
export GOMOD="$(pwd)/go.mod"
export GO111MODULE=on
export GOPROXY="https://proxy.golang.org,direct"

echo $GOMOD
echo $GO111MODULE
echo $GOPROXY

echo $(pwd)
echo $CURRENT_DIR

make clean

if [ "$1" = "linux" ]; then
    echo "====================================="
	echo "Target Binary arch is Linux"
	echo "====================================="
    export CC="gcc"
else
    echo "====================================="
	echo "Target Binary arch is ARM 64-bit"
	echo "====================================="
    export GOARCH=arm64 GOOS=linux
fi

echo $GOOS
function build_server()
{
    make clean
    make build
}

function arrange()
{
    mv bin ${SRC_DIR}/ 
    rm -rf ${TEMP_DIR}      
}
       
build_server