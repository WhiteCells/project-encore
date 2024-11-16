package ping

import "context"

type PingParams struct {
	Name string
}

type PingResponse struct {
	Message string
}

// encore:api public path=/ping
func Ping(ctx context.Context, param *PingParams) (*PingResponse, error) {
	msg := "Ping " + param.Name
	return &PingResponse{Message: msg}, nil
}
