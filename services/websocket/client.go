package websocket

import (
	"daimon/pkg/log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Id      string
	Group   string
	Socket  *websocket.Conn
	Message chan []byte
}

// 从webscoket 读取数据
func (c *Client) Read() {
	defer func() {
		wsManager.unRegisterChan <- c
		log.L.Infof("client [%s] disconnected", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.L.Infof("client [%s] disconnet err: %s", c.Id, err)
		}
	}()

	for {
		messageType, message, err := c.Socket.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			break
		}
		log.L.Infof("client [%s] receive message: %s", c.Id, string(message))
		c.Message <- message
	}

}

// ws 写入数据
func (c *Client) Write() {
	defer func() {
		log.L.Infof("client [%s] disconnect", c.Id)
		if err := c.Socket.Close(); err != nil {
			log.L.Warnf("client [%s] disconnect err: ", c.Id, err)
		}
	}()

	for {
		select {
		case message, ok := <-c.Message:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.L.Infof("client [%s] write message: %s", c.Id, string(message))
			err := c.Socket.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				log.L.Warnf("client [%s] write message error: %s", c.Id, err)
			}
		}
	}

}
