package packet

import (
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/poly1305"
	"golang.zx2c4.com/wireguard/tai64n"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var (
	InitialChainKey [blake2s.Size]byte
	InitialHash     [blake2s.Size]byte
	ZeroNonce       [chacha20poly1305.NonceSize]byte
)

func init() {
	InitialChainKey = blake2s.Sum256([]byte(NoiseConstruction))
	mixHash(&InitialHash, &InitialChainKey, []byte(WGIdentifier))
}

type MessageInitiation struct {
	Type      uint32
	Sender    uint32
	Ephemeral NoisePublicKey
	Static    [NoisePublicKeySize + poly1305.TagSize]byte
	Timestamp [tai64n.TimestampSize + poly1305.TagSize]byte
	MAC1      [blake2s.Size128]byte
	MAC2      [blake2s.Size128]byte
}

type MetricMessage struct {
	Type      uint32
	ID        uint32
	Sender    wgtypes.Key
	Reciever  wgtypes.Key
	TimeStamp int64
}

type ProxyMessage struct {
	Type     uint32
	Sender   [16]byte
	Reciever [16]byte
}
