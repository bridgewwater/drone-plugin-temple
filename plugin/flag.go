package plugin

import (
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/sinlov/drone-info-tools/drone_log"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// IsBuildDebugOpen
// when config or drone build open debug will open debug
func IsBuildDebugOpen(c *cli.Context) bool {
	return c.Bool(NamePluginDebug) || c.Bool(drone_info.NameCliStepsDebug)
}

// BindCliFlag
// check args here
func BindCliFlag(c *cli.Context, cliVersion, cliName string, drone drone_info.Drone) (*Plugin, error) {
	debug := IsBuildDebugOpen(c)

	config := Config{
		Debug:         debug,
		TimeoutSecond: c.Uint(NamePluginTimeOut),

		Webhook: c.String(NameWebHook),
		MsgType: c.String(NameMsgType),
	}

	if config.Debug {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}

	// set default TimeoutSecond
	if config.TimeoutSecond == 0 {
		config.TimeoutSecond = 10
	}

	drone_log.Debugf("args %s: %v", NamePluginTimeOut, config.TimeoutSecond)

	p := Plugin{
		Name:    cliName,
		Version: cliVersion,
		Drone:   drone,
		Config:  config,
	}
	return &p, nil
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
			Name:    NameWebHook,
			Usage:   "webhook for send api",
			EnvVars: []string{EnvWebHook},
		},
		&cli.StringFlag{
			Name:    NameMsgType,
			Usage:   "message type",
			Value:   "text",
			EnvVars: []string{EnvMsgType},
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
			Name:    NamePluginTimeOut,
			Usage:   "do request timeout setting second.",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{EnvPluginTimeOut},
		},
		&cli.BoolFlag{
			Name:    NamePluginDebug,
			Usage:   "debug mode",
			Value:   false,
			EnvVars: []string{drone_info.EnvKeyPluginDebug},
		},
	}
}
