#!/bin/bash

if which go ; then
  echo "Exiting"
  exit 0
fi

wget –quiet https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz >/dev/null 2>&1 && tar -C /usr/local -xzf go* 
echo "export PATH=$PATH:/usr/local/go/bin" >> $HOME/.bashrc
echo "export GOPATH=/vagrant/" >> $HOME/.bashrc

