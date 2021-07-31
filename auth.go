package mp3lify

import "net/http"

func auth(req *http.Request) bool {
	token := req.Header.Get("X-Auth-Token")
	if token == authToken {
		return true
	}
	return false
}
