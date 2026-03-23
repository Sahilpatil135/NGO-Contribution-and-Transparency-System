package services

import (
	"context"
	"io"
	"os"
	"time"

	ipfsapi "github.com/ipfs/go-ipfs-api"
)

type IPFSService interface {
	AddFile(ctx context.Context, r io.Reader) (string, error)
	Cat(ctx context.Context, cid string) (io.ReadCloser, error)
}

type ipfsService struct {
	sh *ipfsapi.Shell
}

func NewIPFSService() IPFSService {
	// Default matches the common go-ipfs-api usage: "localhost:5001".
	addr := os.Getenv("IPFS_API_ADDR")
	if addr == "" {
		addr = "localhost:5001"
	}

	sh := ipfsapi.NewShell(addr)

	// Reasonable default timeout for remote calls.
	sh.SetTimeout(60 * time.Second)

	return &ipfsService{sh: sh}
}

func (s *ipfsService) AddFile(ctx context.Context, r io.Reader) (string, error) {
	// go-ipfs-api's Add doesn't accept ctx directly; it uses context.Background internally.
	// We keep ctx in the signature for future consistency.
	_ = ctx
	return s.sh.Add(r, ipfsapi.Pin(true))
}

func (s *ipfsService) Cat(ctx context.Context, cid string) (io.ReadCloser, error) {
	_ = ctx
	// go-ipfs-api's Cat doesn't accept ctx directly; it uses context.Background internally.
	return s.sh.Cat(cid)
}

