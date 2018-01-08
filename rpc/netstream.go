package rpc

import (
	"io"
	"sync"

	"github.com/andyleap/social-melange/proto/rpc"

	"github.com/golang/protobuf/proto"
)

type NetStream struct {
	rw io.ReadWriter
	writeLock sync.Mutex
}

func (ns *NetStream) Recv() (*rpc.Packet, error) {
	sizeb := make([]byte, 4)
	io.ReadFull(ns.rw, sizeb)
	b := proto.NewBuffer(sizeb)
	size, err := b.DecodeFixed32()
	if err != nil {
		return nil, err
	}

	buf := make([]byte, size)
	_, err = io.ReadFull(ns.rw, buf)
	if err != nil {
		return nil, err
	}

	p := &rpc.Packet{}
	err = proto.Unmarshal(buf, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (ns *NetStream) Send(p *rpc.Packet) error {
	ns.writeLock.Lock()
	defer ns.writeLock.Unlock()
	buf, err := proto.Marshal(p)
	if err != nil {
		return err
	}

	b := proto.NewBuffer(nil)
	b.EncodeFixed32(uint64(len(buf)))
	_, err = ns.rw.Write(b.Bytes())
	if err != nil {
		return err
	}
	
	ns.rw.Write(buf)
	return nil
}

