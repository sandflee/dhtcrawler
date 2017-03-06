package main
import (
	"flag"
	"net"
	"fmt"
	"time"
	"encoding/binary"
	"github.com/sandflee/dhtcrawler/dht"
)

func makeOutput(input []byte) []byte {
	now := time.Now().Unix();
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(now))
	return b
}

func process(input []byte, addr,raddr *net.UDPAddr) ([]byte,error) {
	node := dht.NewNode(dht.GenerateID(), addr.IP)
	krpc := dht.NewKRPC(node)
	fmt.Printf("recived msg:\n", string(input))
	if msg, err := krpc.Decode(string(input[0:]), raddr); err != nil {
		return nil, err
	} else {
		fmt.Printf("encode msg:%+v\n", msg.String())
		pong, err := krpc.EncodeingPong("444")
		fmt.Printf("response msg:%+v\n", pong)
		return []byte(pong),err
	}
}

func main() {
	var host,port string
	flag.StringVar(&host, "host", "0.0.0.0", "host")
	flag.StringVar(&port, "port", "8888", "port")
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp", host + ":" + port)
	if err != nil {
		fmt.Println("resolve upd add failed,", err)
		return
	}

	fmt.Printf("test,%+v\n", addr)

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("listen failed,", err)
		return
	}
	defer conn.Close()

	for {
		input := make([]byte, 4086)
		n, in, err := conn.ReadFromUDP(input)
		if err != nil {
			fmt.Println("read failed,", err)
			return
		}
		//output := makeOutput(input)
		output,err := process(input[0:n], addr, in)
		if err != nil {
			fmt.Println("process failed,", err)
			return
		}
		if _, err := conn.WriteToUDP(output, in); err != nil {
			fmt.Println("write failed,", err)
			return
		}
	}
}