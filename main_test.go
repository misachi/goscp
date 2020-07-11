package main_test

import (
	"flag"
	"testing"
	. "goscp"

	"github.com/urfave/cli"
)

func getContext() *cli.Context {
	set := flag.NewFlagSet("test", 2)
	set.String("key", "/key", "doc")
	set.String("source", "/src", "doc")
	set.String("destination", "/dest", "doc")

	return cli.NewContext(nil, set, nil)
}

func TestBuildArgs(t *testing.T) {
	c := getContext()

	ret := BuildArgs(c)

	if ret[1] != c.String("key") {
		t.Errorf("file error in key: %s", c.String("key"))
	}
	if ret[2] != c.String("source") {
		t.Errorf("file error in source: %s", c.String("source"))
	}
	if ret[3] != c.String("destination") {
		t.Errorf("file error in destination: %s", c.String("destination"))
	}
}

func TestGetFilePath(t *testing.T) {
	c := getContext()

	ret := GetFilePath(c.String("source"))
	if ret != c.String("source") {
		t.Errorf("Error while getting source file path: %s", c.String("source"))
	}

	c.Set("source", "user@1.1.1.1:/remote/server")
	ret2 := GetFilePath(c.String("source"))
	if ret2 != "/remote/server" {
		t.Errorf("Error while parsing source file path: %s", c.String("source"))
	}
}
