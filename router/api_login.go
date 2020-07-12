package router

type loginData struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

type loginPayload struct {
	payloadBase

	LoginInfo loginData `json:"data"`
}

type LoginResp struct {
	csrfParams
	Error string `json:"errorCategory"`
}

func (r *Client) Login(username, password string) LoginResp {
	payload := loginPayload{
		LoginInfo: loginData{
			UserName: "admin",
			Password: HashPassword(username, password, r.session.csrfParams),
		},
	}
	payload.CSRF = r.session.csrfParams

	resp := LoginResp{}
	r.post("/api/system/user_login", payload, &resp)
	r.session.csrfParams = resp.csrfParams

	return resp
}
