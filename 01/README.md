# Lesson one



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