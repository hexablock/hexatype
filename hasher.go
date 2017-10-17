package hexatype

import (
	"crypto/sha1"
	"encoding/json"
	"hash"

	blake2b "github.com/minio/blake2b-simd"
)

// HashAlgorithm is the hashing algorithm a Hasher implements. e.g. SHA1 SHA256 SHA512 etc.
type HashAlgorithm string

// Hasher interface implements an interface to create a new instance of a hash function
type Hasher interface {
	New() hash.Hash
	Algorithm() HashAlgorithm
	ZeroHash() []byte
	Size() int
}

// SHA1Hasher implements the Hasher interface for SHA1
type SHA1Hasher struct{}

// New returns a new instance of a SHA1 hasher
func (hasher *SHA1Hasher) New() hash.Hash {
	return sha1.New()
}

// Algorithm returns the hashing algorithm the hasher implements
func (hasher *SHA1Hasher) Algorithm() HashAlgorithm {
	return HashAlgorithm("SHA1")
}

// ZeroHash returns a zero hash for SHA1
func (hasher *SHA1Hasher) ZeroHash() []byte {
	return make([]byte, 20)
}

// Size returns the hash size
func (hasher *SHA1Hasher) Size() int {
	return 20
}

// MarshalJSON is a custom json marshaller or the hasher
func (hasher SHA1Hasher) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Algorithm": hasher.Algorithm(),
	})
}

// SHA256Hasher implements the Hasher interface for SHA2
type SHA256Hasher struct{}

// New returns a new instance of a SHA1 hasher
func (hasher *SHA256Hasher) New() hash.Hash {
	return blake2b.New256()
}

// Algorithm returns the hashing algorithm the hasher implements
func (hasher *SHA256Hasher) Algorithm() HashAlgorithm {
	return HashAlgorithm("SHA256")
}

// ZeroHash returns a zero hash for SHA256
func (hasher *SHA256Hasher) ZeroHash() []byte {
	return make([]byte, 32)
}

// Size returns the size of the hash
func (hasher *SHA256Hasher) Size() int {
	return 32
}

// MarshalJSON is a custom json marshaller or the hasher
func (hasher SHA256Hasher) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"Algorithm": hasher.Algorithm(),
	})
}
