package model

import (
	"fmt"

	"github.com/cometbft/cometbft/crypto/ed25519"
)

type MsgSendMessage struct {
	Text string
	From ed25519.PubKey
}

type MsgSetBan struct {
	User  ed25519.PubKey
	State bool
}

type MsgSetModerator struct {
	User  ed25519.PubKey
	State bool
}

type MsgCreateUser struct {
	User User
}

/*
This code snippet is defining a function Process for a DB struct. The function takes an interface message
as an argument and returns an error. The function then switches on the type of the message and performs different
actions depending on the type of the message.
If the message is a MsgSendMessage, it finds the user associated with the message, updates some fields in the user
struct, and saves the user to the database.
Similarly, for MsgSetBan and MsgSetModerator, it finds the user, updates some fields, and saves the user to
the database.
For MsgCreateUser, it initializes the user's version to 0 and creates a new user in the database.
If the message type is not supported, an error is returned.
*/
func (db *DB) Process(message interface{}) error {
	switch msg := message.(type) {
	case MsgSendMessage:
		u, err := db.FindUser(msg.From)
		if err != nil {
			return err
		}
		// TODO: implement business logic

		u.Version++
		u.NumMessages++
		return db.SaveUser(u)
	case MsgSetBan:
		u, err := db.FindUser(msg.User)
		if err != nil {
			return err
		}

		u.Version++
		u.Banned = msg.State

		return db.SaveUser(u)
	case MsgSetModerator:
		u, err := db.FindUser(msg.User)
		if err != nil {
			return err
		}
		u.Version++
		u.Moderator = msg.State

		return db.SaveUser(u)
	case MsgCreateUser:
		msg.User.Version = 0
		return db.CreateUser(&msg.User)
	default:
		return fmt.Errorf("message type %T not supported", message)
	}
}
