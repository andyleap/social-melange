package messages

import (
	"crypto"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/ed25519"

	"github.com/golang/protobuf/proto"
)

func (i *Item) Sign(pk crypto.Signer) (*SignedItem, error) {
	pubkey := []byte{}
	switch pubk := pk.Public().(type) {
	case ed25519.PublicKey:
		pubkey = pubk
	default:
		return nil, fmt.Errorf("Unknown public key: %T", pk.Public())
	}
	
	i.Author = &Signer{Key: pubkey}
	
	buf, err := proto.Marshal(i)
	if err != nil {
		return nil, err
	}
	
	sig, err := pk.Sign(rand.Reader, buf, nil)
	if err != nil {
		return nil, err
	}
	
	si := &SignedItem{
		Item: i,
		Signature: sig,
	}
	return si, nil
}

func (fc *FeedChange) Sign(pk crypto.Signer) (*SignedFeedChange, error) {
	buf, err := proto.Marshal(fc)
	if err != nil {
		return nil, err
	}
	
	sig, err := pk.Sign(rand.Reader, buf, nil)
	if err != nil {
		return nil, err
	}
	
	sfc := &SignedFeedChange{
		FeedChange: fc,
		Signature: sig,
	}
	return sfc, nil
}

func (sfc *SignedFeedChange) Verify(pk crypto.PublicKey) bool {
	buf, err := proto.Marshal(sfc.FeedChange)
	if err != nil {
		return false
	}
	switch pk := pk.(type) {
	case ed25519.PublicKey:
		return ed25519.Verify(pk, buf, sfc.Signature)
	}
	return false
}

func (s *Signer) PublicKey() crypto.PublicKey {
	switch s.Kind {
	case KeyKind_ED25519:
		return ed25519.PublicKey(s.GetKey())
	}
	return nil
}

