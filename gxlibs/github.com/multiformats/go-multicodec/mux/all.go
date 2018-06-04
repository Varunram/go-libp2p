package muxcodec

import (
	mc "github.com/libp2p/go-libp2p/gxlibs/github.com/multiformats/go-multicodec"
	cbor "github.com/libp2p/go-libp2p/gxlibs/github.com/multiformats/go-multicodec/cbor"
	json "github.com/libp2p/go-libp2p/gxlibs/github.com/multiformats/go-multicodec/json"
)

func StandardMux() *Multicodec {
	return MuxMulticodec([]mc.Multicodec{
		cbor.Multicodec(),
		json.Multicodec(false),
		json.Multicodec(true),
	}, SelectFirst)
}
