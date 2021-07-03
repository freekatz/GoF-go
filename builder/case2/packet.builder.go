package case2

type PacketBuilder interface {
	Build() Packet
	BuildHeader() Header
	GetBuilder() *PacketBuilder
}
