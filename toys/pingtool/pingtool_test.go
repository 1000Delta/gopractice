package main

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

const (
	name = "baidu"
	ip   = "baidu.com"
)

var host, _ = NewHost(fmt.Sprintf("%s %s", name, ip))

func TestNewHost(t *testing.T) {
	if host.Name != name || host.Host != ip {
		t.FailNow()
	}
}

func TestParseInfo(t *testing.T) {
	infob := []byte(`PING 107.172.246.38 (107.172.246.38) 56(84) bytes of data.
64 bytes from 107.172.246.38: icmp_seq=1 ttl=47 time=196 ms
64 bytes from 107.172.246.38: icmp_seq=2 ttl=47 time=192 ms
64 bytes from 107.172.246.38: icmp_seq=3 ttl=47 time=192 ms
64 bytes from 107.172.246.38: icmp_seq=4 ttl=47 time=193 ms

--- 107.172.246.38 ping statistics ---
4 packets transmitted, 4 received, 0% packet loss, time 3000ms
rtt min/avg/max/mdev = 192.423/193.701/196.488/1.731 ms`)
	info, err := ParseInfo(string(infob))
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if info.IP == "" ||
		info.TTL == 0 ||
		info.Time.Avg == 0.0 ||
		info.Time.Max == 0.0 ||
		info.Time.Min == 0.0 ||
		info.Time.Mdev == 0.0 {
		t.Fail()
	}
}

func TestPing(t *testing.T) {
	info, err := host.Ping()
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	if info.IP == "" ||
		info.TTL == 0 ||
		info.Time.Avg == 0.0 ||
		info.Time.Max == 0.0 ||
		info.Time.Min == 0.0 ||
		info.Time.Mdev == 0.0 {
		t.Fail()
	}
}

func TestTest(t *testing.T) {
	ptnOfTime := regexp.MustCompile(PingResponseTimeFmtRegexp)
	res := ptnOfTime.FindStringSubmatch("rtt min/avg/max/mdev = 32.712/33.277/34.362/0.657 ms")
	min, err := strconv.ParseFloat(res[1], 64)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	t.Log(min)
	t.Log(PingResponseTTLFmtRegexp)
}
