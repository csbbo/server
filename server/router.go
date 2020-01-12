package server

import (
	"net/http"
	account "server/account/controller"
)

func init() {
	http.HandleFunc("/api/user", account.User)
}
