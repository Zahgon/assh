package commands

import (
	"context"
	"io"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"moul.io/assh/v2/pkg/config"
)

type contextKey string

type gatewayErrorMsg struct {
	gateway string
	err     zap.Field
}

var syncContextKey contextKey = "sync"

var proxyCommand = &cobra.Command{
	Use:     "connect",
	Short:   "Connect to host SSH socket, used by ProxyCommand",
	Example: "Argument is a host.",
	Hidden:  true,
	RunE:    runProxyCommand,
}

// nolint:gochecknoinits
func init() {
	proxyCommand.Flags().BoolP("no-rewrite", "", false, "Do not automatically rewrite outdated configuration")
	proxyCommand.Flags().IntP("port", "p", 0, "SSH destination port")
	proxyCommand.Flags().BoolP("dry-run", "", false, "Only show how assh would connect but don't actually do it")
	_ = viper.BindPFlags(proxyCommand.Flags())
}

func runProxyCommand(cmd *cobra.Command, args []string) error {
	_ = "STUB: not implemented"
	return nil
}

// dry-run option
// Setting the 'ASSH_DRYRUN=1' environment variable,
// so 'assh' can use gateways using sub-SSH commands.

// BeforeConfigWrite

// Save

// AfterConfigWrite

// FIXME: handle complete host with json

// nolint:unparam
func computeHost(dest string, portOverride int, conf *config.Config) (*config.Host, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func expandSSHTokens(tokenized string, host *config.Host) string {
	_ = "STUB: not implemented"
	return ""

	// OpenSSH Token Cheatsheet (stolen directly from the man pages)
	//
	// %%    A literal `%'.
	// %C    Shorthand for %l%h%p%r.
	// %d    Local user's home directory.
	// %h    The remote hostname.
	// %i    The local user ID.
	// %L    The local hostname.
	// %l    The local hostname, including the domain name.
	// %n    The original remote hostname, as given on the command line.
	// %p    The remote port.
	// %r    The remote username.
	// %u    The local username.
}

// TODO: Expansion of strings like "%%C" and "%C" are equivalent due to the
//       order that tokens are evaluated.  Should look at how OpenSSH implements
//       the tokenization behavior.

// Expand a home directory ~.  Assume nobody is using
// the ~otheruser syntax.

func prepareHostControlPath(host *config.Host) error { _ = "STUB: not implemented"; return nil }

func proxy(host *config.Host, conf *config.Config, dryRun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// FIXME: dynamically add "-v" flags

// FIXME: detect ssh client version and use netcat if too old
// for now, the workaround is to configure the ProxyCommand of the host to "nc %h %p"

func proxyDirect(host *config.Host, dryRun bool) error { _ = "STUB: not implemented"; return nil }

func runProxy(host *config.Host, command string, dryRun bool) error {
	_ = "STUB: not implemented"
	return nil
}

// #nosec

func hostPrepare(host *config.Host, gateway string) error { _ = "STUB: not implemented"; return nil }

// FIXME: resolve using custom dns server

// #nosec

type exportReadWrite struct {
	written uint64
	err     error
}

// ConnectionStats contains network and timing informations about a connection
type ConnectionStats struct {
	WrittenBytes            uint64
	WrittenBytesHuman       string
	CreatedAt               time.Time
	ConnectedAt             time.Time
	DisconnectedAt          time.Time
	ConnectionDuration      time.Duration
	ConnectionDurationHuman string
	AverageSpeed            float64
	AverageSpeedHuman       string
}

func (c *ConnectionStats) String() string { _ = "STUB: not implemented"; return "" }

// ConnectHookArgs is the struture sent to the hooks and used in Go templates by the hook drivers
type ConnectHookArgs struct {
	Host  *config.Host
	Stats *ConnectionStats
	Error string
}

func (c ConnectHookArgs) String() string { _ = "STUB: not implemented"; return "" }

func proxyGo(host *config.Host, dryRun bool) error { _ = "STUB: not implemented"; return nil }

// BeforeConnect hook

// use GatewayConnectTimeout, fallback on ConnectTimeout

// set to 0 to disable

// OnConnectError hook

// OnConnect hook

// Ignore SIGHUP

// round duraction

// human

// OnDisconnect hook

func readAndWrite(ctx context.Context, r io.Reader, w io.Writer) <-chan exportReadWrite {
	_ = "STUB: not implemented"
	return nil
}
