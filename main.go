package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/jacobsa/go-serial/serial"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	portFlag    = kingpin.Flag("port", "port name (--port=COM3)").Required().String()
	baudRate    = kingpin.Flag("baud-rate", "baud rate (--baud-rate=9600)").Default("9600").Int()
	readTime    = kingpin.Flag("read-time", "read cycle time(ms)").Default("100").Int()
	interactive = kingpin.Flag("interactive", "interactive mode").Short('i').Bool()
)

func main() {
	kingpin.Parse()

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

	if *interactive {
		// interactive mode
		interactiveMode(port)
	} else {
		// one-shot mode
		text, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		_, err = port.Write([]byte(text))
		if err != nil {
			log.Fatalf("port.Write: %v", err)
		}
	}
}

func interactiveMode(port io.ReadWriteCloser) {

	fmt.Fprintln(os.Stdout, "This is Katuobushi interactive mode.")
	fmt.Fprintln(os.Stdout, "Please enter the sending texts...")

	go func() {
		t1 := time.NewTicker(time.Duration(*readTime) * time.Millisecond)
		defer t1.Stop()
		for {
			select {
			case <-t1.C:
				//Read
				buf := make([]byte, 128)
				n, err := port.Read(buf)
				if n != 0 {
					if err != nil {
						if err != io.EOF {
							fmt.Fprintln(os.Stdout, "Error reading from serial port: ", err)
						}
					} else {
						buf = buf[:n]
						//fmt.Println("n =", n)
						fmt.Fprintf(os.Stdout, "%s", buf)
					}
				}
			}
		}
	}()

	// Write
	go func() {
		f := bufio.NewScanner(os.Stdin)
		for f.Scan() {
			text := f.Text()
			text = text + "\n"
			_, err := port.Write([]byte(text))
			if err != nil {
				log.Fatalf("port.Write: %v", err)
			}
		}
	}()

	for {
		// loop
	}
}
