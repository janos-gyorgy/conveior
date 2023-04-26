package ip_utils

import (
	"fmt"
	"net"
)

func GetIPFromInterface(interfaceName string) (string, error) {
	var err error
	itf, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return "", err
	}
	item, err := itf.Addrs()
	if err != nil {
		return "", err
	}
	var ip net.IP
	for _, addr := range item {
		switch v := addr.(type) {
		case *net.IPNet:
			//Verify if IP is IPV4
			if v.IP.To4() != nil {
				ip = v.IP
			}
		}
	}
	if ip != nil {
		return ip.String(), nil
	} else {
		return "", fmt.Errorf("no IP")
	}
}
func GetIPFromInterfaces() (map[string]string, error) {
	out := make(map[string]string)
	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {
		ip, err := GetIPFromInterface(i.Name)
		if err == nil {
			out[i.Name] = ip
		}
	}
	return out, nil
}
