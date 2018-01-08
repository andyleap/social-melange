package fsstore

import (
	"encoding/hex"
	"io/ioutil"
	"path/filepath"
	"os"

	"github.com/andyleap/social-melange/proto/messages"

	"github.com/golang/protobuf/proto"
	"golang.org/x/crypto/sha3"
)

type FSItemStore struct {
	basePath string
	defAlgo messages.Algo
}

func NewItemStore(basePath string) *FSItemStore {
	return &FSItemStore{
		basePath: basePath,
	}
}

func (fs *FSItemStore) GetItem(ref *messages.Ref) (*messages.SignedItem, error) {
	algo := ref.GetAlgo().String()
	hash := hex.EncodeToString(ref.GetHash())
	itempath := filepath.Join(fs.basePath, algo, hash[0:2], hash[0:4], hash)

	buf, err := ioutil.ReadFile(itempath)
	if err != nil {
		return nil, err
	}

	ret := &messages.SignedItem{}
	err = proto.Unmarshal(buf, ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (fs *FSItemStore) PutItem(item *messages.SignedItem) error {
	algo := fs.defAlgo.String()

	buf, err := proto.Marshal(item)
	if err != nil {
		return err
	}

	hash := ""

	switch fs.defAlgo {
	case messages.Algo_SHA3_512:
		sha := sha3.New512()
		sha.Write(buf)
		hash = hex.EncodeToString(sha.Sum(nil))
	}

	dirpath := filepath.Join(fs.basePath, algo, hash[0:2], hash[0:4])
	itempath := filepath.Join(dirpath, hash)
	
	err = os.MkdirAll(dirpath, 0777)
	if err != nil {
		return err
	}	
	
	
	err = ioutil.WriteFile(itempath, buf, 0666)
	if err != nil {
		return err
	}
	return nil	
}

