/*
Explanation:
Define a Validator struct that represents a validator in the Forum application.
The NewValidator function creates a new Validator object with the given name.
The ProposeAddUser, ProposeDeleteUser, ProposeNewModerator, and ProposeBanUser functions create new
Proposal objects with the corresponding actions (add user, delete user, add moderator, and ban user).
The VoteExtension function broadcasts a proposal to all validators for voting.
The FinalizeBlock function finalizes a block and updates the User and Moderator structs accordingly based on
the proposal's action.
*/

package forum

import (
	"fmt"
)

// Validator struct represents a validator in the Forum application
type Validator struct {
	Name string
}

// NewValidator creates a new Validator object with the given name
func NewValidator(name string) *Validator {
	return &Validator{
		Name: name,
	}
}

// ProposeAddUser proposes to add a new user to the Forum
func (v *Validator) ProposeAddUser(user *User) *Proposal {
	// Create a new proposal to add the user
	action := "add_user"
	proposal := NewProposal(user, action)
	proposal.PrepareProposal()
	return proposal
}

// ProposeDeleteUser proposes to delete a user from the Forum
func (v *Validator) ProposeDeleteUser(user *User) *Proposal {
	// Create a new proposal to delete the user
	action := "delete_user"
	proposal := NewProposal(user, action)
	proposal.PrepareProposal()
	return proposal
}

// ProposeNewModerator proposes to add a new moderator to the Forum
func (v *Validator) ProposeNewModerator(newModeratorName string) *Proposal {
	// Create a new proposal to add the new moderator
	action := "add_moderator"
	newModerator := &User{Name: newModeratorName}
	proposal := NewProposal(newModerator, action)
	proposal.PrepareProposal()
	return proposal
}

// ProposeBanUser proposes to ban a user from the Forum
func (v *Validator) ProposeBanUser(user *User) *Proposal {
	// Create a new proposal to ban the user
	action := "ban_user"
	proposal := NewProposal(user, action)
	proposal.PrepareProposal()
	return proposal
}

// VoteExtension broadcasts a proposal to all validators for voting
func (v *Validator) VoteExtension(proposal *Proposal, validators []*Validator) {
	for _, validator := range validators {
		if validator.Name != v.Name {
			proposal.Votes[validator.Name] = false
		}
	}
}

// FinalizeBlock finalizes a block and updates the User and Moderator structs accordingly
func (v *Validator) FinalizeBlock(proposal *Proposal, users []*User, moderators []*Moderator) {
	if proposal.Action == "add_user" {
		// Add the user to the User slice
		users = append(users, proposal.User)
	} else if proposal.Action == "delete_user" {
		// Delete the user from the User slice
		for i, user := range users {
			if user.Name == proposal.User.Name {
				users = append(users[:i], users[i+1:]...)
				break
			}
		}
	} else if proposal.Action == "add_moderator" {
		// Add the new moderator to the Moderator slice
		newModerator := &Moderator{Name: proposal.User.Name}
		moderators = append(moderators, newModerator)
	} else if proposal.Action == "ban_user" {
		// Ban the user if their message contains a hardcoded curse word
		if proposal.User.Message == "bad word" {
			proposal.User.IsBanned = true
		}
	}
	fmt.Println("Block finalized. Users:", users, "Moderators:", moderators)
}
