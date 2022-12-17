package chat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"plugno-api/models"
	"plugno-api/structs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type ChatHandler struct {
	messageModel models.MessageModel
}

func NewChatHandler(s *structs.Server) *ChatHandler {
	return &ChatHandler{
		messageModel: s.MessageModel,
	}
}

type RawMessage struct {
	Message string `json:"message"`
	UserID  int    `json:"userId"`
}

type Client struct {
	chat *Chat

	conn *websocket.Conn

	send chan []byte

	messageModel models.MessageModel
}

const (
	writeWait = 10 * time.Second

	pongWait = 60 * time.Second

	pingPeriod = (pongWait * 9) / 10

	// idk
	maxMessageSize = 512
)

var (
	newline = []byte("\n")
	space   = []byte(" ")
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (c *Client) readPump() {
	defer func() {
		c.chat.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(appData string) error {
		c.conn.SetReadDeadline((time.Now().Add(pongWait)))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error: %v", err)
			}
			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		// decode bytes
		var rawMessage RawMessage
		decoder := bytes.NewReader(message)
		d := json.NewDecoder(decoder)

		err = d.Decode(&rawMessage)

		if err != nil {
			fmt.Println(err.Error())
		}

		insertId, err := c.messageModel.Create(rawMessage.Message, rawMessage.UserID)
		if err != nil {
			fmt.Println("Failed to create message:", err.Error())
			return
		}

		newMessage, err := c.messageModel.FindOne(insertId)
		if err != nil {
			fmt.Println("Failed to find message:", err.Error())
			return
		}

		body, err := json.Marshal(newMessage)
		if err != nil {
			fmt.Println("Failed to send new essage:", err.Error())
			return
		}

		c.chat.broadcast <- body
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (ch *ChatHandler) ServeWs(chat *Chat, c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		chat:         chat,
		conn:         conn,
		send:         make(chan []byte, 256),
		messageModel: ch.messageModel,
	}

	client.chat.register <- client

	go client.writePump()
	go client.readPump()
}
