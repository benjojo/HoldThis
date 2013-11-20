#!/bin/bash
go build hold.go
go build jump.go
mv hold /bin/
mv jump /bin/_jump
echo 'function jump2() { $( _jump "$@" ) ;}' >> ~/.profile
echo 'function jump2() { $( _jump "$@" ) ;}' >> ~/.bashrc
echo 'function jump2() { $( _jump "$@" ) ;}' >> ~/.zshrc
echo "Installed. Restart all shells to take affect"
