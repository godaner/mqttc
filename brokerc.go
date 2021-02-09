package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	initLogger(false)
	cli.HelpFlag = cli.BoolFlag{
		Name:  "help",
		Usage: "show help",
	}
	cli.VersionFlag = cli.BoolFlag{
		Name:  "version",
		Usage: "print the version",
	}
	app := cli.NewApp()
	app.Writer = os.Stdout
	app.ErrWriter = os.Stderr
	app.Name = "brokerc"
	app.HelpName = "brokerc"
	app.Usage = "brokerc is a cross platform publish subscribe client, including mqtt client, amqp client, http client."
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		MQTTPublishCommand,
		MQTTSubscribeCommand,
		AMQPSubscribeCommand,
		AMQPPublishCommand,
		HTTPPublishCommand,
		HTTPSubscribeCommand,
	}
	if err := app.Run(os.Args); err != nil {
		logger.Error(err)
		os.Exit(0)
	}
}
