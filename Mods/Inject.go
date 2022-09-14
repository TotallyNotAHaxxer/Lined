package Z

import (
	"fmt"
	"log"
	"os"

	"github.com/google/gopacket/layers"

	"github.com/google/gopacket/pcapgo"
)

var (
	X error
	F *os.File
)

func (Struct *Data) Inject() {
	F, X := os.Open(Struct.Output)
	if X != nil {
		F, X = os.Create(Struct.Output)
		fmt.Println("[!] File was not found so ceating new one....")
		if X != nil {
			log.Fatal(X)
		}
	} else {
		fmt.Println("[+] File found")
	}
	w := pcapgo.NewWriter(F)
	fmt.Println("[+] Writing header.....")
	w.WriteFileHeader(Struct.Snaplen, layers.LinkTypeEthernet)
	defer F.Close()
	w.WritePacket(Struct.Packet.Metadata().CaptureInfo, Struct.Packet.Data())
	fmt.Println("[+] Data saved to -> ", Struct.Output)
}
