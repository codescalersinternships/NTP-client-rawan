# NTP Client 

This repository implements an NTP client that sends out timestamp request, accepts response, parses it.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)


## Installation

1. Clone the repository

   ```bash
   git clone https://github.com/codescalersinternships/NTP-client-rawan.git
   ```
## APIs

- `Client(server string) ([]byte, error)` : Interfaces with the server, responsible for dialing, sending the timestamp request and fetches the server response
- `ParseNTPResponse(resp []byte) (string, error)` : Responsible for parsing the server response into the required format

## Usage

```go
    resp, err := pkg.Client(server)
	if err != nil {
		panic(err)
	}

	time, err := pkg.ParseNTPResponse(resp)
	if err != nil {
		panic(err)
	}
```

### Flags

   ```bash
    go run cmd/main.go -server <CUSTOM_SERVER>
   ```

- **-server** is an optional flag that specifies a custom server, if not specified the client uses the default server  `pool.ntp.org:123`

