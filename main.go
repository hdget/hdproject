package main

import "hdproject/cmd"

//go:generate go run main.go gen
func main() {
	cmd.Execute()
}
