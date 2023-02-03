package plugin

import (
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/urfave/cli/v2"
	"log"
)

func BindCliFlag(c *cli.Context, cliVersion, cliName string, drone drone_info.Drone) Plugin {
	config := Config{
		Webhook: c.String("config.webhook"),
		Secret:  c.String("config.secret"),
		MsgType: c.String("config.msg_type"),

		Debug: c.Bool("config.debug"),

		TimeoutSecond: c.Uint("config.timeout_second"),
	}

	if config.Debug {
		log.Printf("config.timeout_second: %v", config.TimeoutSecond)
	}

	p := Plugin{
		Name:    cliName,
		Version: cliVersion,
		Drone:   drone,
		Config:  config,
	}
	return p
}

// Flag
// set flag at here
func Flag() []cli.Flag {
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
		// plugin end
		//&cli.StringFlag{
		//	Name:    "config.new_arg,new_arg",
		//	Usage:   "",
		//	EnvVars: []string{"PLUGIN_new_arg"},
		//},
		// file_browser_plugin end
	}
}

// HideFlag
// set plugin hide flag at here
func HideFlag() []cli.Flag {
	return []cli.Flag{
		//&cli.UintFlag{
		//	Name:    "config.timeout_second,timeout_second",
		//	Usage:   "do request timeout setting second.",
		//	Hidden:  true,
		//	Value:   10,
		//	EnvVars: []string{"PLUGIN_TIMEOUT_SECOND"},
		//},
	}
}

// CommonFlag
// Other modules also have flags
func CommonFlag() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    "config.timeout_second,timeout_second",
			Usage:   "do request timeout setting second.",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{"PLUGIN_TIMEOUT_SECOND"},
		},
		&cli.BoolFlag{
			Name:    "config.debug,debug",
			Usage:   "debug mode",
			Value:   false,
			EnvVars: []string{"PLUGIN_DEBUG"},
		},
	}
}
