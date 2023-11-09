package producer

import (
	"bufio"
	"context"
	"fmt"
	"os"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

func Producer(ctx context.Context, topic *pubsub.Topic) {
	publish(ctx, topic)
}

// start publisher to topic
func publish(ctx context.Context, topic *pubsub.Topic) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Printf("enter message to publish: \n")

			msg := scanner.Text()
			if len(msg) != 0 {
				// publish message to topic
				bytes := []byte(msg)
				topic.Publish(ctx, bytes)
			}
		}
	}
}
