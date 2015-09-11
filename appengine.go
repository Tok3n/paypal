package paypal
import (
"appengine/urlfetch"
"appengine"
)

var ctx appengine.Context
var isGAE = false
// NewClient returns a new Client struct
func NewClientForGAE(clientID, secret, APIBase string, c appengine.Context) *Client {
	ctx = c
	isGAE = true
	return &Client{
		urlfetch.Client(ctx),
		clientID,
		secret,
		APIBase,
		nil,
	}
}