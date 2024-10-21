package imvu

import "encoding/json"

type loginRequest struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	CookieAcceptance bool   `json:"gdpr_cookie_acceptance"`
}

func (ic *IMVUClient) Login(username, password string) error {
	req := loginRequest{
		Username:         username,
		Password:         password,
		CookieAcceptance: false,
	}

	json.Marshal(req)

	return nil
}
