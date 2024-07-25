package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var (
	errNoGateway        = errors.New("no gateway found")
	errCantParseGateway = errors.New("can't parse gateway")
	errInvalidMac       = errors.New("invalid mac address")
	errNoMac            = errors.New("no mac address provided")
	errCantParseIP      = errors.New("can't parse ip address")
	errNotFound         = errors.New("mac address not found")
)

func main() {
	if len(os.Args) < 2 {
		panic(errNoMac)
	}

	mac, err := net.ParseMAC(os.Args[1])
	if err != nil {
		panic(errInvalidMac)
	}

	ip, err := discoverGateway()
	if err != nil {
		panic(err)
	}

	nmap, err := nmap(ip)
	if err != nil {
		panic(err)
	}

	ip, err = parseIP(nmap, mac)
	if err != nil {
		panic(err)
	}

	fmt.Printf("IP: %s\n", ip)
}

func discoverGateway() (net.IP, error) {
	routeCmd := exec.Command("route", "print", "0.0.0.0")
	routeCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := routeCmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	ip, err := parseGateway(string(output))
	if err != nil {
		return nil, err
	}

	return ip, nil
}

func parseGateway(route string) (net.IP, error) {
	lines := strings.Split(route, "\n")
	sep := 0
	for idx, line := range lines {
		if sep == 3 {
			if len(lines) <= idx+2 {
				return nil, errNoGateway
			}

			fields := strings.Fields(lines[idx+2])
			if len(fields) < 5 {
				return nil, errCantParseGateway
			}

			ip := net.ParseIP(fields[2])
			if ip == nil {
				return nil, errCantParseGateway
			}

			return ip, nil
		}
		if strings.HasPrefix(line, "=======") {
			sep++
			continue
		}
	}

	return nil, errNoGateway
}

func nmap(ip net.IP) (string, error) {
	nmapCmd := exec.Command("nmap", "-sn", ip.String()+"/24")
	nmapCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := nmapCmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func parseIP(nmap string, mac net.HardwareAddr) (net.IP, error) {
	lines := strings.Split(nmap, "\n")
	for i, line := range lines {
		if strings.Contains(strings.ToLower(line), strings.ToLower(mac.String())) {
			if len(lines) < 2 {
				return nil, errCantParseIP
			}

			fields := strings.Fields(lines[i-2])
			if len(fields) < 5 {
				return nil, errCantParseIP
			}

			ip := net.ParseIP(fields[4])
			if ip == nil {
				return nil, errCantParseIP
			}

			return ip, nil
		}
	}

	return nil, errNotFound
}
