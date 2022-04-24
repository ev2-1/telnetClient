package telnetClient

import (
	"github.com/reiver/go-telnet"
)

type controller struct {
	conn *telnet.Conn
}

func NewController(dest string) (*controller, error) {
	conn, err := telnet.DialTo(dest)
	if err != nil {
		return nil, err
	}

	return &controller{conn: conn}, nil
}

func (c *controller) ReadUntil(b byte) ([]byte, error) {
	var bytes []byte
	notFound := true

	for notFound {
		nbyte := make([]byte, 1)
		_, err := c.conn.Read(nbyte)
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

func (c *controller) Write(cmd string) error {
	_, err := c.conn.Write([]byte(cmd + "\n"))
	return err
}

func (c *controller) Exec(cmd string) (string, error) {
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

func (c *controller) Close() error {
	return c.conn.Close()
}
