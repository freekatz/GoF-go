package builder

type IPv4Packet struct {
	IPPacket

	header IPv4Header
}

type IPv4Header struct {
	IPHeader
	// IPv4 特有的
	TTL int32 `json:"ttl"`
}
