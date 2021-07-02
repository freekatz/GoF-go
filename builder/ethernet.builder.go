package builder

type EthernetPacketBuilder struct {
	packet EthernetPacket
}

func NewEthernetPacketBuilder() *EthernetPacketBuilder {
	return &EthernetPacketBuilder{
		packet: EthernetPacket{},
	}
}

func (builder *EthernetPacketBuilder) Build() EthernetPacket {
	return builder.packet.Copy()
}

func (builder *EthernetPacketBuilder) BuildHeader() EthernetHeader {
	return builder.packet.header.Copy()
}

func (builder *EthernetPacketBuilder) GetBuilder() *EthernetPacketBuilder {
	return builder
}

func (builder *EthernetPacketBuilder) MakeSrcMac(mac string) *EthernetPacketBuilder {
	builder.packet.header.SrcMac = mac

	return builder
}

func (builder *EthernetPacketBuilder) MakeDstMac(mac string) *EthernetPacketBuilder {
	builder.packet.header.DstMac = mac

	return builder
}

func (builder *EthernetPacketBuilder) MakePayload(ipPacket IPPacket) *EthernetPacketBuilder {
	builder.packet.ipPacket = ipPacket

	return builder
}

func (builder *EthernetPacketBuilder) MakePayloadWithBuilder(ipPacketBuilder *IPPacketBuilder) *EthernetPacketBuilder {
	builder.packet.ipPacket = ipPacketBuilder.Build()

	return builder
}
