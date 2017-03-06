package dht
import (
	"time"
)

type NodeContext struct {

}

// 维护路由信息
type Route struct {
	bucket []*RemoteNode
}

// 只负责具体消息的发送和接收
type Transport struct {

}

type KRPC struct {
	NodeInfo
	tid     uint32
}

type NodeManager struct {
	route *Route
	msgRecvC <-chan *KRPCMessage
}

// 信息的中转和协调
type LocalNode struct {
	Info NodeInfo
	transport Transport
	krpc *KRPC
	nm *NodeManager
	processor []MsgProcessor
}

type RemoteNode struct {
	Info NodeInfo
	lastUpdateTime time.Time
	queryTime time.Time
}

type MsgProcessor interface {
	handle(*KRPCMessage)(*KRPCMessage, error)
}

