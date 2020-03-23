package event

import (
	"github.com/Andrew1481432/goVkBot/vk/pojo"
)

type MessageEdit struct {
	PrivateMessage *pojo.PrivateMessage `json:"" map:""`
}

func (m MessageEdit) GetName() string {
	return MessageEditEvent
}
