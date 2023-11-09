package subscriber

import (
	"context"
	"fmt"

	pubsub "github.com/libp2p/go-libp2p-pubsub"

	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"

	I "libp2p-consumer/gossip/init"
	U "libp2p-consumer/utils"
)

func Subscriber(ctx context.Context, host host.Host, topic *pubsub.Topic) {
	// subscribe to topic
	subscriber, err := topic.Subscribe()
	U.CheckErr(err)

	subscribe(subscriber, ctx, host.ID())
}

// start subsriber to topic
func subscribe(subscriber *pubsub.Subscription, ctx context.Context, hostID peer.ID) {
	for {
		msg, err := subscriber.Next(ctx)
		if err != nil {
			panic(err)
		}

		// only consider messages delivered by other peers
		if msg.ReceivedFrom == hostID {
			continue
		}

		// 여기가 메세지 수신했을 때 출력하는 부분
		fmt.Printf("got message: %s, from: %s\n", string(msg.Data), msg.ReceivedFrom)
		I.SendACK(hostID)
	}
}
