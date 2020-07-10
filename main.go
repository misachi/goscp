package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli"
)


func build_args(c *cli.Context) []string {
	return []string{"-i", c.String("key"), c.String("source"), c.String("destination"),}
}

func get_file_path(path string) string {
	if p := strings.Index(path, ":"); p != -1 {
		return path[p+1:]
	} else {
		return path
	}
}

func check_file_exists(c *cli.Context) bool {
	path := get_file_path(c.String("source"))
	i := strings.LastIndex(path, "/")
	new_path := get_file_path(c.String("destination")) + path[i:]

	_, err := os.Stat(new_path)
	return err == nil
}

// func set_args(c *Context) {

// }

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config",
				Aliases: []string{"c"},
				Usage: "Configurations file `FILE`",
			},
			&cli.StringFlag{
				Name: "name",
				Aliases: []string{"n"},
				Usage: "Name of config to use",
			},
			&cli.StringFlag{
				Name: "source",
				Aliases: []string{"s"},
				Usage: "Source server - Where to get the file",
				Required: true,
			},
			&cli.StringFlag{
				Name: "destination",
				Aliases: []string{"d"},
				Usage: "Destination server - Where to copy the file",
				Required: true,
			},
			&cli.StringFlag{
				Name: "key",
				Aliases: []string{"i"},
				Usage: "Key to the source server",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			str := build_args(c)
			if exists := check_file_exists(c); exists {
				fmt.Fprint(os.Stderr, "Source file already exists in destination")
			} else {
				cmd := exec.Command("scp", str...)
				if err := cmd.Run(); err != nil {
				 	return fmt.Errorf("%v\n", err)
				}
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
