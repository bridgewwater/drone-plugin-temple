package plugin

const (
	EnvPluginResultShareHost = "PLUGIN_RESULT_SHARE_HOST"

	NamePluginDebug   = "config.debug"
	EnvPluginTimeOut  = "PLUGIN_TIMEOUT_SECOND"
	NamePluginTimeOut = "config.timeout_second"

	msgTypeText        = "text"
	msgTypePost        = "post"
	msgTypeInteractive = "interactive"

	EnvWebHook  = "PLUGIN_WEBHOOK"
	NameWebHook = "config.webhook"
	EnvMsgType  = "PLUGIN_MSG_TYPE"
	NameMsgType = "config.msg_type"
)

var (
	// supportMsgType
	supportMsgType = []string{
		msgTypeText,
		msgTypePost,
		msgTypeInteractive,
	}

	cleanResultEnvList = []string{
		EnvPluginResultShareHost,
	}
)

type (
	// Config plugin private config
	Config struct {
		Debug         bool
		TimeoutSecond uint

		Webhook string
		MsgType string
	}
)
