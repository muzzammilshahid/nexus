package transport_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/gammazero/nexus/v3/router"
	"github.com/gammazero/nexus/v3/transport"
	"github.com/gammazero/nexus/v3/transport/serialize"
	"github.com/gammazero/nexus/v3/wamp"
	"github.com/stretchr/testify/require"
)

func TestCloseWebsocketPeer(t *testing.T) {
	routerConfig := &router.Config{
		RealmConfigs: []*router.RealmConfig{
			{
				URI: wamp.URI("nexus.test.realm"),
			},
		},
	}
	r, err := router.NewRouter(routerConfig, nil)
	require.NoError(t, err)
	defer r.Close()

	const wsAddr = "127.0.0.1:8000"
	closer, err := router.NewWebsocketServer(r).ListenAndServe(wsAddr)
	require.NoError(t, err)
	defer closer.Close()

	client, err := transport.ConnectWebsocketPeer(
		context.Background(), fmt.Sprintf("ws://%s/", wsAddr), serialize.JSON, nil, r.Logger(), nil)
	require.NoError(t, err)

	// Close the client connection.
	client.Close()

	// Try closing the client connection again. It should not cause an error.
	client.Close()
}
