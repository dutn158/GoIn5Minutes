/**
This program reads data from the connection and writes it to the standard output until and end-of-file condition
or an error occurs. The mustCopy function is a utility used in several examples in this section.
 */
package main

import (
	"net"
	"log"
	"io"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
