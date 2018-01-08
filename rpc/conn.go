package rpc

import (
	"context"
	"log"
	"net"
	"reflect"
	"sync"

	"github.com/andyleap/social-melange/proto/any"
	"github.com/andyleap/social-melange/proto/rpc"

	"github.com/golang/protobuf/proto"
)

type CallHandler interface {
	Resp(proto.Message)
	Error(error)
}

type Conn struct {
	conn net.Conn
	ps     PacketStream
	mux    *Router
	ctx    context.Context
	cancel context.CancelFunc
	closers []func()

	respMapMu sync.Mutex
	respMap   map[uint64]CallHandler
	nextID    uint64
}

func Wrap(c net.Conn, mux *Router) *Conn {
	ps := &NetStream{rw: c}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	conn := &Conn{
		conn: c,
		mux:     mux,
		ps:      ps,
		ctx:     ctx,
		cancel:  cancel,
		respMap: map[uint64]CallHandler{},
		nextID:  1,
	}
	go conn.run()
	return conn
}

type RemoteError string

func (re RemoteError) Error() string {
	return string(re)
}

func (c *Conn) AddCloser(closer func()) {
	c.closers = append(c.closers, closer)
}

func (c *Conn) Close() {
	c.cancel()
	c.conn.Close()
	for _, closer := range c.closers {
		closer()
	}
}

func (c *Conn) run() {
	defer c.Close()
	for {
		p, err := c.ps.Recv()
		if err != nil {
			return
		}
		switch p := p.Packet.(type) {
		case *rpc.Packet_Request:
			c.mux.handleReq(c.ctx, c, p.Request)
		case *rpc.Packet_Response:
			func() {
				c.respMapMu.Lock()
				defer c.respMapMu.Unlock()
				resp, ok := c.respMap[p.Response.GetID()]
				if !ok {
					log.Println("Unsolicitied response!", p.Response.String())
					return
				}
				if p.Response.GetEnd() {
					delete(c.respMap, p.Response.GetID())
				}
				if p.Response.GetError() != "" {
					resp.Error(RemoteError(p.Response.GetError()))
					return
				}
				raw := p.Response.GetResponse()
				typ := raw.GetType()
				t := proto.MessageType(typ)
				if t == nil {
					log.Println("Unknown response type!", p.Response.String())
					return
				}
				msg := reflect.New(t).Interface().(proto.Message)
				go func() {
					defer func() {
						err := recover()
						if err != nil {
							log.Println(err)
						}
					}()
					resp.Resp(msg)
				}()
			}()
		}
	}
}

func (c *Conn) Call(msg proto.Message, ch CallHandler) {
	c.respMapMu.Lock()
	defer c.respMapMu.Unlock()
	id := c.nextID
	c.nextID++
	c.respMap[id] = ch

	typ := proto.MessageName(msg)
	buf, _ := proto.Marshal(msg)
	any := &any.Any{
		Type:  typ,
		Value: buf,
	}

	c.ps.Send(&rpc.Packet{
		Packet: &rpc.Packet_Request{
			Request: &rpc.Request{
				ID:      id,
				Request: any,
			},
		},
	})
}
