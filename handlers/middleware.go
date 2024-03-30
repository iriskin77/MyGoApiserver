package handlers

import (
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get(authorizationHeader)
	if header == "" {

		w.Write([]byte("401"))

	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Authorization" {
		w.Write([]byte("401"))

	}

	if len(headerParts[1]) == 0 {
		w.Write([]byte("401"))

	}

	userId, err := h.services.Users.ParseToken(headerParts[1])
	if err != nil {
		w.Write([]byte("401"))

	}

	bs := []byte(strconv.Itoa(userId))
	w.Write(bs)

}
