package event

import (
	"github.com/Andrew1481432/goVkBot/vk/object"
)

type MessageNew struct {
	PrivateMessage *object.PrivateMessage `json:"" map:""`
}

func (m MessageNew) GetName() string {
	return MessageNewEvent
}
