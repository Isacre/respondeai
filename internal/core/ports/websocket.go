package ports

import "net/http"

type WebSocketConnection interface {
    Upgrade(w http.ResponseWriter, r *http.Request) (WebSocketConn, error)
    ReadMessage(conn WebSocketConn) (messageType int, p []byte, err error)
    WriteMessage(conn WebSocketConn, messageType int, data []byte) error
}

type WebSocketConn interface {
    Close() error
}