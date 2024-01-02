package details

import (
	"net"
	"os"
)

func GetHostName() (string, error) {
	return os.Hostname()
}

func GetLocalIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	addr := conn.LocalAddr().(*net.UDPAddr)
	return addr.IP.String()
}
