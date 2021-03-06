# Purpose
This repository is meant to stimulate ideas on how to store public data records on Hyperledger Sawtooth. 
It's a slightly modified copy of xo_go transaction processor, part of github.com/hyperledger/sawtooth-sdk-go

![alt text](docs/Mdata.png "Diagram")

# System Requirements

1. OS Packages
    ```
    sudo apt install zip curl python3 python3-pip
    ```

2. Install protobuf compilers (make sure to get a 3.x.x version)
    ```
    curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-linux-x86_64.zip
    unzip protoc-3.7.1-linux-x86_64.zip -d protoc3
    sudo mv protoc3/bin/* /usr/local/bin/
    sudo mv protoc3/include/* /usr/local/include/
    ```

2. Install python's grpcio-tools library
    ```
    sudo su - 
    python3 -m pip install grpcio-tools
    ```

3. Install go dependencies
    ```
    go get -u google.golang.org/grpc && \
    go get -u github.com/golang/protobuf/protoc-gen-go && \
    go get github.com/satori/go.uuid && \
    go get github.com/pebbe/zmq4 && \
    go get github.com/golang/mock/gomock && \
    go install github.com/golang/mock/mockgen && \
    go get github.com/jessevdk/go-flags && \
    go get github.com/stretchr/testify/mock
    ```

# Usage
