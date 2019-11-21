package util

import (
	"net"
	"regexp"
)

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

func CheckMacaddr(mac string) bool {
	match, err := regexp.MatchString(macRule, mac)
	if err != nil {
		return false
	}
	return match
}
