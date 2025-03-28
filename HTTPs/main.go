package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
resp = struct {Status int, StatusCode string, Body io.ReaderCloser}

io.ReadCloser interface { Reader interface, Closer interface}
io.Reader interface { Read([]byte) (int, error) }
io.Closer interface { Close() (error) }
*/

// custom Writer
type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Println("Error:", err)
		return
	}
	/*
		Read() takes in bs []byte, injecting the raw data into bs []byte and return (int, error)
		This allows:
			1. The caller to have full control over       byte[], i.e byte[] will not be initialized inside Read().
			2. byte[] can be reused again by passing into Read() next time.
			3. Since byte[] is not initialized inside Read(), no heap allocation. Only once when byte[] was first initialized.
	*/
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// equivalent
	/*
		Copy expects (Writer interface, reader interface)
		resp.Body is a ReadCloser interface, however it also embeds Reader interface therefore it's also a Reader interface
		However Reader interface is not necessarily a ReadCloser interface

		"If interface A embeds interface B, then any type that implements A also implements B — but not the other way around."
	*/

	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

// Now logWriter implements Writer interface
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}
