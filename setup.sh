#!/bin/bash
go build hold.go
go build jump.go
mv hold /bin/
mv jump /bin/_jump
echo "alias jump='$(_jump what)'" >> ~/.profile
echo "alias jump='$(_jump what)'" >> ~/.bashrc
echo "alias jump='$(_jump what)'" >> ~/.zshrc
echo "Installed. Restart all shells to take affect"
