package websocketsrv

import (
    "v1/internal/core/ports"
    "net/http"
)

type WebSocketService struct {
    adapter ports.WebSocketConnection
}

func NewWebSocketService(adapter ports.WebSocketConnection) *WebSocketService {
    return &WebSocketService{adapter: adapter}
}

func (websocket *WebSocketService) HandleWebSocket(writer http.ResponseWriter, reader *http.Request) {
    conn, err := websocket.adapter.Upgrade(writer, reader)
    if err != nil {
        http.Error(writer, "Failed to upgrade to websocket", http.StatusInternalServerError)
        return
    }
    defer conn.Close()

    for {
        messageType, message, err := websocket.adapter.ReadMessage(conn)
        if err != nil {
            return
        }
        if err = websocket.adapter.WriteMessage(conn, messageType, message); err != nil {
            return
        }
    }
}
