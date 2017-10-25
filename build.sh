#!/bin/bash
set -e

GOOS=$(uname | awk '{print tolower($0)}')

arch=$(uname -m)
case $arch in
i686)
    GOARCH=386
    ;;
x86_64)
    GOARCH=amd64
    ;;
*)
    GOARCH=unknown
    ;;
esac


image_name=kube-yaml-cleaner

# Build image with binary
docker build --build-arg "GOOS=$GOOS" --build-arg "GOARCH=$GOARCH" -t $image_name .

# Copy out binary to host
container_id=$(docker create $image_name)
docker cp $container_id:/kube-yaml-cleaner .
docker rm $container_id
docker rmi $image_name
