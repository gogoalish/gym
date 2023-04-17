package server

import (
	"context"
	"net/http"
)

const ctxKeyUser ctxKey = "user"

type ctxKey string

func (h *Handler) contextSetUser(r *http.Request, user *User) *http.Request {
	ctx := context.WithValue(r.Context(), ctxKeyUser, user)
	return r.WithContext(ctx)
}

func (h *Handler) contextGetUser(r *http.Request) *User {
	user, ok := r.Context().Value(ctxKeyUser).(*User)
	if !ok {
		return nil
	}

	return user
}
