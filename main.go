package main

import (
	"fmt"
	"github.com/sinlov/drone-info-tools/drone_info"
	"log"
	"net/url"
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

// bindDroneInfo
// Please do not edit unless you understand drone's environment variables.
func bindDroneInfo(c *cli.Context) drone_info.Drone {
	repoHttpUrl := c.String("repo.http.url")
	repoHost := ""
	repoHostName := ""
	parse, err := url.Parse(repoHttpUrl)
	if err == nil {
		repoHost = parse.Host
		repoHostName = parse.Hostname()
	}
	stageStartT := c.Uint64("stage.started")
	stageStartTime := time.Unix(int64(stageStartT), 0).Format(drone_info.DroneTimeFormatDefault)
	stageFinishedT := c.Uint64("stage.finished")
	stageFinishedTime := time.Unix(int64(stageStartT), 0).Format(drone_info.DroneTimeFormatDefault)
	var drone = drone_info.Drone{
		//  repo info
		Repo: drone_info.Repo{
			ShortName: c.String("repo.name"),
			GroupName: c.String("repo.group"),
			FullName:  c.String("repo.full.name"),
			OwnerName: c.String("repo.owner"),
			Scm:       c.String("repo.scm"),
			RemoteURL: c.String("repo.remote.url"),
			HttpUrl:   repoHttpUrl,
			SshUrl:    c.String("repo.ssh.url"),
			Host:      repoHost,
			HostName:  repoHostName,
		},
		//  drone_info.Build
		Build: drone_info.Build{
			WorkSpace:    c.String("build.workspace"),
			Status:       c.String("build.status"),
			Number:       c.Uint64("build.number"),
			Tag:          c.String("build.tag"),
			TargetBranch: c.String("build.target_branch"),
			Link:         c.String("build.link"),
			Event:        c.String("build.event"),
			StartAt:      c.Uint64("build.started"),
			FinishedAt:   c.Uint64("build.finished"),
			PR:           c.String("pull.request"),
			DeployTo:     c.String("deploy.to"),
			FailedStages: c.String("failed.stages"),
			FailedSteps:  c.String("failed.steps"),
		},
		Commit: drone_info.Commit{
			Link:    c.String("commit.link"),
			Branch:  c.String("commit.branch"),
			Message: c.String("commit.message"),
			Sha:     c.String("commit.sha"),
			Ref:     c.String("commit.ref"),
			Author: drone_info.CommitAuthor{
				Username: c.String("commit.author.username"),
				Email:    c.String("commit.author.email"),
				Name:     c.String("commit.author.name"),
				Avatar:   c.String("commit.author.avatar"),
			},
		},
		Stage: drone_info.Stage{
			StartedAt:    stageStartT,
			StartedTime:  stageStartTime,
			FinishedAt:   stageFinishedT,
			FinishedTime: stageFinishedTime,
			Machine:      c.String("stage.machine"),
			Os:           c.String("stage.os"),
			Arch:         c.String("stage.arch"),
			Variant:      c.String("stage.variant"),
			Type:         c.String("stage.type"),
			Kind:         c.String("stage.kind"),
			Name:         c.String("stage.name"),
		},
		DroneSystem: drone_info.DroneSystem{
			Version:  c.String("drone.system.version"),
			Host:     c.String("drone.system.host"),
			HostName: c.String("drone.system.hostname"),
			Proto:    c.String("drone.system.proto"),
		},
	}
	return drone
}

func main() {
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
