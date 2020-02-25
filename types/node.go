package types

import (
	"net/url"
	"strings"

	"github.com/yuneejang/webserver/config"
)

type Node struct {
	Config   *config.Config
	HTTPNode *url.URL
	WsNode   *url.URL
}

func NewNode(config *config.Config) *Node {

	//	interrupt := make(chan os.Signal, 1)
	//	signal.Notify(interrupt, os.Interrupt)
	httpNode := strings.Join([]string{config.Nodes[0].Host, config.Nodes[0].HttpPort}, ":")
	wsNode := strings.Join([]string{config.Nodes[0].Host, config.Nodes[0].WsPort}, ":")
	node := &Node{
		Config: config,
		//		interrupt: interrupt,
		//		socket:    nil,
		HTTPNode: &url.URL{Scheme: "http", Host: httpNode, Path: "/"},
		WsNode:   &url.URL{Scheme: "ws", Host: wsNode, Path: "/"},
		//		handlers:  make(map[Channel]WsRpcResHandler),
	}
	return node

}
