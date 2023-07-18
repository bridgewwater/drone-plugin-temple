package plugin

import (
	"fmt"
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/sinlov/drone-info-tools/drone_log"
	"github.com/sinlov/drone-info-tools/drone_urfave_cli_v2/exit_cli"
	tools "github.com/sinlov/drone-info-tools/tools/str_tools"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	EnvWebHook = "PLUGIN_WEBHOOK"
	EnvMsgType = "PLUGIN_MSG_TYPE"
)

// BindCliFlag
// check args here
func BindCliFlag(c *cli.Context, cliVersion, cliName string, drone drone_info.Drone) (*Plugin, error) {
	config := Config{
		Webhook: c.String("config.webhook"),
		Secret:  c.String("config.secret"),
		MsgType: c.String("config.msg_type"),

		Debug: c.Bool("config.debug"),

		TimeoutSecond: c.Uint("config.timeout_second"),
	}

	drone_log.Debugf("args config.timeout_second: %v", config.TimeoutSecond)

	if config.Debug {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}

	if config.Webhook == "" {
		err := fmt.Errorf("missing webhook, please set webhook env: %s", EnvWebHook)
		drone_log.Error(err)
		return nil, exit_cli.Err(err)
	}

	if config.MsgType == "" {
		return nil, exit_cli.Format("missing webhook, please set message type env: %s", EnvMsgType)
	}

	if !(tools.StrInArr(config.MsgType, supportMsgType)) {
		return nil, exit_cli.Format("msg type only support %v", supportMsgType)
	}

	// set default TimeoutSecond
	if config.TimeoutSecond == 0 {
		config.TimeoutSecond = 10
	}

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
			Name:       "config.webhook,webhook",
			Usage:      "webhook for send api",
			HasBeenSet: false,
			EnvVars:    []string{EnvWebHook},
		},
		&cli.StringFlag{
			Name:    "config.msg_type,msg_type",
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
			EnvVars: []string{drone_info.EnvKeyPluginDebug},
		},
	}
}
