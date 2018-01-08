package rpc

import (
	"net"
)

type NetServer struct {
	l net.Listener
	mux *Router
	initFuncs []func(c *Conn)
}

func NewNetServer(l net.Listener, mux *Router) *NetServer {
	return &NetServer{
		l: l,
		mux: mux,
	}
}

func (ns *NetServer) Handle() error {
	for {
		c, err := ns.l.Accept()
		if err != nil {
			return err
		}
		conn := Wrap(c, ns.mux)
		for _, initFunc := range ns.initFuncs {
			go initFunc(conn)
		}
	}
}

func (ns *NetServer) AddInit(initFunc func(c *Conn)) {
	ns.initFuncs = append(ns.initFuncs, initFunc)
}

