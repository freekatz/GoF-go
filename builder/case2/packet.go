package case2

type Packet interface { // 定义 packet 方法
	Payload() []byte // 获取下一层数据包, 如 IPv4.Payload() 会返回 Ethernet 的数据包
	Header() []byte  // 输出 header
	Bytes() []byte   // 返回整个数据包
	Copy() Packet    // return a deep copy
}

type Header interface {
	Bytes() []byte
	Copy() Header
}
