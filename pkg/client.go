package NTPclient

import (
	"fmt"
	"net"
	"time"
)

// NTP epoch starts at 1900, Unix at 1970 --> 70 years difference
const seventyYears = 2208988800

func ParseNTPResponse(resp []byte) (string, error) {
	//the timestamp starts at byte 40 of the received packet, concatenate using big-endian

	secsSince1900 := uint32(resp[40])<<24 | uint32(resp[41])<<16 | uint32(resp[42])<<8 | uint32(resp[43])
	frac := uint32(resp[44])<<24 | uint32(resp[45])<<16 | uint32(resp[46])<<8 | uint32(resp[47])

	secsUnix := secsSince1900 - seventyYears // convert to Unix epoch

	nano := (frac * 1e9) / (2 ^ 32) // convert frac part to nanoseconds

	ntpTime := time.Unix(int64(secsUnix), int64(nano))

	return ntpTime.String(), nil

}

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

	if len(resp) < 48 {
		return nil, fmt.Errorf("response too short")
	}

	return resp, nil
}
