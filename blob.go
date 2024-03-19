package goblob

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/kzg4844"
)

func CreateSidecar(blobs *[]kzg4844.Blob) (*types.BlobTxSidecar, error) {
	var commitments []kzg4844.Commitment
	var proofs []kzg4844.Proof
	for _, blob := range *blobs {
		commitment, err := kzg4844.BlobToCommitment(blob)
		if err != nil {
			return nil, err
		}

		proof, err := kzg4844.ComputeBlobProof(blob, commitment)
		if err != nil {
			return nil, err
		}

		commitments = append(commitments, commitment)
		proofs = append(proofs, proof)

	}
	return &types.BlobTxSidecar{
		Blobs:       *blobs,
		Commitments: commitments,
		Proofs:      proofs,
	}, nil
}
