package main

import (
	"net"

	"github.com/andyleap/social-melange/rpc"
	"github.com/andyleap/social-melange/replication"
	"github.com/andyleap/social-melange/messages/fsstore"
)



func main() {
	is := fsstore.NewItemStore("data/items")
	fs := fsstore.NewFeedStore("data/feeds")
	repl := &replication.Replication{
		IS: is,
		FS: fs,
	}
	mux := rpc.NewRouter()
	repl.Register(mux)
	
	l, err := net.Listen("tcp", ":4589")
	ns := rpc.NewNetServer(l, mux)
	
}