package handler

import (
	"net/http"
	"strings"

	"gitlab.com/erloom.id/libraries/go/backend-skeleton/websocket"
)

func (handler *Handler) ServeWS(rw http.ResponseWriter, r *http.Request) {
	secProtocols := strings.Split(r.Header.Get("Sec-Websocket-Protocol"), ", ")
	if len(secProtocols) < 2 {
		return
	}
	tokenString := secProtocols[1]

	token := handler.BackendSkeleton.JWT.ValidateToken(tokenString)
	if token == nil {
		return
	}
	if !token.Valid {
		return
	}

	websocket.ServeWs(handler.BackendSkeleton.WebsocketHub, rw, r)

}
