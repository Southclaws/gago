package gago

import (
	"errors"
	"net/http"
)

const endpoint = "https://www.google-analytics.com/collect"

// Client represents an analytics client. All fields are mandatory.
type Client struct {
	ID                 string
	ClientIDContextKey string
	HitTypeContextKey  string
	Errors             func(error)
}

var client = http.Client{}

// Middleware for analytics.
func (c *Client) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid, ok := r.Context().Value(c.ClientIDContextKey).(string)
		if !ok {
			c.Errors(errors.New("could not extract client ID from request context"))
		}

		ht, ok := r.Context().Value(c.ClientIDContextKey).(string)
		if !ok {
			ht = "pageview"
		}

		client.PostForm(endpoint, map[string][]string{
			"v":   {"1"},
			"tid": {c.ID},
			"cid": {cid},
			"t":   {ht},
		})
		next.ServeHTTP(w, r)
	})
}
