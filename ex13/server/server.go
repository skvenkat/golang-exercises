package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type Server struct {
	listener         net.Listener
	quit             chan struct{}
	exited           chan struct{}
	db               memoryDB
	connections      map[int]net.Conn
	connCloseTimeout time.Duration
}

func NewServer() *Server {
	l, err := net.Listen("tcp", ":9040")
	if err != nil {
		log.Fatal("failed to create listener", err.Error())
	}

	srv := &Server{
		listener:         l,
		quit:             make(chan struct{}),
		exited:           make(chan struct{}),
		db:               newDB(),
		connections:      map[int]net.Conn{},
		connCloseTimeout: 10 * time.Second,
	}

	go srv.serve()
	return srv
}

func (srv *Server) serve() {
	var id int
	fmt.Println("listening for clients")
	for {
		select {
		case <-srv.quit:
			fmt.Println("shutting down the server")
			err := srv.listener.Close()
			if err != nil {
				fmt.Println("could not close listener", err.Error())
			}
			if len(srv.connections) > 0 {
				srv.warnConnections(srv.connCloseTimeout)
				<-time.After(srv.connCloseTimeout)
				srv.closeConnections()
			}
			close(srv.exited)
			return
		default:
			tcpListener := srv.listener.(*net.TCPListener)
			err := tcpListener.SetDeadline(time.Now().Add(2 * time.Second))
			if err != nil {
				fmt.Println("failed to set listener's deadline", err.Error())
			}

			conn, err := tcpListener.Accept()
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				continue
			}

			if err != nil {
				fmt.Println("failed to accept connection : ", err.Error())
			}

			write(conn, "Welcome to MemoryDB Server")
			//handlers.Add(1)
			srv.connections[id] = conn
			go func(connID int) {
				fmt.Printf("client with id %d joined!\n", connID)
				srv.handleConn(conn)
				//handlers.Done()
				delete(srv.connections, connID)
				fmt.Printf("client with id %d left!\n", connID)
			}(id)
			id++
		}
	}
}

func write(c net.Conn, s string) {
	_, err := fmt.Fprintf(c, "%s\n-> ", s)
	if err != nil {
		log.Fatal(err)
	}
}

func (srv *Server) handleConn(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		l := strings.ToLower(strings.TrimSpace(scanner.Text()))
		values := strings.Split(l, " ")

		switch {
		case len(values) == 3 && values[0] == "set":
			srv.db.set(values[1], values[2])
			write(conn, "WRITE_OK")
		case len(values) == 2 && values[0] == "get":
			val, found := srv.db.get(values[1])
			if !found {
				write(conn, fmt.Sprintf("key %s not found", values[1]))
			} else {
				write(conn, val)
			}
		case len(values) == 2 && values[0] == "delete":
			srv.db.delete(values[1])
			write(conn, "DELETE_OK")
		case len(values) == 1 && values[0] == "exit":
			if err := conn.Close(); err != nil {
				fmt.Println("could not close connection", err.Error())
			}
		default:
			write(conn, fmt.Sprintf("UNKNOWN: %s", l))
		}
	}
}

func (srv *Server) warnConnections(time.Duration) {
	for _, conn := range srv.connections {
		write(conn, fmt.Sprintf("host wants to shut down the server in %s seconds", srv.connCloseTimeout.String()))
	}
}

func (srv *Server) closeConnections() {
	fmt.Println("closing all connections")
	for id, conn := range srv.connections {
		err := conn.Close()
		if err != nil {
			fmt.Println("could not close the connection with id %d\n", id, err.Error())
		}
	}
}

func (srv *Server) Stop() {
	fmt.Println("stopping the database server")
	close(srv.quit)
	<-srv.exited
	fmt.Println("saving in-memory records to file")
	srv.db.save()
	fmt.Println("database server successfully stopped")
}
