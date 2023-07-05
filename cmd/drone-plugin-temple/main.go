//go:build !test

package main

import (
	_ "embed"
	"fmt"
	"github.com/bridgewwater/drone-plugin-temple"
	"github.com/sinlov/drone-info-tools/pkgJson"
	"github.com/sinlov/drone-info-tools/template"
	"log"
	"os"
	"time"

	"github.com/bridgewwater/drone-plugin-temple/plugin"
	"github.com/sinlov/drone-info-tools/drone_urfave_cli_v2"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

// action
// do cli Action before flag.
func action(c *cli.Context) error {

	isDebug := c.Bool("config.debug")

	drone := drone_urfave_cli_v2.UrfaveCliBindDroneInfo(c)

	cliVersion := pkgJson.GetPackageJsonVersionGoStyle()
	if isDebug {
		log.Printf("debug: cli version is %s", cliVersion)
		log.Printf("debug: load droneInfo finish at link: %v\n", drone.Build.Link)
	}

	p := plugin.BindCliFlag(c, cliVersion, pkgJson.GetPackageJsonName(), drone)
	err := p.Exec()

	if err != nil {
		log.Fatalf("err: %v", err)
		return err
	}

	return nil
}

func main() {
	pkgJson.InitPkgJsonContent(drone_plugin_temple.PackageJson)
	template.RegisterSettings(template.DefaultFunctions)
	name := pkgJson.GetPackageJsonName()
	if name == "" {
		panic("package.json name is empty")
	}
	app := cli.NewApp()
	app.Version = pkgJson.GetPackageJsonVersionGoStyle()
	app.Name = name
	app.Usage = pkgJson.GetPackageJsonDescription()
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
	flags := drone_urfave_cli_v2.UrfaveCliAppendCliFlag(drone_urfave_cli_v2.DroneInfoUrfaveCliFlag(), plugin.CommonFlag())
	flags = drone_urfave_cli_v2.UrfaveCliAppendCliFlag(flags, plugin.Flag())
	flags = drone_urfave_cli_v2.UrfaveCliAppendCliFlag(flags, plugin.HideFlag())
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
