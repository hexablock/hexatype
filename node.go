package hexatype

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"net"
	"strings"
	"time"

	"github.com/hexablock/vivaldi"
)

// HashID returns the hash the node address
func (node *Node) HashID(h hash.Hash) []byte {
	h.Write(node.Address)
	sh := h.Sum(nil)
	return sh
}

// Port returns the gossip port
func (node *Node) Port() uint16 {
	return binary.BigEndian.Uint16(node.Address[len(node.Address)-2:])
}

// Host returns the host:port string for the Address
func (node *Node) Host() string {
	return node.IP().String() + fmt.Sprintf(":%d", node.Port())
}

// IP returns the gossip ip address of the node
func (node *Node) IP() net.IP {
	return net.IP(node.Address[:len(node.Address)-2])
}

// Metadata parses the metadata bytes into a map
func (node *Node) Metadata() map[string]string {
	m := make(map[string]string)
	for _, k := range node.Meta {
		p := strings.Split(string(k), "=")
		m[p[0]] = p[1]
	}
	return m
}

// MarshalJSON custom encodes the node to json
func (node *Node) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID          string
		Address     string
		LastSeen    time.Time
		Heartbeats  uint32
		Region      string
		Sector      string
		Zone        string
		Meta        []byte
		Coordinates *vivaldi.Coordinate
	}{
		hex.EncodeToString(node.ID),
		node.Host(),
		time.Unix(0, node.LastSeen),
		node.Heartbeats,
		node.Region,
		node.Sector,
		node.Zone,
		node.Meta,
		node.Coordinates,
	})
}
