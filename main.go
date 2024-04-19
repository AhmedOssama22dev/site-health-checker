package main

import (
	"github.com/urfave/cli/v2"
	"fmt"
	"log"
	"os"
)

func main() {
	app :=  &cli.App{
		Name: "Healthchecker",
		Usage: "Website healthchecer tool",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "domain",
				Aliases: []string{"d"},
				Usage: "Domain name",
				Required: true,
			},
			&cli.StringFlag{
				Name: "port",
				Aliases: []string{"p"},
				Usage: "Port number",
				Required: false,
		},
	},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if port == "" {
				port = "80"
			}
			status := check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}
	err := app.Run(os.Args) 
	if err != nil {
		log.Fatal(err);
	}
}