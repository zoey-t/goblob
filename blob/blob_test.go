package blob

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func TestSendBlobTx(t *testing.T) {
	rpc := sepolia_rpc()

	privateKeyHex := privateKey()
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		t.Fatal(err)
	}
	client, err := ethclient.Dial(rpc)
	if err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile("test_data.json")
	if err != nil {
		t.Fatal(err)
	}

	blobTx, err := CreateBlobTx(client, privateKeyHex, data)
	if err != nil {
		t.Fatal(err)
	}

	tx := types.NewTx(blobTx)

	// send tx
	signedTx, err := types.SignTx(tx, types.NewCancunSigner(blobTx.ChainID.ToBig()), privateKey)
	if err != nil {
		t.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("nonce", blobTx.Nonce)
	t.Log(signedTx.Hash())

}

func privateKey() string {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("PRIVATE_KEY")

}

func sepolia_rpc() string {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv("SEPOLIA_RPC")

}
