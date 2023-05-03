/*
Explanation:
The User struct represents a user in the forum application and has three fields: Name (string), Message (string), and
IsBanned (bool).
The Moderator struct represents a moderator in the forum application and has one field: Name (string).
The Proposal struct represents a proposal in the forum application and has four fields: Proposer (string), Action (string),
User (*User), and Votes (map[string]bool).

The package also defines several functions:
NewProposal creates a new Proposal object with the given user and action.
ProposeNewModerator creates a new proposal to add a new moderator and returns the Proposal object.
ValidateProposal validates the given Proposal object to add a new moderator. The method takes a Moderator object and
a slice of Moderator objects as inputs. The method returns a boolean value indicating whether the proposal is valid.
The ValidateProposal method checks whether the given proposal action is "add_moderator". If not, the proposal is considered invalid
and the method returns false. It then calculates the minimum number of moderators required to approve the proposal based on
the number of moderators in the input slice. The method counts the number of moderators who have approved the proposal by iterating
through the input slice and checking the Votes field of the Proposal object. The method ignores the current moderator
(the one calling the method) when counting the votes. Finally, the method returns true if the number of approving moderators
is greater than or equal to the required number, and false otherwise.
*/

package forum

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

// Proposal struct represents a proposal in the Forum application
type Proposal struct {
	Proposer string
	Action   string
	User     *User
	Votes    map[string]bool
}

// NewProposal creates a new Proposal object with the given user and action
func NewProposal(user *User, action string) *Proposal {
	return &Proposal{
		Proposer: user.Name,
		Action:   action,
		User:     user,
		Votes:    make(map[string]bool),
	}
}

// ProposeNewModerator creates a new proposal to add a new moderator
func (m *Moderator) ProposeNewModerator(newModeratorName string) *Proposal {
	// Create a new user object for the new moderator
	newModerator := &User{
		Name: newModeratorName,
	}
	// Set the action to add moderator and create a new proposal
	action := "add_moderator"
	proposal := NewProposal(newModerator, action)
	// Prepare the proposal for voting
	proposal.PrepareProposal()
	return proposal
}

// ValidateProposal validates the given proposal to add a new moderator
func (m *Moderator) ValidateProposal(proposal *Proposal, moderators []Moderator) bool {
	// Check if the action is to add a new moderator
	if proposal.Action != "add_moderator" {
		return false
	}
	// Calculate the number of moderators required to approve the proposal
	required := (2 * len(moderators) / 3) + 1
	// Count the number of moderators who have approved the proposal
	count := 0
	for _, moderator := range moderators {
		// Ignore the current moderator (the one calling the method)
		if moderator.Name != m.Name {
			// Check if the moderator has voted in favor of the proposal
			if proposal.Votes[moderator.Name] {
				count++
			}
		}
	}
	// Return true if the required number of moderators have approved the proposal
	return count >= required
}
