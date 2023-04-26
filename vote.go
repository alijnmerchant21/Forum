/* Explanation

This code defines a Vote struct and a VoteSet struct, which are used to represent votes on a proposal in a forum.

* Vote struct: Has two fields:
	voter - which is a string representing the voter's ID, and
	option - which is a string representing the option they are voting
	for (e.g. "yes" or "no").

* NewVote function: Constructor for the Vote struct, which takes two
arguments for voter and option, and returns a pointer to a
new Vote struct with those values.

* VoteSet struct: Has several fields, including height, round,
and step, which represent the height, round, and step of the proposal
being voted on. It also has slices for prevotes and precommit votes,
maps for votedUsers and votedValidator, and other fields to keep
track of the voting results. The NewVoteSet function is a constructor
for the VoteSet struct, which takes arguments for height, round, step,
proposalHash, and voteLimit, and returns a pointer to a new
VoteSet struct with those values.

* AddVote method: Of the VoteSet struct is used to add a new vote to
the prevotes slice if the voting step is "propose", or add it to
the precommit slice if the voting step is "prevote".
It first checks if the voter has already voted, and if so,
returns without adding the vote.

* GetPrevotes and GetPrecommit methods: Of the VoteSet struct are used
to get the slices of prevotes and precommit votes, respectively.

* HasAllVotes method: Of the VoteSet struct is used to check if all
the eligible users have voted, by comparing the length of the
votedUsers map to the length of the prevotes slice.

* VoteOnProposal method: Of the VoteSet struct is used to add a vote
to the votedValidator map and update the yesVotes count
if the vote option is "yes".
It then returns true if the yesVotes count is greater than or equal
to the voteLimit, indicating that the proposal has been approved,
or false otherwise.
*/

package forum

import (
	"bytes"
	//types "github.com/cometbft/cometbft/abci/types"
)

type Vote struct {
	voter  string
	option string
}

func NewVote(voter string, option string) *Vote {
	return &Vote{
		voter:  voter,
		option: option,
	}
}

func (v *Vote) GetVoter() string {
	return v.voter
}

func (v *Vote) GetOption() string {
	return v.option
}

type VoteSet struct {
	height         int64
	round          int
	step           string
	prevotes       []*Vote
	precommit      []*Vote
	votedUsers     map[string]bool
	proposalHash   []byte
	voteLimit      int
	yesVotes       int
	noVotes        int
	votedValidator map[string]bool
}

func NewVoteSet(height int64, round int, step string, proposalHash []byte, voteLimit int) *VoteSet {
	return &VoteSet{
		height:         height,
		round:          round,
		step:           step,
		prevotes:       []*Vote{},
		precommit:      []*Vote{},
		votedUsers:     make(map[string]bool),
		proposalHash:   proposalHash,
		voteLimit:      voteLimit,
		yesVotes:       0,
		noVotes:        0,
		votedValidator: make(map[string]bool),
	}
}

func (vs *VoteSet) AddVote(vote *Vote) {
	if vs.votedUsers[vote.GetVoter()] {
		return
	}
	vs.votedUsers[vote.GetVoter()] = true

	if vs.step == "propose" {
		vs.prevotes = append(vs.prevotes, vote)
	} else if vs.step == "prevote" {
		vs.precommit = append(vs.precommit, vote)
	}
}

func (vs *VoteSet) GetPrevotes() []*Vote {
	return vs.prevotes
}

func (vs *VoteSet) GetPrecommit() []*Vote {
	return vs.precommit
}

func (vs *VoteSet) HasAllVotes() bool {
	return len(vs.votedUsers) == len(vs.prevotes)
}

func (vs *VoteSet) VoteOnProposal(vote *Vote) bool {
	if vs.votedValidator[vote.GetVoter()] {
		return false
	}

	if !bytes.Equal(vs.proposalHash, []byte(vote.GetOption())) {
		return false
	}

	vs.votedValidator[vote.GetVoter()] = true

	if vote.GetOption() == "yes" {
		vs.yesVotes++
	}

	return vs.yesVotes >= vs.voteLimit
}
