package main

import (
	"os"
	"fmt"
	"github.com/codegangsta/cli"
	"os/signal"
	"robotics.neu.edu.tr/ra27-telemetry/ra/container"
	"syscall"
	"time"
	"log"
)

func timeout(interval time.Duration) {
	time.Sleep(interval)
	log.Println("Timeout! Killing the main thread...")
	os.Exit(-1)
}

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	app := cli.NewApp()
	app.Name = "ra27-telemetry"
	app.Usage = "Ra27 Telemetry"
	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"l"},
			Usage:   "",
			Action: func(c *cli.Context) {
				println("action:", "run")
				config := GetConfig(c.String("config-path"))
				cnt := container.SetupContainer(container.Config(config))
				cnt.Start()
				<-done
				go timeout(8 * time.Second)
				cnt.Stop()

			},

			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "config-path",
					Value: "config.json",
					Usage: "Path to the config file",
				},
			},
		},
	}

	app.Run(os.Args)
}
