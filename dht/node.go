package dht
import (
	"net"
	"github.com/golang/glog"
)

func NewNode(identifier Identifier, addr *net.UDPAddr) *LocalNode {
	info := NodeInfo{
		IP: addr.IP,
		Port: addr.Port,
		ID: identifier,
	}
	kRpc := NewKRPC(info)

	nodeM := NewNodeManager()

	node := &LocalNode{
		Info: info,
		transport: Transport{},
		krpc: kRpc,
		nm: nodeM,
		processor: make([]MsgProcessor,2),
	}
	node.processor = append(node.processor, nodeM)
	node.processor = append(node.processor, node)

	return node
}

func (node *LocalNode) handle(msg *KRPCMessage)(*KRPCMessage, error) {
	return nil,nil
}

func (node *LocalNode) processKrpcMessage(msg *KRPCMessage) (*KRPCMessage) {
	var response *KRPCMessage
	for _,p := range node.processor {
		if res, err := p.handle(msg); err != nil {
			glog.Error("process msg failed,", err)
		} else {
			if res != nil {
				response = res
			}
		}
	}
	return response
}

func (node *LocalNode) Run() {
	recvC := make(chan *KRPCMessage)
	err := node.transport.startUdpServer(recvC)
	if err != nil {
		glog.Error("start udp server failed,")
	}
	nmC := make(chan *KRPCMessage)
	go node.nm.Run(nmC)

	select {
	case krpcMessage := <-recvC:
		resp := node.processKrpcMessage(krpcMessage)
		if resp != nil {
			// to do some process
			node.transport.sendKrpcMessage(resp)
		}
	case msg := <-nmC:
		node.transport.sendKrpcMessage(msg)
	}
}
