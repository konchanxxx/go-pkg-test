package go_pkg_test

import (
	"fmt"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type User struct {
	Name string `toml:"name"`
}

func Bind(path string) {
	u := &User{}
	p, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	_, err = toml.DecodeFile(p, u)
	if err != nil {
		panic(err)
	}

	fmt.Printf("debug user: %#v\n", u)
}
