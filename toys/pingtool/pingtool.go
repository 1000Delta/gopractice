package main

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
)

const (
	// DomainRegexp match domain include localhost like
	DomainRegexp = `(?:(?:\w+\.){0,3}\w+)`
	// IPRegexp match ip address only IPv4
	IPRegexp = `(?:(?:\d{1,3}\.){3}(?:\d{1,3}))`
	// HostFmtRegexp match ip or domain include localhost like
	HostFmtRegexp = `(?:` + DomainRegexp + `|` + IPRegexp + `)`
	// HostLineFmtRegexp is the regexp of host (IP or domain)
	HostLineFmtRegexp = `(.+)[ \t]+(` + HostFmtRegexp + `)`
	// PingResponseIPFmtRegexp is the regexp of ip in ping result
	PingResponseIPFmtRegexp = `PING ` + HostFmtRegexp + ` \((` + IPRegexp + `)\) `
	// PingResponseTTLFmtRegexp is the regexp of TTL in ping result
	PingResponseTTLFmtRegexp = `\d+ bytes from (?:(?:` + HostFmtRegexp + ` \(` + IPRegexp + `\))|` + IPRegexp + `): icmp_seq=\d+ ttl=(\d+) time=\d+\.?\d* ms`
	// PingResponseTimeFmtRegexp is the regexp of time (4 info) in ping result
	PingResponseTimeFmtRegexp = `rtt min\/avg\/max\/mdev = (\d+\.\d+)\/(\d+\.\d+)\/(\d+\.\d+)\/(\d+\.\d+) ms`
)

var (
	// Hosts is the group of host you want to ping
	Hosts []Host

	// ErrParse means parse string failed
	ErrParse = errors.New("Parse failed")
)

// Host is the structural data of host you defined in the .ping file
type Host struct {
	Name string
	Host string
}

// PingInfo record the result of ping
type PingInfo struct {
	IP   string
	TTL  int64
	Time *pingInfoTime
}

type pingInfoTime struct {
	Min  float64
	Avg  float64
	Max  float64
	Mdev float64
}

// NewHost parse bytes to Host object
func NewHost(hostLine string) (*Host, error) {
	ptn := regexp.MustCompile(HostLineFmtRegexp)
	data := ptn.FindStringSubmatch(hostLine)
	if len(data) != 3 {
		return nil, errors.New("Parse failed")
	}

	return &Host{data[1], data[2]}, nil
}

// ParseInfo parse ping result bytes to *Info
func ParseInfo(info string) (*PingInfo, error) {
	pingInfo := &PingInfo{}
	// ip
	ptnOfIP := regexp.MustCompile(PingResponseIPFmtRegexp)
	res := ptnOfIP.FindStringSubmatch(info)
	if len(res) != 2 {
		return nil, ErrParse
	}
	pingInfo.IP = res[1]
	// ttl
	ptnOfTTL := regexp.MustCompile(PingResponseTTLFmtRegexp)
	res = ptnOfTTL.FindStringSubmatch(info)
	if len(res) != 2 {
		return nil, ErrParse
	}
	ttl, _ := strconv.ParseInt(res[1], 10, 32) // regexp can provide no error
	pingInfo.TTL = ttl
	// time
	ptnOfTime := regexp.MustCompile(PingResponseTimeFmtRegexp)
	res = ptnOfTime.FindStringSubmatch(info)
	if len(res) != 5 {
		return nil, ErrParse
	}
	pingInfo.Time = &pingInfoTime{}
	pingInfo.Time.Min, _ = strconv.ParseFloat(res[1], 64)
	pingInfo.Time.Avg, _ = strconv.ParseFloat(res[2], 64)
	pingInfo.Time.Max, _ = strconv.ParseFloat(res[3], 64)
	pingInfo.Time.Mdev, _ = strconv.ParseFloat(res[4], 64)

	return pingInfo, nil
}

// Ping run linux ping command and return result as *Info
func (h *Host) Ping() (*PingInfo, error) {
	// run cmd
	pingCmd := exec.Command("ping", h.Host, "-c5")
	// get info
	infoBytes, err := pingCmd.Output()
	if err != nil {
		return nil, err
	}
	info, err := ParseInfo(string(infoBytes))
	if err != nil {
		return nil, err
	}
	return info, nil
}
