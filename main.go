package main

import (
	"fmt"
	"os"
	"strconv"

	Modz "main/Mods"
)

var (
	D Modz.Data
)

func is_em(ob ...interface{}) bool {
	if ob != nil {
		fmt.Print("\033[32m")
		return false
	} else {
		fmt.Print("\033[31m THIS VALUE SHOULD NOT BE EMPTY! ---- ")
		fmt.Print("\033[31m")
		return true
	}
}

func Is_emint(ob int) bool {
	if ob == 0x00 {
		fmt.Print("\033[31m THIS VALUE SHOULD NOT BE EMPTY! ---- ")
		return true
	} else {
		fmt.Print("\033[32m")
		return false
	}
}

func main() {
	D.Banner()
	D.Filename = os.Args[1]
	D.Output = os.Args[2]
	D.Line_Num, _ = strconv.Atoi(os.Args[3])
	if os.Args[4] == "" {
		D.Snaplen = 1024
	} else {
		u64, _ := strconv.ParseUint(os.Args[4], 10, 32)
		D.Snaplen = uint32(u64)
	}
	if D.Filename == "" || D.Line_Num == 0 || D.Output == "" {
		fmt.Println("[1] Values :")
		fmt.Println("[1]  nil?     Filename                 - ", is_em(D.Filename), "\033[39m")
		fmt.Println("[1]  nil?     Output File              - ", is_em(D.Output), "\033[39m")
		fmt.Println("[1]  nil?     Line number to pull from - ", Is_emint(D.Line_Num), "\033[39m")
	} else {
		fmt.Println("[+] Calling loader....")
		D.Loader()
	}
}

///home/xea43p3x/Desktop/Projects/frizz/src/PCAP/Pcap_Examples/'HTTP - Basic Authentication.pcap' file.pcap 300
