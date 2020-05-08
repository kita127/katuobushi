# Katuobushi
Katsuobushi is a USB serial communication device that runs on your OS terminal.

## Environment
* Windows 10

Other OS may work, but I have not tried it.

## Description
Katsuobushi is a CLI tool that allows USB serial communication on your OS terminal.

You configure port and baud rate when it executes.
The port is configured with `--port`.  You can check the target port with `MODE` command if you are using windows OS.
Baud rate is configured with `--baud-rate`(default 9600).

### One-shot mode
It sends the target some texts received by standard-in.

    $echo hello | katuobushi --port COM3
    or
    $cat input.txt | katuobushi --port COM3

### Interactive mode
`--interactive`(short-flag `-i`) option allows that the tool communicates target in interactive mode.

#### Send
It sends a line as ASCII which are entered by standard-in.
Each text is split by new line.

#### Receive
It outputs received data as ASCII into standard-out.
Receivable data size is 128 bytes.
Received data are read by polling.
That polling timing can be configured by `--read-time`(default 100ms).

## Installation

binalies<br>
https://github.com/kita127/katuobushi/releases

If you use Go,

    $go get github.com/kita127/katuobushi

## Usage

    usage: katuobushi.exe --port=PORT [<flags>]

    Flags:
          --help            Show context-sensitive help (also try --help-long and
                            --help-man).
          --port=PORT       port name (--port=COM3)
          --baud-rate=9600  baud rate (--baud-rate=9600)
          --read-time=100   read cycle time(ms)
      -i, --interactive     interactive mode


## License
This software is released under the MIT License, see LICENSE.
