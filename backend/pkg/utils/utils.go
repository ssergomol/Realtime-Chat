package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("Couldn't read request body")
		return
	}

	err = json.Unmarshal([]byte(body), x)

	if err != nil {
		fmt.Printf("Couldn't unmarshal json data from request")
		return
	}
}

func SetBodyInfoMessage(writer http.ResponseWriter, message string) {
	data := make(map[string]string, 1)
	data["message"] = string(message)
	res, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Couldn't marshal information message\n")
	}
	writer.Write(res)
}

func ExternalIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String()
		}
	}
	panic("Not connected to Network")
}
