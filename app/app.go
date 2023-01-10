package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line app"
	app.Usage = "Fetch IPs and servers from the internet"
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",          // default
			Value: "google.com.br", // default
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Fetch IPs",
			Flags:  flags,
			Action: searchIp,
		},
		{
			Name:   "servers",
			Usage:  "Fetch servers",
			Flags:  flags,
			Action: searchServer,
		},
	}

	return app
}

func searchIp(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(c *cli.Context) {
	host := c.String("host")
	servers, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
