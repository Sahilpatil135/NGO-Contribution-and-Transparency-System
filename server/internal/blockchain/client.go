package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	common "github.com/ethereum/go-ethereum/common"
	crypto "github.com/ethereum/go-ethereum/crypto"
	ethclient "github.com/ethereum/go-ethereum/ethclient"

	_ "github.com/joho/godotenv/autoload"
)

type Client struct {
	EthClient  *ethclient.Client
	PrivateKey *ecdsa.PrivateKey
	PublicKey  common.Address
	ChainID    *big.Int
}

func NewClient() (*Client, error) {
	var ethernetClientUrl string = os.Getenv("ETH_RPC_URL")
	if ethernetClientUrl == "" {
		ethernetClientUrl = "ws://127.0.0.1:8545"
	}

	// Connect to the EthClient RPC
	client, err := ethclient.Dial(ethernetClientUrl)
	if err != nil {
		// log.Fatalf("Failed to connect to the Eth client: %v", err)
		return nil, fmt.Errorf("Failed to connect to Eth Client: %v", err)
	}
	log.Println("Connected to Hardhat node")

	// Get Private Key from Hardhat ENV
	privateKeyHex := os.Getenv("PRIVATE_ETH_KEY")
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		// log.Fatalf("Failed to convert Private Key to ECSDSA: %v", err)
		return nil, fmt.Errorf("Failed to convert Private Key to ECSDSA: %v", err)
	}

	// Get PubKey and Addr from Private Key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		// log.Fatalf("Failed to Create Public Key ECDSA: %v", err)
		return nil, fmt.Errorf("Failed to Create Public Key ECDSA: %v", err)
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// var chainID big.Int = *big.NewInt(31337)
	chainID, err := client.ChainID(context.Background())

	// Replace with your actual deployed contract address
	// contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS_1"))

	// Load the contract
	// instance, err := donationLedger.NewDonationLedger(contractAddress, client)
	// if err != nil {
	// log.Fatalf("Failed to load the contract: %v", err)
	// }

	// recordDonation(instance, auth)
	// getDonationsByCauseId(instance, auth)

	return &Client{
		EthClient:  client,
		PrivateKey: privateKey,
		PublicKey:  address,
		ChainID:    chainID,
	}, nil
}

func (c *Client) GetNonce(ctx context.Context) (uint64, error) {
	return c.EthClient.PendingNonceAt(ctx, c.PublicKey)
}

func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.EthClient.SuggestGasPrice(ctx)
}

func (c *Client) NewTransactor(ctx context.Context) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(c.PrivateKey, c.ChainID)
	if err != nil {
		return nil, err
	}

	nonce, err := c.GetNonce(ctx)
	if err != nil {
		return nil, err
	}

	gasPrice, err := c.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	return auth, nil
}
