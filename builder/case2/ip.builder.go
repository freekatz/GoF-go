package case2

type IPPacketBuilder struct {
	packet IPPacket
}

func (builder *IPPacketBuilder) Build() IPPacket {
	return builder.packet.Copy()
}

func (builder *IPPacketBuilder) BuildHeader() IPHeader {
	return builder.packet.header.Copy()
}

func (builder *IPPacketBuilder) GetBuilder() *IPPacketBuilder {
	return builder
}

func (builder *IPPacketBuilder) MakeVersion(version IPVersion) *IPPacketBuilder {
	builder.packet.header.Version = version

	return builder
}

func (builder *IPPacketBuilder) MakeSrcAddr(addr string) *IPPacketBuilder {
	builder.packet.header.SrcAddr = addr

	return builder
}

func (builder *IPPacketBuilder) MakeDstAddr(addr string) *IPPacketBuilder {
	builder.packet.header.DstAddr = addr
	return builder
}

func (builder *IPPacketBuilder) MakePayload(payload []byte) *IPPacketBuilder {
	builder.packet.payload = payload

	return builder
}
