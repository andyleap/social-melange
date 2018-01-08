package messages

import (
	"github.com/andyleap/social-melange/proto/messages"
)

type ItemStore interface {
	GetItem(*messages.Ref) (*messages.SignedItem, error)
	PutItem(*messages.SignedItem) error
}

type FeedStore interface {
	GetSeq(feed *messages.Feed, seq uint64) (*messages.SignedFeedChange, error)
	PutSeq(feed *messages.Feed, fc *messages.SignedFeedChange) error
	GetSeqs(feed *messages.Feed) ([]uint64, error)
}