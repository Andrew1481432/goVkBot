package event

import (
	"github.com/Andrew1481432/goVkBot/vk/object"
)

type MessageEdit struct {
	PrivateMessage *object.PrivateMessage `json:"" map:""`
}

func (m MessageEdit) GetName() string {
	return MessageEditEvent
}
