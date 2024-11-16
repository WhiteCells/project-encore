package queryparameters

import "context"

type ListBlogPost struct {
	PageLimit int `query:"limit"`
	Author    string
}

type ListParams struct {
	Limit  uint16
	Offset uint16
}

type ListResponse struct {
	Post []*ListBlogPost
}

//encore:api public method=GET path=/blog
func ListBlogPosts(ctx context.Context, opts *ListParams) (*ListResponse, error) {
	return &ListResponse{Post: []*ListBlogPost{
		{PageLimit: 1, Author: "cells"},
		{2, "visitor"},
	}}, nil
}
