package main

import (
	"fmt"
	_ "fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/jacobsa/go-serial/serial"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	portFlag = kingpin.Flag("port", "port name (--port=COM3)").Required().String()
	baudRate = kingpin.Flag("baud-rate", "baud rate (--baud-rate=9600)").Default("9600").Int()
)

func main() {
	kingpin.Parse()

	text, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	// Set up options.
	options := serial.OpenOptions{
		PortName:        *portFlag,
		BaudRate:        uint(*baudRate),
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	defer port.Close()

	_, err = port.Write([]byte(text))
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	go func() {
		t1 := time.NewTicker(100 * time.Millisecond)
		for {
			select {
			case <-t1.C:
				//Read
				buf := make([]byte, 32)
				n, err := port.Read(buf)
				if n != 0 {
					if err != nil {
						if err != io.EOF {
							fmt.Println("Error reading from serial port: ", err)
						}
					} else {
						buf = buf[:n]
						fmt.Println("n =", n)
						fmt.Printf("Rx: %s\n", buf)
					}
				}
			}
		}
	}()

	for {
	}

	// fmt.Println("Wrote", n, "bytes.")
}
