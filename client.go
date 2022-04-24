package telnetClient

import (
	"github.com/reiver/go-telnet"
)

type Controller struct {
	Conn *telnet.Conn
}

func NewController(dest string) (*Controller, error) {
	conn, err := telnet.DialTo(dest)
	if err != nil {
		return nil, err
	}

	return &Controller{Conn: conn}, nil
}

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

func (c *Controller) Write(cmd string) error {
	_, err := c.Conn.Write([]byte(cmd + "\n"))
	return err
}

func (c *Controller) Exec(cmd string) (string, error) {
	err := c.Write(cmd)
	if err != nil {
		return "", err
	}

	bytes, err := c.ReadUntil('\n')
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (c *Controller) Close() error {
	return c.Conn.Close()
}
