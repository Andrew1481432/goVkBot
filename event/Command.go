package event

import (
	"github.com/Andrew1481432/goVkBot/vk/pojo"
)

type Command struct {
	Command        string
	Args           []string
	PrivateMessage *pojo.PrivateMessage
}

func (c *Command) GetName() string {
	return CommandEvent
}
