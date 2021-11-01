package main

import (
	"time"

	"github.com/gorilla/websocket"
)

//clientはチャットを行っている1人のユーザーを表します
type client struct {
	//socketはこのクライアントのためのWebSocketです
	socket *websocket.Conn
	//send はメッセージが送られるチャネルです
	send chan *message
	//roomはこのクライアントが参加しているチャットルームです。
	room *room
	// userDataはユーザーに関する情報を保持します
	userData map[string]interface{}
}

func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now().Format("2006/01/02 15:04:05")
			msg.Name = c.userData["name"].(string)
			if avatarURL, ok := c.userData["avatar_url"]; ok {
				msg.AvatarURL = avatarURL.(string)
			}
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
