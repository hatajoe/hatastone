package match

import (
	"github.com/hatajoe/hatastone/match/event"
	"github.com/hatajoe/hatastone/match/rule"
	player "github.com/hatajoe/hatastone/player/context/match"
)

type Context struct {
	rule        rule.IRule // match rule
	firstPlayer player.IPlayer
	firstEvent  event.Events // first player event listeners
	afterPlayer player.IPlayer
	afterEvent  event.Events // after player event listeners

	currentPlayer player.IPlayer
	currentEvents event.Events
	state         IState // game state
}

func NewContext(rule rule.IRule, firstPlayer, afterPlayer player.IPlayer, firstEvent, afterEvent event.Events) *Context {
	return &Context{
		rule:          rule,
		firstPlayer:   firstPlayer,
		firstEvent:    firstEvent,
		afterPlayer:   afterPlayer,
		afterEvent:    afterEvent,
		currentPlayer: firstPlayer,
		currentEvents: firstEvent,
		state:         &InCoinToss{},
	}
}

func (c *Context) Exec() error {
	return c.state.Exec(c)
}

func (c *Context) SetState(s IState) {
	c.state = s
}

func (c Context) GetState() IState {
	return c.state
}

func (c Context) GetOpponent() player.IPlayer {
	if c.firstPlayer.GetID() != c.currentPlayer.GetID() {
		return c.firstPlayer
	}
	return c.afterPlayer
}

func (c Context) GetSeed() int64 {
	return c.rule.GetSeed()
}

func (c Context) GetFirstEvent() event.Events {
	return c.firstEvent
}

func (c Context) GetAfterEvent() event.Events {
	return c.afterEvent
}

func (c Context) GetCurrentEvents() event.Events {
	return c.currentEvents
}

func (c *Context) SwitchCurrentPlayer() player.IPlayer {
	if c.currentPlayer.GetID() == c.firstPlayer.GetID() {
		c.currentPlayer = c.afterPlayer
		c.currentEvents = c.afterEvent
		return c.currentPlayer
	}
	c.currentPlayer = c.firstPlayer
	c.currentEvents = c.firstEvent
	return c.currentPlayer
}

func (c Context) IsHeroDead() bool {
	return c.firstPlayer.IsHeroDead() || c.afterPlayer.IsHeroDead()
}

func (c *Context) SwapPlayOrder() {
	c.firstEvent, c.afterEvent = c.afterEvent, c.firstEvent
}
