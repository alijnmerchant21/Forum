package model

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"sync"

	"github.com/cometbft/cometbft/crypto/ed25519"
	"github.com/dgraph-io/badger"
	"github.com/timshannon/badgerhold"
)

var bufPool = sync.Pool{New: func() any { return &bytes.Buffer{} }}

type DB struct {
	store *badgerhold.Store
}

func New(dbPath string) (*DB, error) {
	store, err := badgerhold.Open(
		badgerhold.Options{
			Encoder: func(v interface{}) ([]byte, error) {
				jby, err := json.Marshal(v)
				if err != nil {
					return nil, err
				}

				buf := bufPool.Get().(*bytes.Buffer)
				defer bufPool.Put(buf)
				buf.Reset()
				gz, err := flate.NewWriter(buf, 5)
				if err != nil {
					return nil, err
				}
				defer gz.Close()

				if _, err := gz.Write(jby); err != nil {
					return nil, err
				}

				return buf.Bytes(), nil
			},
			Decoder: func(in []byte, val interface{}) error {
				buf := bufPool.Get().(*bytes.Buffer)
				defer bufPool.Put(buf)
				buf.Reset()

				gz := flate.NewReader(buf)
				defer gz.Close()

				if _, err := gz.Read(in); err != nil {
					return err
				}
				return json.Unmarshal(buf.Bytes(), val)
			},
			Options: badger.DefaultOptions(dbPath),
		},
	)
	if err != nil {
		return nil, err
	}

	return &DB{
		store: store,
	}, nil

}

func (db *DB) Close() error { return db.store.Close() } // This code defines a function called Close() that is a method of a struct type DB. The function returns an error and when called, it will close the database connection by calling a method Close() on an object store of the same type as DB.

func (db *DB) CreateUser(u *User) error { return db.store.Insert(u.PubKey, u) } // This Go code defines a method called CreateUser which takes a pointer to a DB struct and a pointer to a User struct as its arguments. The method returns an error. Inside the method, the Insert method of a store field on the DB struct is called with the User struct's PubKey and the User struct itself as arguments.

func (db *DB) SaveUser(u *User) error                     { return db.store.Update(u.PubKey, u) } // This code snippet defines a method SaveUser for a DB struct that takes a pointer to a User struct and returns an error. The method updates the User object in the database by calling the Update method of a store field in the DB struct, passing in the PubKey field of the User object and the User object itself.
func (db *DB) FindUser(key ed25519.PubKey) (*User, error) { return nil, nil }                     // This code snippet defines a method FindUser for a DB struct that takes an ed25519 public key and returns either a pointer to a User struct or an error. However, the implementation is incomplete as it always returns nil for the User and nil for the error.
