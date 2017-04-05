// +build ignore

package main

import (
	"io"
	"net"
	"time"

	"github.com/erikdubbelboer/fasthttp"
)

type mockConn struct {
	data    []byte
	written bool
}

func (c *mockConn) Read(b []byte) (n int, err error) {
	if len(c.data) == 0 {
		return 0, io.EOF
	}
	n = copy(b, c.data)
	c.data = c.data[n:]
	return
}

func (c *mockConn) Write(b []byte) (n int, err error) {
	c.written = true
	return len(b), nil
}

func (c *mockConn) Close() error {
	return nil
}

func (c *mockConn) LocalAddr() net.Addr {
	return &net.TCPAddr{net.IP{127, 0, 0, 1}, 49706, ""}
}

func (c *mockConn) RemoteAddr() net.Addr {
	return &net.TCPAddr{net.IP{127, 0, 0, 1}, 49706, ""}
}

func (c *mockConn) SetDeadline(t time.Time) error {
	return nil
}

func (c *mockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *mockConn) SetWriteDeadline(t time.Time) error {
	return nil
}

func main() {
	conn := &mockConn{
		data: []byte("0 \nTransfer-Encoding:\n\n\xff"),
	}

	if err := fasthttp.ServeConn(conn, func(ctx *fasthttp.RequestCtx) {
		println("handler")
	}); err != nil {
		println(err.Error())
	}

	println("done")
}
