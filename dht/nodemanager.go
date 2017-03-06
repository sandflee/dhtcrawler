package dht

BOOTSTRAP_NODES = (
("router.bittorrent.com", 6881),
("dht.transmissionbt.com", 6881),
("router.utorrent.com", 6881)
)

func NewNodeManager() *NodeManager  {
	r := &Route{}

	msgRecvC := make(chan *KRPCMessage)
	return &NodeManager{
		route: r,
		msgRecvC:msgRecvC,
	}
}


func (nm *NodeManager) handle(*KRPCMessage)(*KRPCMessage, error) {
	return nil,nil
}

func nodeSelector(chan<- *KRPCMessage) {

}

func nodeExpireChecker(chan<- *KRPCMessage) {

}

func processKrpcMsg(msg *KRPCMessage) {

}

func (nm *NodeManager) Run(sender chan<-*KRPCMessage) {
	go nodeSelector(sender)
	go nodeExpireChecker(sender)

	select {
	case msg := <-nm.msgRecvC:
		processKrpcMsg(msg)
	}
}
