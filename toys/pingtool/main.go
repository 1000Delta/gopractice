/*
pingtool is a tool to ping a group of ip and return some info about these.
*/
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sync"
)

var (
	pingFileName = "hosts.ping"
	wg           = &sync.WaitGroup{}
)

func main() {
	if !flag.Parsed() {
		flag.PrintDefaults()
	}
	flag.Parse()
	switch flag.NArg() {
	case 0:
	case 1:
		pingFileName = flag.Args()[0]
	default:
		flag.PrintDefaults()
		return
	}
	// set logger
	log.SetPrefix("[pingtool]")
	// read pingFile
	log.Println("Reading hosts from: " + pingFileName)
	hdata, err := ioutil.ReadFile(pingFileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	hosts := bytes.Split(hdata, []byte{'\n'})
	if len(hosts) == 1 {
		if isMatch, _ :=
			regexp.Match(
				HostLineFmtRegexp,
				bytes.Split(hosts[0], []byte{'#'})[0],
			); !isMatch {
			log.Fatalln("the ping file cannot been parse")
		}
	}
	fmt.Printf("Name\t\t\tIP\t\tTTL\t\tTime:\tmin\tavg\tmax\tmdev\n")
	// ping
	for id, hostLine := range hosts {
		hostLine = bytes.TrimSpace(hostLine)
		hostLine = bytes.Split(hostLine, []byte{'#'})[0]
		if len(hostLine) == 0 {
			continue
		}
		wg.Add(1)
		// concurent ping
		go func(line int, hostLine []byte) {
			defer wg.Done()
			host, err := NewHost(string(hostLine))
			if err != nil {
				log.Printf("Cannot parse this line: line %d\n", line)
				return
			}
			info, err := host.Ping()
			if err != nil {
				fmt.Printf("%-10s\t\t%s\n", host.Name, err.Error())
				return
			}
			fmt.Printf(
				"%-10s\t\t%s\t%d\t\t\t%.2f\t%.2f\t%.2f\t%.2f\n",
				host.Name,
				info.IP,
				info.TTL,
				info.Time.Min,
				info.Time.Avg,
				info.Time.Max,
				info.Time.Mdev,
			)
		}(id, hostLine)
	}
	wg.Wait()
}
