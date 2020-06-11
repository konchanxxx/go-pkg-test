package go_pkg_test

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type User struct {
	Name string `toml:"name"`
}

func Bind(path string) {
	u := &User{}
	_, err := toml.DecodeFile(path, u)
	if err != nil {
		panic(err)
	}

	fmt.Printf("debug user: %#v\n", u)
}
