package forum

import (
	"os"

	"github.com/cometbft/cometbft/app"
	"github.com/cometbft/cometbft/cmd"
	"github.com/cometbft/cometbft/consensus"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cometbft/cometbft/p2p"
	"github.com/cometbft/cometbft/proxy"
	"github.com/cometbft/cometbft/rpc"
	"github.com/cometbft/cometbft/state"
	"github.com/cometbft/cometbft/store"
)

/*
This code is the entry point for a CometBFT application. It initializes various components of
the application such as the consensus engine, state machine, store, P2P layer, RPC layer,
proxy layer, and command-line interface. It then starts each of these components and waits
forever.
*/

func main() {
	logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))

	// Create the main CometBFT app
	application := app.NewCometApp()

	// Create the CometBFT consensus engine
	consensus := consensus.NewCometConsensus()

	// Create the CometBFT state machine
	state := state.NewCometState()

	// Create the CometBFT store
	store := store.NewCometStore()

	// Create the CometBFT P2P layer
	p2p := p2p.NewCometP2P()

	// Create the CometBFT RPC layer
	rpc := rpc.NewCometRPC()

	// Create the CometBFT proxy layer
	proxy := proxy.NewCometProxy()

	// Create the CometBFT command-line interface
	cmd := cmd.NewCometCLI()

	// Start the CometBFT app
	if err := application.Start(); err != nil {
		logger.Error("error starting application", "err", err)
		return
	}

	// Start the CometBFT consensus engine
	if err := consensus.Start(); err != nil {
		logger.Error("error starting consensus engine", "err", err)
		return
	}

	// Start the CometBFT state machine
	if err := state.Start(); err != nil {
		logger.Error("error starting state machine", "err", err)
		return
	}

	// Start the CometBFT store
	if err := store.Start(); err != nil {
		logger.Error("error starting store", "err", err)
		return
	}

	// Start the CometBFT P2P layer
	if err := p2p.Start(); err != nil {
		logger.Error("error starting p2p layer", "err", err)
		return
	}

	// Start the CometBFT RPC layer
	if err := rpc.Start(); err != nil {
		logger.Error("error starting rpc layer", "err", err)
		return
	}

	// Start the CometBFT proxy layer
	if err := proxy.Start(); err != nil {
		logger.Error("error starting proxy layer", "err", err)
		return
	}

	// Start the CometBFT command-line interface
	if err := cmd.Start(); err != nil {
		logger.Error("error starting command-line interface", "err", err)
		return
	}

	// Wait forever
	select {}
}
