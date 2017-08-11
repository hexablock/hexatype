package hexatype

import "encoding/json"

// MarshalJSON is a custom marshaller to handle the entry key
func (kvp KeyValuePair) MarshalJSON() ([]byte, error) {
	obj := struct {
		Key   string
		Value []byte
		Entry *Entry
	}{
		Key:   string(kvp.Key),
		Value: kvp.Value,
		Entry: kvp.Entry,
	}
	return json.Marshal(obj)
}
