package vk

import (
	"github.com/Andrew1481432/goVkBot/vk/pojo"
)

type Response struct {
	Response interface{}
	Error    *pojo.ResponseError
	Raw      map[string]interface{}
}
