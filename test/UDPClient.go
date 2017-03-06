package main
import (
	"flag"
	"net"
	"fmt"
	"github.com/sandflee/dhtcrawler/dht"
)


func generateOutPutMsg(srcAddr *net.UDPAddr)([]byte,error) {
	node := dht.NewNode(dht.GenerateID(), srcAddr.IP)
	krpc := dht.NewKRPC(node)
	if _, output, err := krpc.EncodeingPing(); err == nil {
		fmt.Println("output:",output)
		return []byte(output),nil
	} else {
		return nil,err
	}
}

func decodeMeg(msg string, srcAddr *net.UDPAddr) {
	node := dht.NewNode(dht.GenerateID(), srcAddr.IP)
	krpc := dht.NewKRPC(node)
	fmt.Println("recv msg:", msg)
	if m, err := krpc.Decode(msg, srcAddr); err != nil {
		fmt.Println("decode failed,", err)
		return
	} else {
		fmt.Println("decode msg:", m)
	}

}

func main() {
	var host string
	var port int
	flag.StringVar(&host, "host", "0.0.0.0", "host")
	flag.IntVar(&port, "port", 8888, "port")
	flag.Parse()

	srcAddr := &net.UDPAddr{IP:net.IPv4zero, Port:0}
	dstAddr := &net.UDPAddr{IP:net.ParseIP(host), Port: port}

	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println()
		return
	}
	defer conn.Close()

	output,err := generateOutPutMsg(srcAddr)
	if err != nil {
		fmt.Println("generate output failed", err)
		return
	}

	_, err = conn.Write(output)
	if err != nil {
		fmt.Println("write failed", err)
		return
	}

	buf := make([]byte, 1024)
	if n, err := conn.Read(buf); err != nil {
		fmt.Println("read failed", err)
	} else {
		fmt.Println("<%s>,", conn.RemoteAddr())
		decodeMeg(string(buf[0:n]), dstAddr)
	}
}