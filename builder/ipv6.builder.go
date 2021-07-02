package builder

type IPv6PacketBuilder struct {
	IPPacketBuilder
	packet IPv6Packet
}

func NewIPv6PacketBuilder() *IPv6PacketBuilder {
	return &IPv6PacketBuilder{
		packet: IPv6Packet{},
	}
}

func (builder *IPv6PacketBuilder) MakeHopLimit(hopLimit int32) *IPv6PacketBuilder {
	builder.packet.header.HopLimit = hopLimit

	return builder
}
