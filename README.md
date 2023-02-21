# Go HTTP Server

## Setup

1. Clone the repo to your machine and run.
```
git clone git@github.com:yeoffrey/go-http-server.git
cd go-http-server
go run main.go
```
2. Go to [localhost:3333](http://localhost:3333) or [localhost:4444](http://localhost:4444) to see the endpoints.

## Description
This project is a basic HTTP server that I made by following [this](https://www.digitalocean.com/community/tutorials/how-to-make-an-http-server-in-go) Digital Ocean tutorial and then making my own modifications.

## Modifications Made
Made a function `serveHttp()` that makes setting up a new server easy. This can then be passed into a go subroutine in the main function which makes my life way easier :-)

## Future Work
Want to expand this onto my personal server to host a website.