package case2

import (
	"bytes"
	"encoding/json"
)

type EthernetPacket struct {
	header   EthernetHeader
	ipPacket IPPacket
}

func (packet *EthernetPacket) Payload() []byte {
	return packet.ipPacket.Bytes()
}

func (packet *EthernetPacket) Header() []byte {
	return packet.header.Bytes()
}

func (packet *EthernetPacket) Bytes() []byte {
	return bytes.Join([][]byte{packet.header.Bytes(), packet.ipPacket.Bytes()}, []byte{})
}

func (packet *EthernetPacket) Copy() EthernetPacket {
	copyPacket := *packet
	return copyPacket
}

type EthernetHeader struct {
	SrcMac string `json:"srcMac"`
	DstMac string `json:"dstMac"`
}

func (header *EthernetHeader) Bytes() []byte {
	headerAsBytes, _ := json.Marshal(header)
	return headerAsBytes
}

func (header *EthernetHeader) Copy() EthernetHeader {
	copyHeader := *header
	return copyHeader
}
