package Z

import "github.com/google/gopacket"

type Data struct {
	Filename string
	Output   string
	Line_Num int
	Snaplen  uint32
	Packet   gopacket.Packet
}
