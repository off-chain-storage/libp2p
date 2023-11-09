package init

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/host"

	"github.com/multiformats/go-multiaddr"

	U "libp2p-consumer/utils"
)

func Init(ctx context.Context, isBootstrap bool, address string) (*pubsub.Topic, host.Host) {
	host, err := libp2p.New(libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/7654"))
	U.CheckErr(err)

	// view host details and addresses
	fmt.Printf("host ID %s\n", host.ID())
	fmt.Printf("following are the assigned addresses\n")
	for _, addr := range host.Addrs() {
		fmt.Printf("%s\n", addr.String())
	}
	fmt.Printf("\n")

	// create a new PubSub service using the GossipSub router
	gossipSub, err := pubsub.NewGossipSub(ctx, host)
	U.CheckErr(err)

	var discoveryPeers []multiaddr.Multiaddr
	if isBootstrap {
		// When the node is a bootstrap node, don't use any discovery peers
		discoveryPeers = []multiaddr.Multiaddr{}
		fmt.Printf("Bootstrap %s\n", address)
	} else {
		// When the node is not a bootstrap node, use the provided address
		multiAddr, err := multiaddr.NewMultiaddr(address)
		U.CheckErr(err)
		discoveryPeers = []multiaddr.Multiaddr{multiAddr}
		fmt.Printf("Not Bootstrap %s\n", address)
	}

	// discoveryPeers := []multiaddr.Multiaddr{}
	dht, err := NewDHT(ctx, host, discoveryPeers)
	U.CheckErr(err)

	// setup peer discovery
	go Discover(ctx, host, dht, "librum")

	// join the pubsub topic called librum
	room := "librum"
	topic, err := gossipSub.Join(room)
	U.CheckErr(err)

	return topic, host
}
