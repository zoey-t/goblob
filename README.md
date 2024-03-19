# goblob
utils for generating Ethereum blob transaction

function `CreateSidecarAndVersionedHashes` creates the `sidecar` and `BlobHashes` fields in the blob transaction.
`sidecar` struct contains the blobs, kzg commitments and kzg proofs.
The data that needs to be sent to Ethereum in the blob transaction will be transformed into blobs first.
The kzg commitment and proof are generated for each blob.

Blobhashes contain the unique versioned hashes for each blob attached in the blob transaction. It's the sha256 hash of a version prefix and blob's kzg commitment.

function `CreateBlobTx` creates the blob transaction with the given data.

example: a blob on-chain [transaction](https://sepolia.etherscan.io/tx/0x2b780e9fe5aba6272a6059b4b449bd74e667599ae1fe06ae0840c894217a0c0d)
