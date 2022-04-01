package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/nexsabre/mikropoker-go/db"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// connection is an middleman between the websocket connection and the hub.
type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}

// readPump pumps messages from the websocket connection to the hub.
func (s subscription) readPump() {
	c := s.conn
	defer func() {
		h.unregister <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}

		m := handleCmds(msg, s.room)
		h.broadcast <- m
	}
}

type WSAction struct {
	Action  string `json:"action"`
	Payload string `json:"payload"`
}

func handleActions(action *WSAction, sessionID int) {
	payload := strings.Split(action.Payload, ",")
	// TODO NexSabre clean this in UI
	switch action.Action {
	case "salle":
		// payload:"NexSabre,32" (nickname, salle)
		salle, _ := strconv.ParseFloat(payload[2], 32)
		db.UserPointsForUser(db.GetDB(), sessionID, payload[1], float32(salle))
	case "reveal":
		// payload:"false" (reveal_or_hide)
		reveal, _ := strconv.ParseBool(payload[1])
		db.RevealSession(db.GetDB(), sessionID, reveal)
	case "restart":
		// paylod:"" ()
		db.RestartSession(db.GetDB(), sessionID)
	}
}

func handleCmds(msg []byte, sessionID string) message {
	if len(msg) > 0 {
		wsAction := &WSAction{}
		err := json.Unmarshal(msg, wsAction)
		if err != nil {
			fmt.Printf("Canno handle action '%+v'", msg)
		} else {
			intSessionID, _ := strconv.Atoi(sessionID)
			handleActions(wsAction, intSessionID)
		}
	}

	_sessionID, _ := strconv.Atoi(sessionID)
	session := db.GetSession(db.GetDB(), _sessionID)
	msgFromDB, err := json.Marshal(session)
	if err != nil {
		fmt.Printf("Cannot handle command %+v", err)
	}

	return message{msgFromDB, sessionID}
}

// write writes a message with the given message type and payload.
func (c *connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket connection.
func (s *subscription) writePump() {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func serveWs(w http.ResponseWriter, r *http.Request, roomId string) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws}
	s := subscription{c, roomId}
	h.register <- s
	go s.writePump()
	go s.readPump()
}
