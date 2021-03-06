package lmgtfy

import (
	"fmt"
	"net/url"

	"github.com/zquestz/s/providers"
)

func init() {
	providers.AddProvider("lmgtfy", &Provider{})
}

// Provider merely implements the Provider interface.
type Provider struct{}

// BuildURI generates a search URL for LMGTFY.
func (p *Provider) BuildURI(q string) string {
	return fmt.Sprintf("http://lmgtfy.com/?q=%s", url.QueryEscape(q))
}
