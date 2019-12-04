# cmsc491BlockchainFinalProject

This project is adapted from the Book: Blockchain by Example by Bellaj Badr; Xun Wu; Richard Horrocks 

Prerequisites:

    -Virtual box available at https://www.virtualbox.org/wiki/Downloads or native Linux (Ubuntu)
    -Git (https://git-scm.com/)
    -Go Language (https://golang.org/dl/)
    -Docker engine available at https://docs.docker.com/install/linux/docker-ce/ubuntu/ (version +17.03)
    -Docker Compose available https://docs.docker.com/compose/install/ ( version +1.8)
    -Curl (https://curl.haxx.se/)
    -Node available at https://nodejs.org/en/ (version 8.x < 9)
    -NPM  (version 5.x)
    -Python 2.7.x and Pip

Once you install Go, you should set the environments variables properly : 
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin

You can refer to their official documentation for further installation instructions.

Download the Hyperledger fabric samples in the same directory: https://github.com/hyperledger/fabric-samples

How to run: 
After all the prerequisites are downloaded
1. docker kill $(docker ps -qa)
2. docker rm $(docker ps -qa)
3. rm -rf crypto-config
4. ./fscn.sh -m generate
5. ./fscn.sh -m up
