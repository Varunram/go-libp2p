package metrics

import (
	protocol "github.com/libp2p/go-libp2p/gxlibs/github.com/libp2p/go-libp2p-protocol"
	peer "github.com/libp2p/go-libp2p/gxlibs/github.com/libp2p/go-libp2p-peer"
)

type StreamMeterCallback func(int64, protocol.ID, peer.ID)
type MeterCallback func(int64)

type Reporter interface {
	LogSentMessage(int64)
	LogRecvMessage(int64)
	LogSentMessageStream(int64, protocol.ID, peer.ID)
	LogRecvMessageStream(int64, protocol.ID, peer.ID)
	GetBandwidthForPeer(peer.ID) Stats
	GetBandwidthForProtocol(protocol.ID) Stats
	GetBandwidthTotals() Stats
}
