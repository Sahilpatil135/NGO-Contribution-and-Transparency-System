package blockchain

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"

	"server/internal/blockchain/contracts"
)

func UUIDToBytes16(id uuid.UUID) [16]byte {
	return [16]byte(id)
}

func Bytes16ToUUID(bytes [16]byte) uuid.UUID {
	return uuid.UUID(bytes)
}

type DonationChainService struct {
	client   *Client
	contract *contracts.DonationLedger
	address  common.Address
}

func NewDonationChainService(
	client *Client,
	contractAddress string,
) (*DonationChainService, error) {

	addr := common.HexToAddress(contractAddress)

	instance, err := contracts.NewDonationLedger(addr, client.EthClient)
	if err != nil {
		return nil, err
	}

	return &DonationChainService{
		client:   client,
		contract: instance,
		address:  addr,
	}, nil
}

func (s *DonationChainService) RecordDonation(
	ctx context.Context,
	donationID uuid.UUID,
	causeID uuid.UUID,
	donorID uuid.UUID,
	amount *big.Int,
	paymentRef string,
) (string, error) {

	auth, err := s.client.NewTransactor(ctx)
	if err != nil {
		return "", err
	}

	tx, err := s.contract.RecordDonation(
		auth,
		UUIDToBytes16(donationID),
		UUIDToBytes16(causeID),
		UUIDToBytes16(donorID),
		amount,
		paymentRef,
	)

	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s *DonationChainService) GetDonation(
	ctx context.Context,
	donationID uuid.UUID,
) (*contracts.DonationLedgerDonation, error) {

	donation, err := s.contract.GetDonation(
		nil,
		UUIDToBytes16(donationID),
	)

	if err != nil {
		return nil, err
	}

	return &donation, nil
}

func (s *DonationChainService) GetDonationsByCause(
	ctx context.Context,
	causeID uuid.UUID,
) ([][16]byte, error) {

	donations, err := s.contract.GetDonationsByCause(
		nil,
		UUIDToBytes16(causeID),
	)

	if err != nil {
		return nil, err
	}

	return donations, nil
}

func (s *DonationChainService) VerifyTransaction(
	ctx context.Context,
	txHash string,
) (bool, error) {

	hash := common.HexToHash(txHash)

	receipt, err := s.client.EthClient.TransactionReceipt(ctx, hash)
	if err != nil {
		return false, err
	}

	return receipt.Status == 1, nil
}
