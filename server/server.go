// 一个简单的telnet server

package server

import (
	"bufio"
	"net"
	tk "rank/toolkit"
	"strconv"
	"strings"
)

type Command struct {
	Cmd     string
	Args    []string
	Conn    net.Conn
	Result  string
	Session *Session
}

type Handler interface {
	Handle(*Command)
}

type ServerInfo struct {
	protocol string
	port     int
	running  bool
	events   chan Command
	Router   map[string]Handler
}

type Session struct {
	Id      int64
	Conn    net.Conn
	Quit    bool
	Alive   bool
	IsAdmin bool
}

func handleConnection(conn net.Conn, ch chan Command) {
	// 处理连接的逻辑代码
	tk.L("Accepted once ")
	defer conn.Close()
	session := Session{
		Id:      tk.Alloc(),
		Conn:    conn,
		Quit:    false,
		Alive:   true,
		IsAdmin: false,
	}
	var onAcceptCmd Command = Command{
		Cmd:     "_init",
		Conn:    conn,
		Session: &session,
	}
	ch <- onAcceptCmd
	reader := bufio.NewReader(conn)
	var p []byte = make([]byte, 32)
	reader.Read(p)
	for session.Alive {
		input, err := reader.ReadString('\n')
		if err != nil {
			tk.L("Error reading from client:", err)
			return
		}
		input = strings.Trim(input, "\r\n")
		tk.L("Received: " + input)
		// if input == "exit" {
		// 	tk.L("client leave")
		// 	return
		// }
		strs := strings.Split(input+",", ",")
		cmd := Command{
			Cmd:     strs[0],
			Args:    strs[1:],
			Conn:    conn,
			Result:  "",
			Session: &session,
		}
		ch <- cmd
	}
}

func InitServer(protocol string, port int) *ServerInfo {
	var server *ServerInfo = new(ServerInfo)
	server.port = port
	server.protocol = protocol
	server.events = CreateChannel()
	server.Router = make(map[string]Handler, 128)

	return server
}

func (server *ServerInfo) RegHandler(cmd string, h Handler) {
	server.Router[cmd] = h
}

func (server *ServerInfo) LoopForever() {
	server.running = true
	listener, err := net.Listen(server.protocol, ":"+strconv.Itoa(server.port))
	defer listener.Close()
	go LoopEvent(server.events, server.Router)
	if err != nil {
		tk.F(err)
	}
	for server.running {
		conn, err := listener.Accept()
		if err != nil {
			tk.F(err)
		}
		go handleConnection(conn, server.events)
	}
}
