package event

import (
	"github.com/Andrew1481432/goVkBot/vk/object"
)

type MessageReply struct {
	Message *object.Message `json:"" map:""`
}

func (m MessageReply) GetName() string {
	return MessageReplyEvent
}
