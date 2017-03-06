package main
import (
	"github.com/sandflee/dhtcrawler/dht"
	"flag"
	"net"
	"github.com/golang/glog"
)

const (
	DEFAULT_DHT_LISTEN_PORT = ""
)

func main() {
	var ip,port string
	flag.StringVar(&ip, "ip", "127.0.0.1","external ip")
	flag.StringVar(&port, "port", DEFAULT_DHT_LISTEN_PORT, "dht listen port")
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", ip + ":" + port)
	if err != nil {
		glog.Error("resolve upd add failed,", err)
		return
	}

	node := dht.NewNode(dht.GenerateID(), addr)
	node.Run()
}