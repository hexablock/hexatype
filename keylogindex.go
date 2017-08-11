package hexatype

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"log"
)

// NewKeylogIndex creates a new keylog index
func NewKeylogIndex(key, locID []byte) *KeylogIndex {
	return &KeylogIndex{Key: key, Location: locID, Entries: [][]byte{}}
}

// Append appends the id to the index making sure prev matches the id of the current
// last entry.  If this is the first entry, prev should be a zero hash.
func (idx *KeylogIndex) Append(id, prev []byte) error {
	last := idx.Last()
	if last == nil {
		if !isZeroBytes(prev) {
			return ErrPreviousHash
		}
	} else if bytes.Compare(last, prev) != 0 {
		return ErrPreviousHash
	}

	idx.Entries = append(idx.Entries, id)
	return nil
}

// Last returns the last entry id or nil if there are no entries.
func (idx *KeylogIndex) Last() []byte {
	l := len(idx.Entries)
	if l == 0 {
		return nil
	}
	return idx.Entries[l-1]
}

// Rollback removes the last entry from the index if it contains entries returning the
// remaining entry count
func (idx *KeylogIndex) Rollback() int {
	l := len(idx.Entries)
	if l > 0 {
		l--
		idx.Entries = idx.Entries[:l]
	}

	return l
}

// Count returns the number of entries in the index
func (idx *KeylogIndex) Count() int {
	return len(idx.Entries)
}

// Iter iterates through each entry id starting fromt he seek position.  If seek is nil
// all entries are traversed.  If the callback returns true the function exits immediately
func (idx *KeylogIndex) Iter(seek []byte, cb func(id []byte) error) (err error) {
	var s int

	// Find the seek position from the index
	if seek != nil {
		s = -1
		for i, e := range idx.Entries {
			if bytes.Compare(seek, e) == 0 {
				s = i
				break
			}
		}
		// Return error if we can't find the seek position
		if s < 0 {
			return ErrEntryNotFound
		}
	}

	log.Printf("[DEBUG] seek=%x pos=%d", seek, s)

	l := idx.Count()
	// Start from seek issueing callbacks
	for i := s; i < l; i++ {
		// Bail if true
		if err = cb(idx.Entries[i]); err != nil {
			break
		}
	}

	return err
}

// MarshalJSON is a custom marshaller to handle encoding byte slices to hex.  We do not
// have an unmarshaller as the index is not directly written to.
func (idx *KeylogIndex) MarshalJSON() ([]byte, error) {

	obj := struct {
		Key      string
		Location string
		Entries  []string
	}{
		Key:      string(idx.Key),
		Location: hex.EncodeToString(idx.Location),
		Entries:  make([]string, len(idx.Entries)),
	}

	for i, e := range idx.Entries {
		obj.Entries[i] = hex.EncodeToString(e)
	}

	return json.Marshal(obj)
}

func isZeroBytes(b []byte) bool {
	for _, v := range b {
		if v != 0 {
			return false
		}
	}
	return true
}
