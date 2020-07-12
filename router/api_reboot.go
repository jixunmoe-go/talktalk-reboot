package router

type rebootPayload struct {
	payloadBase
}

type RebootResp struct {
	csrfParams
	ErrorCode int `json:"errcode"`
}

func (r *Client) Reboot() RebootResp {
	payload := rebootPayload{}
	payload.CSRF = r.session.csrfParams

	resp := RebootResp{}
	r.post("/api/service/reboot.cgi", payload, &resp)
	r.session.csrfParams = resp.csrfParams

	return resp
}
