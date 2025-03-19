package types

import "math/big"

type IDType = [32]byte

type Job struct {
	Id                     IDType
	Creator                string
	Task                   []byte
	Type                   string
	Reward                 *big.Int
	NodeVersion            *big.Int
	MaxResults             *big.Int
	MinResultsForConsensus *big.Int
	CreatedAt              *big.Int
}

type JobPlugin interface {
	GetSupportedJobTypes() []string
	ExecuteJob(job *Job) ([]byte, error)
}
