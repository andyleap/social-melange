package replication

import (
	"context"
	"time"
	"sync"

	"github.com/andyleap/social-melange/messages"
	pmsg "github.com/andyleap/social-melange/proto/messages"
	"github.com/andyleap/social-melange/proto/replication"
	"github.com/andyleap/social-melange/rpc"
)

type Replication struct {
	FS messages.FeedStore
	IS messages.ItemStore

	currentConnMu sync.Mutex
	currentConns map[*rpc.Conn][][]byte
}

func New() *Replication {
	return &Replication{
		currentConns: map[*rpc.Conn][][]byte{},
	}
}

func (r *Replication) handleSignerRequest(ctx context.Context, req *replication.SignerRequest, rw rpc.ResponseWriter) {
	done := ctx.Done()

	curPos := req.GetFeeds()
	for {
		msg := map[string]uint64{}
		for feed, val := range curPos {
			seqs, _ := r.FS.GetSeqs(&pmsg.Feed{Signer: req.GetSigner(), Name: feed})
			newVal := val
			for _, s := range seqs {
				if s > val {
					newVal = s
				}
			}
			if newVal > val {
				msg[feed] = newVal
			}
			curPos[feed] = newVal
		}
		rw.Send(&replication.SignerUpdate{Feeds: msg}, false)
		time.Sleep(30 * time.Second)
		select {
		case <-done:
			return
		default:
		}
	}
}

func (r *Replication) handleFeedRequest(ctx context.Context, req *replication.FeedRangeRequest, rw rpc.ResponseWriter) {
	ret := map[uint64]*pmsg.SignedFeedChange{}
	for l1 := req.Start; l1 <= req.Stop; l1++ {
		ref, err := r.FS.GetSeq(req.GetFeed(), l1)
		if err != nil {
			continue
		}
		ret[l1] = ref
	}
	rw.Send(&replication.FeedRangeResponse{FeedChanges: ret}, true)
}

func (r *Replication) handleItemRequest(ctx context.Context, req *replication.ItemRequest, rw rpc.ResponseWriter) {
	i, err := r.IS.GetItem(req.GetRef())
	if err != nil {
		rw.Error(err)
		return
	}
	rw.Send(&replication.ItemResponse{Item: i}, true)
}

func (r *Replication) Register(mux *rpc.Router) {
	mux.Handle(r.handleSignerRequest)
	mux.Handle(r.handleFeedRequest)
	mux.Handle(r.handleItemRequest)
}

func (r *Replication) NewConn(c *rpc.Conn) {
	r.currentConnMu.Lock()
	defer r.currentConnMu.Unlock()
	r.currentConns[c] = [][]byte{}
	c.AddCloser(func() {
		r.currentConnMu.Lock()
		defer r.currentConnMu.Unlock()
		delete(r.currentConns, c)
	})
}






