package main

import (
	"fmt"
	"github.com/sinlov/drone-info-tools/template"
	"log"
	"os"
	"time"

	"github.com/bridgewwater/drone-plugin-temple/plugin"
	"github.com/sinlov/drone-info-tools/drone_urfave_cli_v2"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

const (
	// Version of cli
	Version = "v0.1.2"
	Name    = "drone-plugin-temple"
)

// action
// do cli Action before flag.
func action(c *cli.Context) error {

	isDebug := c.Bool("config.debug")

	drone := drone_urfave_cli_v2.UrfaveCliBindDroneInfo(c)

	if isDebug {
		log.Printf("debug: cli version is %s", Version)
		log.Printf("debug: load droneInfo finish at link: %v\n", drone.Build.Link)
	}

	config := plugin.Config{
		Webhook: c.String("config.webhook"),
		Secret:  c.String("config.secret"),
		MsgType: c.String("config.msg_type"),

		Debug: c.Bool("config.debug"),

		TimeoutSecond: c.Uint("config.timeout_second"),
	}

	if isDebug {
		log.Printf("config.timeout_second: %v", config.TimeoutSecond)
	}

	p := plugin.Plugin{
		Name:    Version,
		Version: Version,
		Drone:   drone,
		Config:  config,
	}
	err := p.Exec()

	if err != nil {
		log.Fatalf("err: %v", err)
		return err
	}

	return nil
}

// pluginFlag
// set plugin flag at here
func pluginFlag() []cli.Flag {
	return []cli.Flag{
		// plugin start
		// new flag string template if no use, please replace this
		&cli.StringFlag{
			Name:    "config.new_arg,new_arg",
			Usage:   "",
			EnvVars: []string{"PLUGIN_new_arg"},
		},
		&cli.StringFlag{
			Name:       "config.webhook,webhook",
			Usage:      "webhook for send api",
			HasBeenSet: false,
			EnvVars:    []string{"PLUGIN_WEBHOOK"},
		},
		&cli.StringFlag{
			Name:    "config.msg_type,msg_type",
			Usage:   "message type",
			Value:   "text",
			EnvVars: []string{"PLUGIN_MSG_TYPE"},
		},

		&cli.BoolFlag{
			Name:    "config.debug,debug",
			Usage:   "debug mode",
			Value:   false,
			EnvVars: []string{"PLUGIN_DEBUG"},
		},
		// plugin end
	}
}

// pluginHideFlag
// set plugin hide flag at here
func pluginHideFlag() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    "config.timeout_second,timeout_second",
			Usage:   "do request timeout setting second.",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{"PLUGIN_TIMEOUT_SECOND"},
		},
	}
}

func main() {
	template.RegisterSettings(template.DefaultFunctions)
	app := cli.NewApp()
	app.Version = Version
	app.Name = "Drone Plugin"
	app.Usage = ""
	year := time.Now().Year()
	app.Copyright = fmt.Sprintf("© 2022-%d sinlov", year)
	author := &cli.Author{
		Name:  "sinlov",
		Email: "sinlovgmppt@gmail.com",
	}
	app.Authors = []*cli.Author{
		author,
	}

	app.Action = action
	flags := drone_urfave_cli_v2.UrfaveCliAppendCliFlag(drone_urfave_cli_v2.DroneInfoUrfaveCliFlag(), pluginFlag())
	flags = drone_urfave_cli_v2.UrfaveCliAppendCliFlag(flags, pluginHideFlag())
	app.Flags = flags

	// kubernetes runner patch
	if _, err := os.Stat("/run/drone/env"); err == nil {
		errDotEnv := godotenv.Overload("/run/drone/env")
		if errDotEnv != nil {
			log.Fatalf("load /run/drone/env err: %v", errDotEnv)
		}
	}

	// app run as urfave
	if err := app.Run(os.Args); nil != err {
		log.Println(err)
	}
}
