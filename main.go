package main

import (
	_ "fmt"
	"github.com/jacobsa/go-serial/serial"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"log"
	"os"
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

	// fmt.Println("Wrote", n, "bytes.")
}
