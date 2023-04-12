package wInterface

import . "github.com/kanengo/okexsdkv5/ws/wImpl"

// 请求数据
type WSParam interface {
	EventType() Event
	ToMap() *map[string]string
}
