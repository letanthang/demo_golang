package main

import (
	"fmt"
	"github.com/urfave/cli/v2" // imports as package "cli"
	"log"
	"os"
)

var app = cli.NewApp()

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

func info() {
	app.Name = "Thang"
	app.Usage = "Search Organization, Ticket, User"
	// app.Author = "Jeroenouw"
	app.Version = "1.0.0"
}
func commands() {
	app.Commands = []*cli.Command{
		{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "Show help",
			Action: func(c *cli.Context) error {
				fmt.Println("help info")
				return nil
			},
		},
		{
			Name:    "quit",
			Aliases: []string{"q"},
			Usage:   "Quit the app",
			Action: func(c *cli.Context) error {
				fmt.Println("app quited")
				return nil
			},
		},
		{
			Name:    "abc",
			Aliases: []string{"q"},
			Usage:   "abc the app",
			Action: func(c *cli.Context) error {
				fmt.Println("app abc")
				return nil
			},
		},
	}
}
