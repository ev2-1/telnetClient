package telnetClient

import (
	"github.com/reiver/go-telnet"
)

// The Controller struct
type Controller struct {
	Conn           *telnet.Conn
	ResponseStream chan string
}

// Create a new controller that automatically revices and dumps content to the ResponseStream channel in Controller
// selftype is the first thing send my client to identify
func NewReciveController(dest, selftype string) (*Controller, error) {
	conn, err := telnet.DialTo(dest)
	if err != nil {
		return nil, err
	}

	c := &Controller{
		Conn:           conn,
		ResponseStream: make(chan string),
	}

	c.Write(selftype)

	// read:
	c.ReadToChannel('\r')

	return c, nil
}

// creates a new controller
func NewController(dest string) (*Controller, error) {
	conn, err := telnet.DialTo(dest)
	if err != nil {
		return nil, err
	}

	return &Controller{
		Conn:           conn,
		ResponseStream: make(chan string),
	}, nil
}

// reads until b is encounterd
func (c *Controller) ReadUntil(b byte) ([]byte, error) {
	var bytes []byte
	notFound := true

	for notFound {
		nbyte := make([]byte, 1)
		_, err := c.Conn.Read(nbyte)
		if err != nil {
			return nil, err
		}

		if nbyte[0] == b {
			notFound = false
		} else {
			bytes = append(bytes, nbyte...)
		}
	}

	return bytes, nil
}

// writes string
func (c *Controller) Write(cmd string) error {
	_, err := c.Conn.Write([]byte(cmd + "\n"))
	return err
}

// writes cmd and returns when reciving response
func (c *Controller) Exec(cmd string) ([]string, error) {
	err := c.Write(cmd)
	if err != nil {
		return make([]string, 0), err
	}

	bytes, err := c.ReadUntil('\n')
	if err != nil {
		return make([]string, 0), err
	}

	return ParseResponse(string(bytes))
}

// reads until b, then sends response to ResponseStream channel
func (c *Controller) ReadToChannel(b byte) {
	go func() {
		for {
			res, err := c.ReadUntil(b)
			if err != nil {
				continue
			} else {
				c.ResponseStream <- string(res)
			}
		}
	}()
}

// closes the connection
func (c *Controller) Close() error {
	return c.Conn.Close()
}
