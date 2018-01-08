package rpc

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/andyleap/social-melange/proto/any"
	"github.com/andyleap/social-melange/proto/rpc"

	"github.com/golang/protobuf/proto"
)

type PacketStream interface {
	Send(*rpc.Packet) error
	Recv() (*rpc.Packet, error)
}

type ResponseWriter interface {
	Send(msg proto.Message, end bool) error
	Error(err error)
}

type responseWriter struct {
	ps PacketStream
	id uint64
}

func (rw *responseWriter) Send(msg proto.Message, end bool) error {
	rsp := &rpc.Response{
		ID:  rw.id,
		End: end,
	}
	buf, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	rsp.Response = &any.Any{
		Type:  proto.MessageName(msg),
		Value: buf,
	}
	return rw.ps.Send(&rpc.Packet{
		Packet: &rpc.Packet_Response{
			Response: rsp,
		},
	})
}

func (rw *responseWriter) Error(err error) {
	rw.ps.Send(&rpc.Packet{
		Packet: &rpc.Packet_Response{
			Response: &rpc.Response{
				ID:    rw.id,
				Error: err.Error(),
				End:   true,
			},
		},
	})
}

type handler struct {
	f   reflect.Value
	typ reflect.Type
}

type Router struct {
	handlers map[string]*handler
}

func NewRouter() *Router {
	return &Router{
		handlers: map[string]*handler{},
	}
}

func (r *Router) Handle(handle interface{}) {
	val := reflect.ValueOf(handle)
	t := val.Type()
	if t.NumIn() != 3 {
		panic(fmt.Errorf("Unexpected number of arguments for handler!  Expected 3, got %d!", t.NumIn()))
	}
	if t.In(0) != reflect.ValueOf(context.Context(nil)).Type() {
		panic(fmt.Errorf("Unexpected first argument for handler!  Expected context.Context, got %s!", t.In(0).String()))		
	}
	if !t.In(1).Implements(reflect.ValueOf(proto.Message(nil)).Type()) {
		panic(fmt.Errorf("Unexpected second argument for handler!  Expected implements proto.Message, got %s!", t.In(1).String()))		
	}
	if t.In(2) != reflect.ValueOf(ResponseWriter(nil)).Type() {
		panic(fmt.Errorf("Unexpected third argument for handler!  Expected rpc.ResponseWriter, got %s!", t.In(2).String()))		
	}
	msg := reflect.New(t.In(1)).Interface().(proto.Message)
	typ := proto.MessageName(msg)
	r.handlers[typ] = &handler{
		f: val,
		typ: t.In(1),
	}
}

func (r *Router) handleReq(ctx context.Context, c *Conn, req *rpc.Request) {
	rw := responseWriter{c.ps, req.GetID()}
	reqAny := req.GetRequest()
	typ := reqAny.GetType()

	handler := r.handlers[typ]
	if handler == nil {
		rw.Error(fmt.Errorf("Unable to find handler for type: %s", typ))
		return
	}

	msg := reflect.New(handler.typ).Interface().(proto.Message)
	proto.Unmarshal(reqAny.GetValue(), msg)

	go func() {
		defer func() {
			err := recover()
			if err != nil {
				rw.Error(fmt.Errorf("Error while handling request"))
				log.Println(err)
			}
		}()
		handler.f.Call([]reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(msg), reflect.ValueOf(rw)})
	}()
}
