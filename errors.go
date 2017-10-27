package hexatype

import (
	"errors"

	"google.golang.org/grpc"
)

// This file contains errors across various hexablock products.  They are kept
// here for implemented interfaces to use.

var (
	// ErrPreviousHash is used when the previous hash of an entry does not match
	ErrPreviousHash = errors.New("previous hash mismatch")
	// ErrEntryNotFound is used when an entry is not found
	ErrEntryNotFound = errors.New("entry not found")
	// ErrKeyNotFound is used when a key is not found
	ErrKeyNotFound = errors.New("key not found")
	// ErrKeyExists is used if the key already exists
	ErrKeyExists = errors.New("key exists")
	// ErrKeyInvalid gets used when an invalid key is seen
	ErrKeyInvalid = errors.New("key invalid")
	// ErrKeyDegraded is used when a key contains a marker and a write is performed
	ErrKeyDegraded = errors.New("key degraded")
	// ErrInsufficientPeers is used when there aren't enough peers for the request
	ErrInsufficientPeers = errors.New("insufficient peers")
)

// ParseGRPCError parses a grpc error into a hexatype error
func ParseGRPCError(e error) error {
	var (
		str = grpc.ErrorDesc(e)
		err error
	)

	switch str {
	case ErrPreviousHash.Error():
		err = ErrPreviousHash

	case ErrEntryNotFound.Error():
		err = ErrEntryNotFound

	case ErrKeyNotFound.Error():
		err = ErrKeyNotFound

	case ErrKeyExists.Error():
		err = ErrKeyExists

	case ErrKeyDegraded.Error():
		err = ErrKeyDegraded

	case ErrKeyInvalid.Error():
		err = ErrKeyInvalid

	default:
		err = e

	}

	return err
}
