package hexatype

import "errors"

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
)
