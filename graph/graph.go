package graph

import (
	"github.com/andyleap/social-melange/graph/proto"
	pmsg "github.com/andyleap/social-melange/proto/messages"
	"github.com/andyleap/social-melange/messages"
	
	"github.com/golang/protobuf/proto"
)

type Follow struct {
	users map[uint64]*pmsg.Signer
	
	FS messages.FeedStore
	IS messages.ItemStore
}

func (f *Follow) ProcessMessage(feed *pmsg.Feed, sfc *pmsg.SignedFeedChange) {
	for _, delSeq := range sfc.GetFeedChange().GetDelete() {
		delete(f.users, delSeq)
	}
	item, err := f.IS.GetItem(sfc.GetFeedChange().GetRef())
	if err != nil {
		return
	}
	if item.GetItem().GetMessage().Type == "graph.follow" {
		follow := &graph.Follow{}
		err := proto.Unmarshal(item.GetItem().GetMessage().GetValue(), follow)
		if err != nil {
			return
		}
		f.users[sfc.GetFeedChange().GetSequence()] = follow.GetFollow()
	}
}