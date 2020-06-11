package main

import (
	go_pkg_test "github.com/konchanxxx/go-pkg-test"
)

const (
	configPath = "configs"
)

func main() {
	configFilePath := configPath + "/" + "sample.toml"
	go_pkg_test.Bind(configFilePath)
}
