package main

import (
	"errors"
	"fmt"
	"net"
	"testing"
	"time"
)

/*type Conn interface {
	Read(b []byte) (n int, err error)

	Write(b []byte) (n int, err error)

	Close() error

	LocalAddr()

	RemoteAddr()

	SetDeadline(t time.Time) error

	SetReadDeadline(t time.Time) error

	SetWriteDeadline(t time.Time) error
}*/

type Conn struct {
	buf  []byte
	r, w int
}

/*type ConnBuffer struct {
	fd *conn
}*/

func (b *Conn) LocalAddr() net.Addr {
	//some code
	return nil
}
func (b *Conn) RemoteAddr() net.Addr {
	// some code
	return nil
}
func (b *Conn) SetDeadline(t time.Time) error {
	return nil
}
func (b *Conn) SetReadDeadline(t time.Time) error {
	return nil
}
func (b *Conn) SetWriteDeadline(t time.Time) error {
	return nil
}

func (b *Conn) Read(p []byte) (n int, err error) {
	if b.r == b.w {
		errReadEmpty := errors.New("errReadEmpty")
		return 0, errReadEmpty
	}
	n = copy(p, b.buf[b.r:b.w])
	b.r += n
	if b.r == b.w {
		b.r = 0
		b.w = 0
	}
	return n, nil
}

// check length

func (c *Conn) Close() error {

	return nil
}

func (b *Conn) Write(p []byte) (n int, err error) {
	// Slide existing data to beginning.
	if b.r > 0 && len(p) > len(b.buf)-b.w {
		copy(b.buf, b.buf[b.r:b.w])
		b.w -= b.r
		b.r = 0
	}

	// Write new data.
	n = copy(b.buf[b.w:], p)
	b.w += n
	if n < len(p) {
		err = errors.New("errWriteFull")
	} else {
		err = nil
	}
	return n, err
}

func TestServerRead(t *testing.T) {
	b := make([]byte, 4)
	b[0] = 0x00
	b[1] = 0x01
	b[2] = 10
	c := &Conn{buf: b, r: 0, w: 4}
	rCount, err1 := c.Read(b)
	fmt.Println("read count is : \n", rCount)
	fmt.Println(b)
	if err1 != nil && rCount != 4 {
		t.Fatal("failed to read ")
	}
}
func TestServerReadForReadingLengthZero(t *testing.T) {
	b := make([]byte, 4)
	b[0] = 0x00
	b[1] = 0x01
	b[2] = 10
	c := &Conn{buf: b, r: 0, w: 0}

	rCount, err1 := c.Read(b)
	fmt.Println("read count is : \n", rCount, err1)
	fmt.Println(b)
	if err1 != nil {
		fmt.Println("\nSuccess\n")
	} else {
		t.Fatal("failed to read zero Length ")
	}
}

func TestServerMsgCheck(t *testing.T) {

	msg := checkmessage("Hello")
	if msg != "Hi" {
		t.Fatal("failed to Interpret message ")
	}
}
func TestServerWrite(t *testing.T) {
	b := make([]byte, 4)
	b[0] = 0x00
	b[1] = 0x01
	b[2] = 10
	c := &Conn{buf: b, r: 0, w: 0}

	wCount, err := c.Write(b)
	fmt.Println("write count is : \n", wCount)
	fmt.Println(c.buf)
	if err != nil {
		t.Fatal("failed to write ")
	}

}
func TestServerWriteForWrittingLengthZero(t *testing.T) {
	b := make([]byte, 4)
	b[0] = 0x00
	b[1] = 0x01
	b[2] = 10
	c := &Conn{buf: b, r: 0, w: 4}

	wCount, err := c.Write(b)
	fmt.Println("write count is : \n", err, wCount)
	fmt.Println(c.buf)
	if wCount != 0 {
		t.Fatal("failed to write blank ")
	}

}
func TestServerGenerateMsg(t *testing.T) {

	b := make([]byte, 20)
	c := &Conn{buf: b, r: 0, w: 4}
	err := generateMsg(c, b)
	if err != nil {
		t.Fatal("failed to Handle Read Write Operation ")
	}

}
