package plugin

import (
	"errors"
	"fmt"
	"github.com/bridgewwater/drone-plugin-temple/drone_info"
	"github.com/bridgewwater/drone-plugin-temple/tools"
	"log"
	"os"
)

type (
	// Plugin plugin all config
	Plugin struct {
		Drone  drone_info.Drone
		Config Config
	}
)

func (p *Plugin) Exec() error {
	if p.Config.Debug {
		for _, e := range os.Environ() {
			log.Println(e)
		}
	}

	var err error

	if p.Config.Webhook == "" {
		msg := "missing webhook, please set webhook"
		return errors.New(msg)
	}

	if p.Config.MsgType == "" {
		msg := "missing msg type setting, please set message type"
		return errors.New(msg)
	}

	if !(tools.StrInArr(p.Config.MsgType, supportMsgType)) {
		return fmt.Errorf("msg type only support %v", supportMsgType)
	}

	// set default TimeoutSecond
	if p.Config.TimeoutSecond == 0 {
		p.Config.TimeoutSecond = 10
	}

	log.Printf("dev use Webhook: %v\n", p.Config.Webhook)
	log.Printf("dev use MsgType: %v\n", p.Config.MsgType)

	return err
}
