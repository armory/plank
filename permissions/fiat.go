package permissions

import (
	"github.com/armory/plank/client"
)

const (
	fiatWritePerm = "WRITE"
	fiatBaseURL   = "http://fiat"
)

// TODO: package variable for this is just asking for trouble. There will most certainly be side-effects.
var defaultFiatClient, _ = client.New(client.BaseURL(fiatBaseURL))
