package wander

import "math/rand"

type appType struct {
	Medium
	Modules map[string]Module
}

type Bot struct {
	// This is the math.Rand type. Do not use it for anything cryptographically secure.
	Rand rand.Rand
	Applications map[string]*appType
}

func (b *Bot) RegisterMedium(m Medium) {
	app := &appType{}
	app.Medium = m
	b.Applications[m.Name()] = app
}

func NewBot() Bot {
	return Bot{}
}
