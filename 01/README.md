# Lesson one


# install

```
https://go.dev/doc/install
```

```
wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go 
sudo tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz

```

## add in bashrc

```
gedit ~/.bashrc

export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$(go env GOPATH)/bin
export GOPATH=$(go env GOPATH)

source ~/.bashrc
```


```
go mod init
go mod tidy
go run main.go

```

 Every Go program is made up of packages. 

 By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand. 


 Build 

```
go build -o build/app
./build/app
```