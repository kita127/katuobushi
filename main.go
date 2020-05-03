package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/jacobsa/go-serial/serial"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	portFlag = kingpin.Flag("port", "port name (--port=COM3)").Required().String()
	baudRate = kingpin.Flag("baud-rate", "baud rate (--baud-rate=9600)").Default("9600").Int()
	readTime = kingpin.Flag("read-time", "read cycle time(ms)").Default("100").Int()
)

func main() {
	kingpin.Parse()

	// 全部読み
	//text, err := ioutil.ReadAll(os.Stdin)
	//if err != nil {
	//	log.Fatal(err)
	//}

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

	go func() {
		t1 := time.NewTicker(time.Duration(*readTime) * time.Millisecond)
		defer t1.Stop()
		for {
			select {
			case <-t1.C:
				//Read
				buf := make([]byte, 32)
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
			_, err = port.Write([]byte(text))
			if err != nil {
				log.Fatalf("port.Write: %v", err)
			}
		}
	}()

	for {
		// loop
	}

	// fmt.Println("Wrote", n, "bytes.")
}
