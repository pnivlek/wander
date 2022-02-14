package module

import "github.com/pnivlek/wander/pkg/wander"

type BaseModule struct {
	name string
	LoadFunc    wander.LoadFunc    `json:"-"`
	SaveFunc    wander.SaveFunc    `json:"-"`
	MessageFunc wander.MessageFunc `json:"-"`
	HelpFunc    wander.HelpFunc    `json:"-"`
	StatsFunc   wander.StatsFunc   `json:"-"`
}

func (p *BaseModule) Name() string {
	return p.name
}

func (p *BaseModule) Load(bot *wander.Bot, medium wander.Medium, data []byte) error {
	if p.LoadFunc != nil {
		return p.LoadFunc(bot, medium, data)
	}
	return nil
}

func (p *BaseModule) Save() (data []byte, error error) {
	if p.SaveFunc != nil {
		return p.SaveFunc()
	}
	return nil, nil
}

func (p *BaseModule) Help(bot *wander.Bot, medium wander.Medium, message wander.Message, detailed bool) []string {
	if p.HelpFunc != nil {
		return p.HelpFunc(bot, medium, message, detailed)
	}
	return nil
}

func (p *BaseModule) Message(bot *wander.Bot, medium wander.Medium, message wander.Message) {
	if p.MessageFunc != nil {
		p.MessageFunc(bot, medium, message)
	}
}
