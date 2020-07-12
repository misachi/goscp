package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
)

// BuildArgs returns scp commandline flags as a list
func BuildArgs(c *cli.Context) []string {
	return []string{"-i", c.String("key"), c.String("source"), c.String("destination")}
}

// GetFilePath returns a slice of file path if its a combination of
// the server name and file path e.g if path is user@1.1.1.1:/file/path, it
// returns /file/path
func GetFilePath(path string) string {
	if p := strings.Index(path, ":"); p != -1 {
		return path[p+1:]
	}
	return path
}

// CheckFileExists ensures the file to be copied is not already in destination
func CheckFileExists(c *cli.Context) bool {
	path := GetFilePath(c.String("source"))
	i := strings.LastIndex(path, "/")
	newpath := GetFilePath(c.String("destination")) + path[i:]

	_, err := os.Stat(newpath)
	return err == nil
}

// func set_args(c *Context) {

// }

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Configurations file `FILE`",
			},
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "Name of config to use",
			},
			&cli.StringFlag{
				Name:     "source",
				Aliases:  []string{"s"},
				Usage:    "Source server - Where to get the file",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "destination",
				Aliases:  []string{"d"},
				Usage:    "Destination server - Where to copy the file",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "key",
				Aliases:  []string{"i"},
				Usage:    "Key to the source server",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			str := BuildArgs(c)
			if exists := CheckFileExists(c); exists {
				fmt.Fprint(os.Stderr, "Source file already exists in destination")
			} else {
				cmd := exec.Command("scp", str...)
				if err := cmd.Run(); err != nil {
					return fmt.Errorf("%v", err)
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
