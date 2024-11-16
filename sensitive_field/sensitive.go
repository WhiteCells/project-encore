package sensitivefield

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Updates struct {
	Author      string    `json:"author"`
	PublishTime time.Time `json:"publishTime"`
}

type BatchUpdateParams struct {
	Requester     string    `header:"X-Requester"`
	RequestTime   time.Time `header:"X-Requester-Time"`
	CurrentAuthor string    `query:"author"`
	Updates       *Updates  `json:"updates"`
	SecretKey     string    `encore:"sensitive"`
}

type BathcUpdateResponse struct {
	ServeBy    string      `header:"X-Served-By"`
	UpdatedIDs []uuid.UUID `json:"update_ids"`
}

//encore:api public method=POST path=/section/:sectionID/posts
func BatchUpdate(ctx context.Context, sectionID string, params *BatchUpdateParams) (*BathcUpdateResponse, error) {
	// if params.SecretKey == "" {
	// 	return nil, encore.NewError("missing secret_key", "")
	// }
	return &BathcUpdateResponse{
		ServeBy: "cells",
		UpdatedIDs: []uuid.UUID{
			uuid.New(),
			uuid.New(),
		},
	}, nil
}
