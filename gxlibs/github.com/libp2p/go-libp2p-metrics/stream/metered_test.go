package meterstream

import (
	"io"
	"io/ioutil"
	"math/rand"
	"testing"
	"time"

	inet "github.com/libp2p/go-libp2p/gxlibs/github.com/libp2p/go-libp2p-net"
	protocol "github.com/libp2p/go-libp2p/gxlibs/github.com/libp2p/go-libp2p-protocol"
	peer "github.com/libp2p/go-libp2p/gxlibs/github.com/libp2p/go-libp2p-peer"
)

type FakeStream struct {
	ReadBuf io.Reader
	inet.Stream
}

func (fs *FakeStream) Read(b []byte) (int, error) {
	return fs.ReadBuf.Read(b)
}

func (fs *FakeStream) Write(b []byte) (int, error) {
	return len(b), nil
}

func (fs *FakeStream) Protocol() protocol.ID {
	return "TEST"
}

func TestCallbacksWork(t *testing.T) {
	fake := new(FakeStream)

	var sent int64
	var recv int64

	sentCB := func(n int64, proto protocol.ID, p peer.ID) {
		sent += n
	}

	recvCB := func(n int64, proto protocol.ID, p peer.ID) {
		recv += n
	}

	ms := newMeteredStream(fake, peer.ID("PEER"), recvCB, sentCB)

	toWrite := int64(100000)
	toRead := int64(100000)

	a := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := rand.New(rand.NewSource(time.Now().UnixNano() + 1))
	fake.ReadBuf = io.LimitReader(a, toRead)
	writeData := io.LimitReader(b, toWrite)

	n, err := io.Copy(ms, writeData)
	if err != nil {
		t.Fatal(err)
	}

	if n != toWrite {
		t.Fatal("incorrect write amount")
	}

	if toWrite != sent {
		t.Fatal("incorrectly reported writes", toWrite, sent)
	}

	n, err = io.Copy(ioutil.Discard, ms)
	if err != nil {
		t.Fatal(err)
	}

	if n != toRead {
		t.Fatal("incorrect read amount")
	}

	if toRead != recv {
		t.Fatal("incorrectly reported reads")
	}
}
