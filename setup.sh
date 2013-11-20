#!/bin/bash
go build hold.go
go build jump.go
mv hold /bin/
mv jump /bin/_jump
echo 'function jump() { $( _jump "$@" ) ;}' >> ~/.profile
echo 'function jump() { $( _jump "$@" ) ;}' >> ~/.bashrc
echo 'function jump() { $( _jump "$@" ) ;}' >> ~/.zshrc
echo "Installed. Restart all shells to take affect"
