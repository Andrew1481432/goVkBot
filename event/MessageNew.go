package event

import (
	"github.com/Andrew1481432/goVkBot/vk/pojo"
)

type MessageNew struct {
	PrivateMessage *pojo.PrivateMessage `json:"" map:""`
}

func (m MessageNew) GetName() string {
	return MessageNewEvent
}
