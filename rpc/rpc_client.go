package rpc

import (
	"eth2-exporter/types"

	"github.com/sirupsen/logrus"
)

// Client provides an interface for RPC clients
type Client interface {
	GetChainHead() (*types.ChainHead, error)
	GetEpochData(epoch uint64) (*types.EpochData, error)
	GetValidatorQueue() (*types.ValidatorQueue, error)
	GetEpochAssignments(epoch uint64) (*types.EpochAssignments, error)
	GetBlocksBySlot(slot uint64) ([]*types.Block, error)
	GetValidatorParticipation(epoch uint64) (*types.ValidatorParticipation, error)
	GetNewBlockChan() chan *types.Block
	GetBlockStatusByEpoch(slot uint64) ([]*types.CanonBlock, error)
	GetFinalityCheckpoints(epoch uint64) (*types.FinalityCheckpoints, error)
}

var logger = logrus.New().WithField("module", "rpc")
