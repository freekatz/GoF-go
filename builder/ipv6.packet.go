package builder

type IPv6Packet struct {
	IPPacket

	header IPv6Header
}

type IPv6Header struct {
	IPHeader
	// IPv6 特有的
	HopLimit int32 `json:"hopLimit"`
}
