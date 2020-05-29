package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 开Server
func ServRpc(host string, service interface{}) error {
	rpc.Register(service)

	listener, err := net.Listen("tcp4", host)

	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
	return nil
}

// 获取server的链接
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp4", host)
	if nil != err {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
