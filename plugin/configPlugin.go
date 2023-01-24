package plugin

const (
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
)

type (
	// Config plugin private config
	Config struct {
		Debug         bool
		TimeoutSecond int
		Webhook       string
		Secret        string
		MsgType       string
	}
)
