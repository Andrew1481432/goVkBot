package event

import (
	"github.com/Andrew1481432/goVkBot/vk/pojo"
)

type MessageReply struct {
	Message *pojo.Message `json:"" map:""`
}

func (m MessageReply) GetName() string {
	return MessageReplyEvent
}
