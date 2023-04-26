# Forum

Forum based on ABCI++

## This is a Forum Application based on ABCI++

Many thanks to the original team for brainstorming and bringing forth this idea. Their original repo can be found [here](https://github.com/interchainio/forum)

**The tutorial demonstrate the use of:**

- Prepare Proposal
- Process Proposal
- Vote Extension
- Finalize Block

### Here is a basic structure of the *Forum Application*

#### We can divide the application structure into the following files/modules:

**main.go** - The entry point of the application that initializes the ABCI application and starts the server.

**app.go** - The main application file that defines the ForumApplication struct and implements the ABCI interface methods.

**validator.go** - This file/module defines the Validator struct and methods for proposing to add or delete a user.

**moderator.go** - This file/module defines the Moderator struct and methods for proposing new moderators and validating the proposal.

**user.go** - This file/module defines the User struct and methods for adding and banning a user.

**proposal.go** - This file/module defines the Proposal struct and methods for preparing and processing proposals.

**vote.go** - This file/module defines the Vote struct and methods for voting on proposals using the Vote Extension.

### The flow of the application can be summarized as follows

- On initialization, the ForumApplication struct is created with the set of validators defined in the genesis file.

- A validator can propose to add or delete a user by creating a Proposal object and calling the PrepareProposal method of the Proposal module.

- The proposal is then broadcasted to other validators for voting using the Vote Extension.

- The ProcessProposal method of the Proposal module is called once the proposal has received enough votes.

- If the proposal is approved, the User struct is updated to add or delete the user accordingly.

- Validators can propose new moderators by creating a Proposal object and calling the PrepareProposal method of the Proposal module.

- The proposal is then broadcasted to other validators for voting using the Vote Extension.

- The ProcessProposal method of the Proposal module is called once the proposal has received enough votes.

- If the proposal is approved, the Moderator struct is updated to add the new moderator.

- To ban a user, a moderator can create a Proposal object and call the PrepareProposal method of the Proposal module.

- The proposal is then broadcasted to other validators for voting using the Vote Extension.

- The ProcessProposal method of the Proposal module is called once the proposal has received enough votes.

- If the proposal is approved and the proposed user's message contains a hardcoded curse word, the User struct is updated to ban the user.

### The main functions of each module are as follows

1. **PrepareProposal** - This method prepares a proposal object and sets its attributes.

2. **ProcessProposal** - This method processes the proposal object and executes the necessary action if the proposal is approved.

3. **Vote Extension** - This extension allows validators to vote on proposals and reach consensus.

4. **Finalize Block** - This method is called after all transactions have been processed in a block and finalizes the state of the application for the next block.
