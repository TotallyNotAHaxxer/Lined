package Z

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var counter = 0
var handle *pcap.Handle
var x error

func (Struct *Data) Loader() {
	handle, x = pcap.OpenOffline(Struct.Filename)
	if x != nil {
		log.Fatal(x)
	} else {
		fmt.Println("[!] handler opened")
	}
	defer handle.Close()
	fmt.Println("[+] Link handler closed")
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	fmt.Println("[!] Link opened")
	for packet := range packetSource.Packets() {
		counter++
		if counter == Struct.Line_Num {
			if packet != nil {
				fmt.Println("[+] INFORMATION ::::: FOUND PACKET LINE ")
				Struct.Packet = packet
				fmt.Println("-------------------------------------------------")
				fmt.Println(packet)
				fmt.Println("-------------------------------------------------")
				fmt.Println("[!] Loading injector")
				break
			} else {
				fmt.Println("[!] Could not write packet to file, it seems as if it was not a valid line within the packet")
			}
		}
	}
	Struct.Inject()
}
