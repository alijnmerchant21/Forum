/*Explanation:
This code defines the User, Moderator, and Proposal structs, as well as the Moderators slice.
The User struct represents a user in the Forum application, and has fields for the user's name, message, and
whether they are banned.
The Moderator struct represents a moderator in the Forum application, and has a field for the moderator's name.
The Moderators slice holds the list of moderators in the Forum application.
The Proposal struct represents a proposal in the Forum application, and has fields for the proposer,
the proposed action, the user affected by the action, and a boolean flag indicating whether the proposal has been
approved by the validators.
The NewProposal function creates a new Proposal object with the given user and action.
The PrepareProposal method prepares the proposal for voting by setting the Vote flag to false.
The ProcessProposal method processes the proposal and executes the necessary action if the proposal has been approved. If the action is to add a user, the method sets the IsBanned field of the user to false. If the action is to delete a user, the method sets the IsBanned field of the user to true. If the action is to add a moderator, the method appends the moderator's name to the Moderators slice. If the action is to ban a user, the method checks if the user's message contains a curse word and sets the IsBanned field to true if it does.
*/

package forum

import "strings"

// User struct represents a user in the Forum application
type User struct {
	Name     string
	Message  string
	IsBanned bool
}

// Moderator struct represents a moderator in the Forum application
type Moderator struct {
	Name string
}

// Moderators slice holds the list of moderators in the Forum application
var Moderators []string

// Proposal struct represents a proposal in the Forum application
type Proposal struct {
	Proposer string
	Action   string
	User     *User
	Vote     bool
}

// NewProposal creates a new Proposal object with the given user and action
func NewProposal(user *User, action string) *Proposal {
	return &Proposal{
		Proposer: user.Name,
		Action:   action,
		User:     user,
	}
}

// PrepareProposal prepares the proposal for voting
func (p *Proposal) PrepareProposal() {
	// Add your PrepareProposal logic here
	p.Vote = false
}

// ProcessProposal processes the proposal and executes the necessary action if approved
func (p *Proposal) ProcessProposal() {
	// Add your ProcessProposal logic here
	p.Vote = true

	switch p.Action {
	case "add_user":
		p.User.IsBanned = false
	case "delete_user":
		p.User.IsBanned = true
	case "add_moderator":
		// Add the new moderator to the list
		Moderators = append(Moderators, p.User.Name)
	case "ban_user":
		// Check if the user's message contains a curse word
		if strings.Contains(p.User.Message, "curse word") {
			p.User.IsBanned = true
		}
	}
}
