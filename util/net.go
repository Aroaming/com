package util

import (
	"net"
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
