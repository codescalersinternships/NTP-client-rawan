package NTPclient

import (
	"fmt"
	"net"
)

func Client() ([]byte, error) {
	conn, err := net.Dial("udp", "pool.ntp.org:123")
	if err != nil {
		return nil, fmt.Errorf("error dialing ntp server: %v", err)
	}

	req := make([]byte, 48) // min NTP packet size is 48 bytes
	// | LI (2 bits) | VN (3 bits) | Mode (3 bits) | --> needs to be set by client
	req[0] = 0b00011011 // LI = 00 (normal), VN = 3 , Mode = 3 (client)

	conn.Write(req)

	resp := make([]byte, 48)
	_, err = conn.Read(resp)

	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	return resp, nil
}
