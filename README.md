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

## Usage

### Flags

   ```bash
    go run cmd/main.go -server <CUSTOM_SERVER>
   ```

- **-server** is an optional flag that specifies a custom server, if not specified the client uses the default server  `pool.ntp.org:123`

