package module

import (
	"math"

	"github.com/pnivlek/wander/common"
)

type EightBallModule struct {
	common.BasePlugin
	responses []string
}

func (eb *EightBallModule) Name() string {
	return "8ball"
}

func (eb *EightBallModule) Message(b *common.Bot, m common.Medium, msg common.Message) {
	response := math.Floor(b.Rand.Float64() * float64(len(eb.responses)))
	err := m.SendMessage(msg.Channel(), eb.responses[int(response)])
	if err != nil {
		m.Log(m, common.LogLevelError, "could not send message")
	}
}
