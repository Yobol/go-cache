package tcp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"

	cache "github.com/yobol/go-cache/core"
)

const (
	OpTypeSet = 'S'
	OpTypeGet = 'G'
	OpTypeDel = 'D'

	SP = ' '
)

func New(c cache.Cache) *Server {
	return &Server{c}
}

type Server struct {
	cache.Cache
}

func (s *Server) Listen() {
	l, err := net.Listen("tcp", ":10616")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go s.process(conn)
	}
}

func (s *Server) process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		log.Println("start to handle the connection from", conn.RemoteAddr().String())
		op, err := reader.ReadByte()
		if err != nil {
			if err != io.EOF {
				log.Println("close connection, because", err)
				return
			}
		}
		switch op {
		case OpTypeSet:
			err = s.set(conn, reader)
		case OpTypeGet:
			err = s.get(conn, reader)
		case OpTypeDel:
			err = s.del(conn, reader)
		default:
			log.Println("close connection, because of unknown operator type", op)
		}
		if err != nil {
			log.Println("close connection, because", err)
		}
	}
}

func (s *Server) set(conn net.Conn, reader *bufio.Reader) error {
	k, v, err := s.readKeyAndValue(reader)
	if err != nil {
		return err
	}
	return s.sendResponse(nil, s.Set(k, v), conn)
}

func (s *Server) get(conn net.Conn, reader *bufio.Reader) error {
	k, err := s.readKey(reader)
	if err != nil {
		return err
	}
	v, err := s.Get(k)
	return s.sendResponse(v, err, conn)
}

func (s *Server) del(conn net.Conn, reader *bufio.Reader) error {
	k, err := s.readKey(reader)
	if err != nil {
		return err
	}
	return s.sendResponse(nil, s.Del(k), conn)
}

func (s *Server) readKeyAndValue(reader *bufio.Reader) (string, []byte, error) {
	klen, err := s.readLen(reader)
	if err != nil {
		return "", nil, err
	}
	vlen, err := s.readLen(reader)
	if err != nil {
		return "", nil, err
	}

	k := make([]byte, klen)
	_, err = io.ReadFull(reader, k)
	if err != nil {
		return "", nil, err
	}
	v := make([]byte, vlen)
	_, err = io.ReadFull(reader, v)
	if err != nil {
		return "", nil, err
	}
	return string(k), v, err
}

func (s *Server) readKey(reader *bufio.Reader) (string, error) {
	klen, err := s.readLen(reader)
	if err != nil {
		return "", err
	}
	k := make([]byte, klen)
	_, err = io.ReadFull(reader, k)
	if err != nil {
		return "", err
	}
	return string(k), nil
}

func (s *Server) readLen(reader *bufio.Reader) (int, error) {
	tmp, err := reader.ReadString(SP)
	if err != nil {
		return 0, err
	}
	l, err := strconv.Atoi(strings.TrimSpace(tmp))
	if err != nil {
		return 0, err
	}
	return l, nil
}

func (s *Server) sendResponse(value []byte, e error, conn net.Conn) error {
	if e != nil {
		errMsg := e.Error()
		msg := fmt.Sprintf("-%d %s", len(errMsg), errMsg)
		_, err := conn.Write([]byte(msg))
		return err
	}
	msg := fmt.Sprintf("%d ", len(value))
	_, err := conn.Write(append([]byte(msg), value...))
	return err
}
