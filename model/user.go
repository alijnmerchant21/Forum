/*
This code defines a struct User in the model package. The User struct has the following fields:

Name: a string representing the name of the user

PubKey: an ed25519.PubKey type, which is a public key for the user. This field is tagged with badgerhold:"index" which indicates that it should be indexed in the Badger DB.

Moderator: a boolean indicating whether the user is a moderator or not

Banned: a boolean indicating whether the user is banned or not

NumMessages: an integer representing the number of messages the user has sent

Version: an unsigned 64-bit integer representing the version of the user

SchemaVersion: an integer representing the schema version of the user

The User struct is used to represent user data in a Go program, and can be used to store user data in a database or manipulate it in other ways.
The ed25519.PubKey type is imported from the github.com/cometbft/cometbft/crypto/ed25519 package, which provides cryptographic functions for use in a distributed systems environment.
*/

package model

import (
	"github.com/cometbft/cometbft/crypto/ed25519"
)

type User struct {
	Name          string
	PubKey        ed25519.PubKey `badgerhold:"index"` // this is just a wrapper around bytes
	Moderator     bool
	Banned        bool
	NumMessages   int64
	Version       uint64
	SchemaVersion int
}
