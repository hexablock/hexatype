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

// NewNode returns a new node with the given ip and port
func NewNode(addr string, port int) *Node {
	ip := net.ParseIP(addr)
	pb := make([]byte, 2)
	binary.BigEndian.PutUint16(pb, uint16(port))

	return &Node{Address: append(ip, pb...)}
}

// LamportTime returns the nodes virtual cluster time
func (node *Node) LamportTime() LamportTime {
	return LamportTime(node.LTime)
}

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

// SetMetadata sets the key value to the meta bytes.  It does not check for
// repeats of keys
func (node *Node) SetMetadata(key, value string) {
	str := strings.Join([]string{key, value}, "=")
	if node.Meta == nil || len(node.Meta) == 0 {
		node.Meta = []byte(str)
	} else {
		node.Meta = append(node.Meta, append([]byte(","), []byte(str)...)...)
	}
}

// Metadata parses the metadata bytes into a map.  In case of duplicate keys the
// last key will be used
func (node *Node) Metadata() map[string]string {
	m := make(map[string]string)
	if node.Meta == nil || len(node.Meta) == 0 {
		return m
	}

	lines := strings.Split(string(node.Meta), ",")

	for _, k := range lines {
		if p := strings.Split(k, "="); len(p) == 2 {
			m[p[0]] = p[1]
		}
	}
	return m
}

// MarshalJSON custom encodes the node to json
func (node *Node) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID          string
		Address     string
		LastSeen    time.Time
		LTime       uint64
		Heartbeats  uint32
		Latency     int64
		Region      string
		Sector      string
		Zone        string
		Meta        []byte
		Coordinates *vivaldi.Coordinate
	}{
		hex.EncodeToString(node.ID),
		node.Host(),
		time.Unix(0, node.LastSeen),
		node.LTime,
		node.Heartbeats,
		node.Latency,
		node.Region,
		node.Sector,
		node.Zone,
		node.Meta,
		node.Coordinates,
	})
}
