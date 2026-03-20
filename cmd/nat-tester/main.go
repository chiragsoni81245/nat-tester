package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/chiragsoni81245/nat-tester/internal/detector"
	"github.com/chiragsoni81245/nat-tester/internal/stunclient"
	"github.com/chiragsoni81245/nat-tester/internal/types"
)

func main() {
	serversFlag := flag.String("stun", "stun.l.google.com:19302,stun1.l.google.com:19302", "Comma-separated STUN servers")
	timeoutFlag := flag.Int("timeout", 3, "Timeout in seconds")
	verbose := flag.Bool("v", false, "Verbose output")

	flag.Parse()

	servers := strings.Split(*serversFlag, ",")

	conn, err := net.ListenUDP("udp", nil)
	if err != nil {
		log.Fatal("Failed to open UDP socket:", err)
	}
	defer conn.Close()

	fmt.Println("Local address:", conn.LocalAddr())
	fmt.Println("Using STUN servers:", servers)

	results := make([]types.Result, 0)
	timeout := time.Duration(*timeoutFlag) * time.Second

	fmt.Println("\n[+] Running STUN tests...")

	for _, server := range servers {
		addr, err := stunclient.Query(conn, server, timeout)
		if err != nil {
			fmt.Println("[-] Failed:", server, err)
			continue
		}

		fmt.Printf("[+] %s → %s\n", server, addr.String())

		results = append(results, types.Result{
			Server: server,
			Addr:   addr,
		})
	}

	if len(results) < 2 {
		fmt.Println("\n[-] Not enough results to determine NAT type")
		return
	}

	natType := detector.Detect(results)

	fmt.Println("\n========== RESULT ==========")
	fmt.Println("NAT Type:", natType)

	if *verbose {
		fmt.Println("\n[DEBUG] Raw Results:")
		for _, r := range results {
			fmt.Printf("- %s → %s\n", r.Server, r.Addr)
		}
	}
}
