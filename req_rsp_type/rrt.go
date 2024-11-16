package rrt

import "context"

type Request1 struct {
	Header string `header:"X-Header"`
	Query  string `query:"query"`
	Nested struct {
		Header2 string `header:"X-Header2"`
		Query   string `query:"query2"`
	} `json:"nested"`
}

type Response1 struct {
	Header string `header:"X-Header"`
	Query  string `query:"query"`
	Nested struct {
		Header2 string `header:"X-Header2"`
		Query   string `query:"query2"`
	} `json:"nested"`
}

type Request struct {
	XCode string `header:"X-Header"`
}

type Response struct {
	Status string `header:"X-Header"`
}

type LoginResponse struct {
	SessionOID string `header:"Set-Cookie"`
}

//encore:api public method=POST path=/login
func Login(ctx context.Context) (*LoginResponse, error) {
	return &LoginResponse{SessionOID: "session=a=1;b=2"}, nil
}
