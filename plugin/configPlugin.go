package plugin

const (
	EnvPluginResultShareHost = "PLUGIN_RESULT_SHARE_HOST"

	msgTypeText        = "text"
	msgTypePost        = "post"
	msgTypeInteractive = "interactive"
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
		Webhook string
		Secret  string
		MsgType string

		Debug bool

		TimeoutSecond uint
	}
)
