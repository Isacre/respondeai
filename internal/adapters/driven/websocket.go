package infra

import (
	"net/http"
	"v1/internal/core/ports"

	"github.com/gorilla/websocket"
)

type WebSocketAdapter struct {
	upgrader websocket.Upgrader
}

func NewWebSocketAdapter() *WebSocketAdapter {
	return &WebSocketAdapter{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

type WebSocketConn struct {
	*websocket.Conn
}

func (a *WebSocketAdapter) Upgrade(w http.ResponseWriter, r *http.Request) (ports.WebSocketConn, error) {
	conn, err := a.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return &WebSocketConn{Conn: conn}, nil
}

func (a *WebSocketAdapter) ReadMessage(conn ports.WebSocketConn) (int, []byte, error) {
	return conn.(*WebSocketConn).ReadMessage()
}

func (a *WebSocketAdapter) WriteMessage(conn ports.WebSocketConn, messageType int, data []byte) error {
	return conn.(*WebSocketConn).WriteMessage(messageType, data)
}
