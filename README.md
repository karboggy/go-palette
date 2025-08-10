# Introduction
Generate a color-scheme palette from an image file.

# Build from source
```shell
# install golang (apt, dnf, ...)
sudo dnf install golang

# build go-palette
go build -o $HOME/.local/bin/go-palette
```

# Dev tips
```shell
# generate go.mod
go mod init go-palette

# pull module colorthief-go
go get github.com/cascax/colorthief-go

# run 
go run main.go
```
