package plugin_test

import (
	"github.com/bridgewwater/drone-plugin-temple/plugin"
	"github.com/sinlov/drone-info-tools/drone_info"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPlugin(t *testing.T) {
	// mock Plugin
	t.Logf("~> mock Plugin")
	p := plugin.Plugin{
		Name:    mockName,
		Version: mockVersion,
	}
	// do Plugin
	t.Logf("~> do Plugin")
	if envCheck(t) {
		return
	}

	// use env:ENV_DEBUG
	p.Config.Debug = envDebug

	err := p.Exec()
	if nil == err {
		t.Fatal("args [ webhook ] empty error should be catch!")
	}

	p.Config.Webhook = envPluginWebhook

	err = p.Exec()
	if nil == err {
		t.Fatal("args [ msg_type ] empty error should be catch!")
	}

	p.Config.MsgType = "mock" // not support type
	err = p.Exec()
	if nil == err {
		t.Fatal("args [ msg_type ] not support error should be catch!")
	}

	envMsgType := os.Getenv("PLUGIN_MSG_TYPE")

	if envMsgType == "" {
		t.Error("please set env:PLUGIN_MSG_TYPE then test")
	}

	p.Config.MsgType = envMsgType

	p.Drone = *drone_info.MockDroneInfo("success")
	// verify Plugin

	err = p.CleanResultEnv()
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "", os.Getenv(plugin.EnvPluginResultShareHost))
}
