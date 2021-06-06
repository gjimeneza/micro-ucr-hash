# Micro UCR Hash

## Overview
This project consists of a functional description of a hashing algorithm in Go. It has a small CLI for ease of use.

## Requirements
* [Go](https://golang.org/doc/install)
* [Cobra CLI](https://github.com/spf13/cobra#installing)

## Installation
Clone repository with:
```
go get github.com/gjimeneza/micro-ucr-hash
```
or with:
```
git clone https://github.com/gjimeneza/micro-ucr-hash.git
```
### Local installation
```
make
```
Run with:
```
cd bin
./micro-ucr-hash
```
### Global installation
Ensure you updated your path in $HOME/.profile by appending the following:
```
export GOPATH=$(go env GOPATH) 
export PATH=$PATH:$GOPATH/bin
```
Install with:
```
go install
```
Run with:
```
micro-ucr-hash
```
## Usage
For bounty/nonce generation from a given 12 byte payload and target focused on area reduction:

Syntax:
```
./micro-ucr-hash area -p <hex payload> -t <target>
```
Example:
```
./micro-ucr-hash area -p ed18be0f984ae0e2e3128efe -t 10
```

For bounty/nonce generation from a given 12 byte payload and target focused on speed:

Syntax:
```
./micro-ucr-hash speed -p <hex payload> -t <target>
```
Example:
```
./micro-ucr-hash speed -p ed18be0f984ae0e2e3128efe -t 10
```
