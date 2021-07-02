package builder

type PacketBuilder interface {
	Build() Packet
	BuildHeader() Header
	GetBuilder() *PacketBuilder
}
