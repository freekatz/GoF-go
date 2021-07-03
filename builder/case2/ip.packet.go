package case2

import (
	"bytes"
	"encoding/json"
)

type IPPacket struct {
	header  IPHeader
	payload []byte
}

func (packet *IPPacket) Payload() []byte {
	return packet.payload
}

func (packet *IPPacket) Header() []byte {
	return packet.header.Bytes()
}

func (packet *IPPacket) Bytes() []byte {
	return bytes.Join([][]byte{packet.header.Bytes(), packet.payload}, []byte{})
}

func (packet *IPPacket) Copy() IPPacket {
	copyPacket := *packet
	return copyPacket
}

type IPVersion int

const (
	IPv4 IPVersion = 4
	IPv6 IPVersion = 6
)

type IPHeader struct {
	Version IPVersion `json:"version"`
	SrcAddr string    `json:"srcAddr"`
	DstAddr string    `json:"dstAddr"`
}

func (header *IPHeader) Bytes() []byte {
	headerAsBytes, _ := json.Marshal(header)
	return headerAsBytes
}

func (header *IPHeader) Copy() IPHeader {
	copyHeader := *header
	return copyHeader
}
