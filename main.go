package main

import "github.com/hdget/hdproject/cmd"

//go:generate go run main.go gen
func main() {
	cmd.Execute()
}
