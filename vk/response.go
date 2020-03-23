package vk

import (
	"github.com/Andrew1481432/goVkBot/vk/object"
)

type Response struct {
	Response interface{}
	Error    *object.ResponseError
	Raw      map[string]interface{}
}
