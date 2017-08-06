package hexatype

import "errors"

var (
	ErrPreviousHash  = errors.New("previous hash mismatch")
	ErrEntryNotFound = errors.New("entry not found")
	ErrKeyNotFound   = errors.New("key not found")
	ErrKeyExists     = errors.New("key exists")
	ErrKeyInvalid    = errors.New("key invalid")
)
