# gloco
Convert Docker Compose files to systemd units.
# Why?
Because I got tired of rewriting them.
# Building and installation
Grab the files:
```shell
git clone https://github.com/nevermarine/gloco
```
Move to project directory:
```shell
cd gloco
```
Build:
```shell
make
```
Or build directly with Go:
```shell
go build .
```
To install, just copy file for your operating system to any directory in your PATH.
```shell
sudo cp ./gloco-linux /usr/bin/gloco
```
# Usage
```shell
$ gloco -h                           
Usage of ./gloco:
  -c string
    	Path to Compose file (default "docker-compose.yml")
  -i string
    	Path to resulting service file (default "result.service")
```
If you already have docker-compose.yml in the working directory (the repo has one), you can just run it.
```shell
$ gloco
```
The default output file is `result.service` as indicated by help string.
