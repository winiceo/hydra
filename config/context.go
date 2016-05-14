package config

import (
	"github.com/ory-am/fosite/hash"
	"github.com/ory-am/hydra/firewall"
	"github.com/ory-am/ladon"
	"github.com/ory-am/fosite/handler/core"
	"github.com/ory-am/hydra/pkg"
)

type Context struct {
	Connection     interface{}

	Hasher         hash.Hasher
	Warden         firewall.Firewall
	LadonManager   ladon.Manager
	FositeStrategy core.CoreStrategy
	FositeStore    pkg.FositeStorer
}