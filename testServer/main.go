package main

import "github.com/Atticus08/GoLangFun/GoREST/serveFiles"

func main() {
	serveFiles.RunFileServer(":8080")
}
