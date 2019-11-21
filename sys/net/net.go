package net

import (
	"net"
	"regexp"
)

//IPv4List get local ipv4
func IPv4List() ([]net.IP, error) {
	itfs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var (
		itf      net.Interface
		addrs    []net.Addr
		addr     net.Addr
		ipNet    *net.IPNet
		ok       bool
		ipv4     net.IP
		ipv4List []net.IP
	)
	for _, itf = range itfs {
		if itf.Flags&net.FlagUp == 0 {
			continue
		}
		addrs, err = itf.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr = range addrs {
			ipNet, ok = addr.(*net.IPNet)
			if !ok || ipNet.IP.IsLoopback() {
				continue
			}
			ipv4 = ipNet.IP.To4()
			if ipv4 == nil {
				continue
			}
			ipv4List = append(ipv4List, ipv4)
		}
	}
	return ipv4List, nil
}

//GetMacAddrs return local macaddrs
func GetMacAddrs() map[string]string {
	mapMacs := make(map[string]string)
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return mapMacs
	}
	for _, netInterface := range netInterfaces {
		addr := netInterface.HardwareAddr.String()
		if len(addr) == 0 {
			continue
		}
		mapMacs[netInterface.Name] = netInterface.HardwareAddr.String()
	}
	return mapMacs
}

const macRule = "([A-Fa-f0-9]{2}[-,:]){5}[A-Fa-f0-9]{2}"

//CheckMacaddr check the validity of macaddr
func CheckMacaddr(mac string) bool {
	match, err := regexp.MatchString(macRule, mac)
	if err != nil {
		return false
	}
	return match
}
