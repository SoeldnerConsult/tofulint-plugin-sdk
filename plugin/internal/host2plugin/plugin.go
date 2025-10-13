package host2plugin

import (
	"context"

	"github.com/arsiba/tofulint-plugin-sdk/plugin/internal/proto"
	"github.com/arsiba/tofulint-plugin-sdk/tflint"
	plugin "github.com/hashicorp/go-plugin"
	"github.com/hashicorp/go-version"
	"google.golang.org/grpc"
)

// SDKVersion is the SDK version.
const SDKVersion = "0.23.0"

// minTofuLintVersionConstraint presents the minimum version of TofuLint that this SDK supports.
var minTofuLintVersionConstraint = version.MustConstraints(version.NewConstraint(">= 0.0.1"))

// handShakeConfig is used for UX. ProcotolVersion will be updated by incompatible changes.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  11,
	MagicCookieKey:   "TFLINT_RULESET_PLUGIN",
	MagicCookieValue: "5adSn1bX8nrDfgBqiAqqEkC6OE1h3iD8SqbMc5UUONx8x3xCF0KlPDsBRNDjoYDP",
}

// RuleSetPlugin is a wrapper to satisfy the interface of go-plugin.
type RuleSetPlugin struct {
	plugin.NetRPCUnsupportedPlugin

	impl tflint.RuleSet
}

var _ plugin.GRPCPlugin = &RuleSetPlugin{}

// GRPCServer returns an gRPC server acting as a plugin.
func (p *RuleSetPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterRuleSetServer(s, &GRPCServer{
		impl:   p.impl,
		broker: broker,
	})
	return nil
}

// GRPCClient returns an RPC client for the host.
func (*RuleSetPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		client: proto.NewRuleSetClient(c),
		broker: broker,
	}, nil
}
