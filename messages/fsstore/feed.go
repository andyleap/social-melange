package fsstore

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sort"

	"github.com/andyleap/social-melange/proto/messages"

	"github.com/golang/protobuf/proto"
)

type FSFeedStore struct {
	basePath string
}

func NewFeedStore(basePath string) *FSFeedStore {
	return &FSFeedStore{
		basePath: basePath,
	}
}

func (fs *FSFeedStore) GetSeq(feed *messages.Feed, seq uint64) (*messages.SignedFeedChange, error) {
	key := hex.EncodeToString(feed.GetSigner().GetKey())

	itempath := filepath.Join(fs.basePath, key[0:2], key, feed.GetName(), strconv.Itoa(int(seq)))

	buf, err := ioutil.ReadFile(itempath)
	if err != nil {
		return nil, err
	}
	ret := &messages.SignedFeedChange{}

	err = proto.Unmarshal(buf, ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (fs *FSFeedStore) PutSeq(feed *messages.Feed, sfc *messages.SignedFeedChange) error {
	key := hex.EncodeToString(feed.GetSigner().GetKey())

	dirpath := filepath.Join(fs.basePath, key[0:2], key, feed.GetName())
	itempath := filepath.Join(dirpath, strconv.Itoa(int(sfc.FeedChange.GetSequence())))

	err := os.MkdirAll(dirpath, 0777)
	if err != nil {
		return err
	}

	buf, err := proto.Marshal(sfc)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(itempath, buf, 0666)
}

func (fs *FSFeedStore) GetSeqs(feed *messages.Feed) ([]uint64, error) {
	key := hex.EncodeToString(feed.GetSigner().GetKey())

	dirpath := filepath.Join(fs.basePath, key[0:2], key, feed.GetName())
	dir, err := os.Open(dirpath)
	if err != nil {
		return nil, err
	}
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	seqs := make([]uint64, 0, len(names))
	for _, name := range names {
		seq, err := strconv.Atoi(name)
		if err != nil {
			continue
		}
		seqs = append(seqs, uint64(seq))
	}
	sort.Slice(seqs, func(i, j int) bool { return seqs[i] < seqs[j] })
	return seqs, nil
}
