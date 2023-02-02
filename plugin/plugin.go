package plugin

import (
	"errors"
	"fmt"
	"github.com/bridgewwater/drone-plugin-temple/tools"
	"github.com/sinlov/drone-info-tools/drone_info"
	"log"
	"math/rand"
	"os"
	"time"
)

type (
	// Plugin plugin all config
	Plugin struct {
		Name    string
		Version string
		Drone   drone_info.Drone
		Config  Config
	}
)

func (p *Plugin) CleanResultEnv() error {
	for _, envItem := range cleanResultEnvList {
		err := os.Unsetenv(envItem)
		if err != nil {
			return fmt.Errorf("at FileBrowserPlugin.CleanResultEnv [ %s ], err: %v", envItem, err)
		}
	}
	return nil
}

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

	log.Printf("=> plugin %s version %s", p.Name, p.Version)

	return err
}

// randomStr
// new random string by cnt
func randomStr(cnt uint) string {
	var letters = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	result := make([]byte, cnt)
	keyL := len(letters)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(keyL)]
	}
	return string(result)
}

// randomStr
// new random string by cnt
func randomStrBySed(cnt uint, sed string) string {
	var letters = []byte(sed)
	result := make([]byte, cnt)
	keyL := len(letters)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(keyL)]
	}
	return string(result)
}

func setEnvFromStr(key string, val string) {
	err := os.Setenv(key, val)
	if err != nil {
		log.Fatalf("set env key [%v] string err: %v", key, err)
	}
}
