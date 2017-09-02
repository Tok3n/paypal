package paypal
import (
	"google.golang.org/appengine/urlfetch"
	"golang.org/x/net/context"
)

var ctx context.Context
var isGAE = false
// NewClient returns a new Client struct
func NewClientForGAE(clientID, secret, APIBase string, c context.Context) *Client {
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