package case2

type IPv4PacketBuilder struct {
	IPPacketBuilder
	packet IPv4Packet
}

func NewIPv4PacketBuilder() *IPv4PacketBuilder {
	return &IPv4PacketBuilder{
		packet: IPv4Packet{},
	}
}

func (builder *IPv4PacketBuilder) MakeTTL(ttl int32) *IPv4PacketBuilder {
	builder.packet.header.TTL = ttl

	return builder
}
