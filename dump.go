package dump

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	corelog "github.com/coredns/coredns/plugin/log"
	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/coredns/coredns/plugin/pkg/replacer"

	"github.com/mholt/caddy"
	"github.com/miekg/dns"
)

// Ignore implement the plugin interface.
type Ignore struct {
	Next plugin.Handler
}

func init() {
	caddy.RegisterPlugin("ignore", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	for c.Next() {
		if c.NextArg() {
			return plugin.Error("ignore", c.ArgErr())
		}
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Ignore{Next: next}
	})

	return nil
}

// ServeDNS implements the plugin.Handler interface.
func (d Dump) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	w.

	return plugin.NextOrFailure(d.Name(), d.Next, ctx, w, r)
}

// Name implements the Handler interface.
func (d Dump) Name() string { return "dump" }
