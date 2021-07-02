package builder

import (
	"log"
	"testing"
)

func Test(t *testing.T) {
	ethernetPacketBuilder := NewEthernetPacketBuilder()
	ipv4PacketBuilder := NewIPv4PacketBuilder()
	ipv6PacketBuilder := NewIPv6PacketBuilder()

	// test ip builder
	ipv4PacketBuilder.MakeVersion(IPv4)
	ipv4PacketBuilder.MakeSrcAddr("1.1.1.1")
	ipv4PacketBuilder.MakeDstAddr("255.255.255.255")
	ipv4PacketBuilder.MakeTTL(255)
	ipv4PacketBuilder.MakePayload([]byte("ipv4 packet"))

	ipv6PacketBuilder.MakeVersion(IPv6)
	ipv6PacketBuilder.MakeSrcAddr("11::11")
	ipv6PacketBuilder.MakeDstAddr("ff::ff")
	ipv6PacketBuilder.MakeHopLimit(1024)
	ipv6PacketBuilder.MakePayload([]byte("ipv6 packet"))

	// test ethernet builder

	ethernetPacketBuilder.MakeSrcMac("0-0-0-0-0-0").MakeDstMac("f-f-f-f-f-f")
	ethernetPacketBuilder.MakePayload(ipv4PacketBuilder.Build())

	ethernetv4Packet := ethernetPacketBuilder.Build()
	ethernetv4Header := ethernetPacketBuilder.BuildHeader()

	log.Println(ethernetv4Packet, ethernetv4Header)

	ethernetPacketBuilder.MakePayloadWithBuilder(ipv6PacketBuilder.GetBuilder())

	ethernetv6Packet := ethernetPacketBuilder.Build()
	ethernetv6Header := ethernetPacketBuilder.BuildHeader()

	log.Println(ethernetv6Packet, ethernetv6Header)

	if &ethernetv4Packet == &ethernetv6Packet || &ethernetv4Header == &ethernetv6Header {
		t.Errorf("build same objects.")
	}

	log.Println(string(ethernetv4Packet.Bytes()))
	log.Println(string(ethernetv4Packet.Header()))
	log.Println(string(ethernetv4Packet.Payload()))
	log.Println(string(ethernetv6Packet.ipPacket.Bytes()))
	log.Println(string(ethernetv6Packet.ipPacket.Header()))
	log.Println(string(ethernetv6Packet.ipPacket.Payload()))
	ipp := ethernetv6Packet.ipPacket.Copy()
	log.Println(&(ethernetv6Packet.ipPacket) == &ipp)
}
