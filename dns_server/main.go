package main

import (
	"log"
	"time"

	"github.com/miekg/dns"
)

func handleDNS(w dns.ResponseWriter, req *dns.Msg) {
	client := &dns.Client {
		Net:     "udp",
		Timeout: 3 * time.Second,
	}

	resp, _, err := client.Exchange(req, "1.1.1.1:53")
	if err != nil {
		log.Printf("upstream error: %v", err)

		// Return SERVFAIL to the client.
		fail := new(dns.Msg)
		fail.SetRcode(req, dns.RcodeServerFailure)
		_ = w.WriteMsg(fail)
		return
	}

	// Forward upstream response to the client.
	if err := w.WriteMsg(resp); err != nil {
		log.Printf("response error: %v", err)
	}
}

func main() {
	server := &dns.Server{
		Addr:    ":5353",
		Net:     "udp",
		Handler: dns.HandlerFunc(handleDNS),
	}
	log.Println("DNS server listening on :5353")
	log.Fatal(server.ListenAndServe())
}
