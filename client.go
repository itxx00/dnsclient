package main

import (
    "github.com/miekg/dns"
    "net"
    "os"
    "log"
    "time"
    "fmt"
)

func main() {
    c := new(dns.Client)
    m := new(dns.Msg)
    server := os.Args[1]
    domain := os.Args[2]
    m.SetQuestion(dns.Fqdn(domain), dns.TypeA)
    m.RecursionDesired = false

    var port string = "53"
    start := time.Now()
    r, _, err := c.Exchange(m, net.JoinHostPort(server, port))
    end := time.Now()
    delta := end.Sub(start)
    if r == nil {
        log.Fatalf("*** error: %s\n", err.Error())
    }

    if r.Rcode != dns.RcodeSuccess {
        log.Fatalf("*** invalid answer name %s after A query for %s\n", domain, domain)
    }

    for _, a := range r.Answer {
        fmt.Printf("%v %s\n", a, delta)
    }
}
