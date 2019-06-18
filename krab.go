package krab

import (
	"bytes"
	"errors"
	"strings"

	"github.com/RTradeLtd/crypto/v2"
	datastore "github.com/ipfs/go-datastore"
	ds "github.com/ipfs/go-datastore"
	namespace "github.com/ipfs/go-datastore/namespace"
	"github.com/ipfs/go-datastore/query"
	kstore "github.com/ipfs/go-ipfs-keystore"
	ci "github.com/libp2p/go-libp2p-core/crypto"
)

// compile time check for interface compatability
var _ kstore.Keystore = (*Keystore)(nil)

// Keystore is used to manage an encrypted IPFS keystore
type Keystore struct {
	em *crypto.EncryptManager
	ds datastore.Batching
}

// NewKeystore is used to create a new krab ipfs keystore manager
func NewKeystore(ds datastore.Batching, passphrase string) (*Keystore, error) {
	wrapDS := namespace.Wrap(ds, datastore.NewKey("/krabkeystore"))
	return &Keystore{
		em: crypto.NewEncryptManager(passphrase),
		ds: wrapDS,
	}, nil
}

// Has is used to check whether or not the given key name exists
func (km *Keystore) Has(name string) (bool, error) {
	if err := validateName(name); err != nil {
		return false, err
	}
	if has, err := km.ds.Has(ds.NewKey(name)); err != nil {
		return false, err
	} else if !has {
		return false, errors.New(ErrNoSuchKey)
	}
	return true, nil
}

// Put is used to store a key in our keystore
func (km *Keystore) Put(name string, privKey ci.PrivKey) error {
	if err := validateName(name); err != nil {
		return err
	}
	if has, err := km.Has(name); err == nil {
		return errors.New(ErrKeyExists)
	} else if has {
		return errors.New(ErrKeyExists)
	}
	pkBytes, err := privKey.Bytes()
	if err != nil {
		return err
	}
	reader := bytes.NewReader(pkBytes)
	// encrypt the private key
	encryptedPK, err := km.em.Encrypt(reader)
	if err != nil {
		return err
	}
	return km.ds.Put(ds.NewKey(name), encryptedPK)
}

// Get is used to retrieve a key from our keystore
func (km *Keystore) Get(name string) (ci.PrivKey, error) {
	if err := validateName(name); err != nil {
		return nil, err
	}
	if has, err := km.Has(name); err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New(ErrNoSuchKey)
	}
	encryptedPKBytes, err := km.ds.Get(ds.NewKey(name))
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(encryptedPKBytes)
	pkBytes, err := km.em.Decrypt(reader)
	if err != nil {
		return nil, err
	}
	return ci.UnmarshalPrivateKey(pkBytes)
}

// Delete is used to remove a key from our keystore
func (km *Keystore) Delete(name string) error {
	if err := validateName(name); err != nil {
		return err
	}
	return km.ds.Delete(ds.NewKey(name))
}

// List is used to list all key identifiers in our keystore
func (km *Keystore) List() ([]string, error) {
	entries, err := km.ds.Query(query.Query{})
	if err != nil {
		return nil, err
	}
	keys, err := entries.Rest()
	if err != nil {
		return nil, err
	}
	var ids []string
	for _, v := range keys {
		ids = append(ids, strings.Split(v.Key, "/")[1])
	}
	return ids, nil
}
