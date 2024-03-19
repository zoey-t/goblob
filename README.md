# goblob
utils for generating ethereum blob transacions

function `CreateSidecarAndVersionedHashes` creates the `sidecar` and `BlobHashes` fields in the blob transaction.
`sidecar` struct contains the blobs, kzg commitments and kzg proofs.
The data that needs to send to etheruem in the blob transaction will be transformed to blobs first.
The kzg commiment and proof are generated for each blob.

Blobhashes contains the unique versioned hashes for each blob attached in the blob transaction. It's the sha256 hash of a version prefix and blob's kzg commitment.

function `CreateBlobTx` create the blob transaction with the given data.