package main

import (
	"context"
	"flag"
	"os"

	I "libp2p-consumer/gossip/init"
	P "libp2p-consumer/gossip/producer"
	S "libp2p-consumer/gossip/subscriber"
)

var address string

func main() {
	ctx := context.Background()

	listenCmd := flag.NewFlagSet("start", flag.ExitOnError)
	isBootstrap := listenCmd.Bool("isBootstrap", false, "Is Bootstrap Peer?")
	bootstrapIP := listenCmd.String("bootstrap", "0.0.0.0", "Bootstrap Peer IP Address")
	peerID := listenCmd.String("peerid", "12", "Peer ID")

	switch os.Args[1] {

	case "producer":
		listenCmd.Parse(os.Args[2:])
		address = "/ip4/" + *bootstrapIP + "/tcp/7654/p2p/" + *peerID
		topic, _ := I.Init(ctx, *isBootstrap, address)
		P.Producer(ctx, topic)

	case "subscriber":
		listenCmd.Parse(os.Args[2:])
		address = "/ip4/" + *bootstrapIP + "/tcp/7654/p2p/" + *peerID
		topic, host := I.Init(ctx, *isBootstrap, address)
		S.Subscriber(ctx, host, topic)

	default:
		os.Exit(1)
	}
}
