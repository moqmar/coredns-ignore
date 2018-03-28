package ignore

import (
	"context"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"

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
func (d Ignore) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	w.Hijack()
	w.Close()

	return plugin.NextOrFailure(d.Name(), nil, ctx, w, r)
}

// Name implements the Handler interface.
func (d Ignore) Name() string { return "ignore" }
