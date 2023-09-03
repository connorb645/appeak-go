package bootstrap

import (
	"github.com/connorb645/appeak-go/store"
	"github.com/connorb645/appeak-go/zendesk"
)

func NewDocumentStore() store.HelpCenterProvider {
	return zendesk.NewZendesk(
		"blacksapps.zendesk.com",
		"connor@blacksapps.co.uk",
		"dnZIeFIDAhYXtFmQhTpKrm09zhLRlPQld4THYXM1",
	)
}
