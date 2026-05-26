package config

import (
	"io"

	composeyaml "github.com/docker/libcompose/yaml"
)

// Host defines the configuration flags of a host
type Host struct {
	// ssh-config fields
	AddKeysToAgent                   string                    `yaml:"addkeystoagent,omitempty,flow" json:"AddKeysToAgent,omitempty"`
	AddressFamily                    string                    `yaml:"addressfamily,omitempty,flow" json:"AddressFamily,omitempty"`
	AskPassGUI                       string                    `yaml:"askpassgui,omitempty,flow" json:"AskPassGUI,omitempty"`
	BatchMode                        string                    `yaml:"batchmode,omitempty,flow" json:"BatchMode,omitempty"`
	BindAddress                      string                    `yaml:"bindaddress,omitempty,flow" json:"BindAddress,omitempty"`
	CanonicalDomains                 string                    `yaml:"canonicaldomains,omitempty,flow" json:"CanonicalDomains,omitempty"`
	CanonicalizeFallbackLocal        string                    `yaml:"canonicalizefallbacklocal,omitempty,flow" json:"CanonicalizeFallbackLocal,omitempty"`
	CanonicalizeHostname             string                    `yaml:"canonicalizehostname,omitempty,flow" json:"CanonicalizeHostname,omitempty"`
	CanonicalizeMaxDots              string                    `yaml:"canonicalizemaxDots,omitempty,flow" json:"CanonicalizeMaxDots,omitempty"`
	CanonicalizePermittedCNAMEs      string                    `yaml:"canonicalizepermittedcnames,omitempty,flow" json:"CanonicalizePermittedCNAMEs,omitempty"`
	CASignatureAlgorithms            composeyaml.Stringorslice `yaml:"casignaturealgorithms,omitempty,flow" json:"CASignatureAlgorithms,omitempty"`
	CertificateFile                  composeyaml.Stringorslice `yaml:"certificatefile,omitempty,flow" json:"CertificateFile,omitempty"`
	ChallengeResponseAuthentication  string                    `yaml:"challengeresponseauthentication,omitempty,flow" json:"ChallengeResponseAuthentication,omitempty"`
	CheckHostIP                      string                    `yaml:"checkhostip,omitempty,flow" json:"CheckHostIP,omitempty"`
	Cipher                           string                    `yaml:"cipher,omitempty,flow" json:"Cipher,omitempty"`
	Ciphers                          composeyaml.Stringorslice `yaml:"ciphers,omitempty,flow" json:"Ciphers,omitempty"`
	ClearAllForwardings              string                    `yaml:"clearallforwardings,omitempty,flow" json:"ClearAllForwardings,omitempty"`
	Compression                      string                    `yaml:"compression,omitempty,flow" json:"Compression,omitempty"`
	CompressionLevel                 int                       `yaml:"compressionlevel,omitempty,flow" json:"CompressionLevel,omitempty"`
	ConnectionAttempts               string                    `yaml:"connectionattempts,omitempty,flow" json:"ConnectionAttempts,omitempty"`
	ConnectTimeout                   int                       `yaml:"connecttimeout,omitempty,flow" json:"ConnectTimeout,omitempty"`
	ControlMaster                    string                    `yaml:"controlmaster,omitempty,flow" json:"ControlMaster,omitempty"`
	ControlPath                      string                    `yaml:"controlpath,omitempty,flow" json:"ControlPath,omitempty"`
	ControlPersist                   string                    `yaml:"controlpersist,omitempty,flow" json:"ControlPersist,omitempty"`
	DynamicForward                   composeyaml.Stringorslice `yaml:"dynamicforward,omitempty,flow" json:"DynamicForward,omitempty"`
	EnableSSHKeysign                 string                    `yaml:"enablesshkeysign,omitempty,flow" json:"EnableSSHKeysign,omitempty"`
	EscapeChar                       string                    `yaml:"escapechar,omitempty,flow" json:"EscapeChar,omitempty"`
	ExitOnForwardFailure             string                    `yaml:"exitonforwardfailure,omitempty,flow" json:"ExitOnForwardFailure,omitempty"`
	FingerprintHash                  string                    `yaml:"fingerprinthash,omitempty,flow" json:"FingerprintHash,omitempty"`
	ForwardAgent                     string                    `yaml:"forwardagent,omitempty,flow" json:"ForwardAgent,omitempty"`
	ForwardX11                       string                    `yaml:"forwardx11,omitempty,flow" json:"ForwardX11,omitempty"`
	ForwardX11Timeout                int                       `yaml:"forwardx11timeout,omitempty,flow" json:"ForwardX11Timeout,omitempty"`
	ForwardX11Trusted                string                    `yaml:"forwardx11trusted,omitempty,flow" json:"ForwardX11Trusted,omitempty"`
	GatewayPorts                     string                    `yaml:"gatewayports,omitempty,flow" json:"GatewayPorts,omitempty"`
	GlobalKnownHostsFile             composeyaml.Stringorslice `yaml:"globalknownhostsfile,omitempty,flow" json:"GlobalKnownHostsFile,omitempty"`
	GSSAPIAuthentication             string                    `yaml:"gssapiauthentication,omitempty,flow" json:"GSSAPIAuthentication,omitempty"`
	GSSAPIClientIdentity             string                    `yaml:"gssapiclientidentity,omitempty,flow" json:"GSSAPIClientIdentity,omitempty"`
	GSSAPIDelegateCredentials        string                    `yaml:"gssapidelegatecredentials,omitempty,flow" json:"GSSAPIDelegateCredentials,omitempty"`
	GSSAPIKeyExchange                string                    `yaml:"gssapikeyexchange,omitempty,flow" json:"GSSAPIKeyExchange,omitempty"`
	GSSAPIRenewalForcesRekey         string                    `yaml:"gssapirenewalforcesrekey,omitempty,flow" json:"GSSAPIRenewalForcesRekey,omitempty"`
	GSSAPIServerIdentity             string                    `yaml:"gssapiserveridentity,omitempty,flow" json:"GSSAPIServerIdentity,omitempty"`
	GSSAPITrustDNS                   string                    `yaml:"gssapitrustdns,omitempty,flow" json:"GSSAPITrustDNS,omitempty"`
	HashKnownHosts                   string                    `yaml:"hashknownhosts,omitempty,flow" json:"HashKnownHosts,omitempty"`
	HostbasedAuthentication          string                    `yaml:"hostbasedauthentication,omitempty,flow" json:"HostbasedAuthentication,omitempty"`
	HostbasedKeyTypes                string                    `yaml:"hostbasedkeytypes,omitempty,flow" json:"HostbasedKeyTypes,omitempty"`
	HostKeyAlgorithms                composeyaml.Stringorslice `yaml:"hostkeyalgorithms,omitempty,flow" json:"HostKeyAlgorithms,omitempty"`
	HostKeyAlias                     string                    `yaml:"hostkeyalias,omitempty,flow" json:"HostKeyAlias,omitempty"`
	IdentitiesOnly                   string                    `yaml:"identitiesonly,omitempty,flow" json:"IdentitiesOnly,omitempty"`
	IdentityAgent                    string                    `yaml:"identityagent,omitempty,flow" json:"IdentityAgent,omitempty"`
	IdentityFile                     composeyaml.Stringorslice `yaml:"identityfile,omitempty,flow" json:"IdentityFile,omitempty"`
	IgnoreUnknown                    string                    `yaml:"ignoreunknown,omitempty,flow" json:"IgnoreUnknown,omitempty"`
	IPQoS                            composeyaml.Stringorslice `yaml:"ipqos,omitempty,flow" json:"IPQoS,omitempty"`
	KbdInteractiveAuthentication     string                    `yaml:"kbdinteractiveauthentication,omitempty,flow" json:"KbdInteractiveAuthentication,omitempty"`
	KbdInteractiveDevices            composeyaml.Stringorslice `yaml:"kbdinteractivedevices,omitempty,flow" json:"KbdInteractiveDevices,omitempty"`
	KexAlgorithms                    composeyaml.Stringorslice `yaml:"kexalgorithms,omitempty,flow" json:"KexAlgorithms,omitempty"`
	KeychainIntegration              string                    `yaml:"keychainintegration,omitempty,flow" json:"KeychainIntegration,omitempty"`
	LocalCommand                     string                    `yaml:"localcommand,omitempty,flow" json:"LocalCommand,omitempty"`
	RemoteCommand                    string                    `yaml:"remotecommand,omitempty,flow" json:"RemoteCommand,omitempty"`
	LocalForward                     composeyaml.Stringorslice `yaml:"localforward,omitempty,flow" json:"LocalForward,omitempty"`
	LogLevel                         string                    `yaml:"loglevel,omitempty,flow" json:"LogLevel,omitempty"`
	MACs                             composeyaml.Stringorslice `yaml:"macs,omitempty,flow" json:"MACs,omitempty"`
	Match                            string                    `yaml:"match,omitempty,flow" json:"Match,omitempty"`
	NoHostAuthenticationForLocalhost string                    `yaml:"nohostauthenticationforlocalhost,omitempty,flow" json:"NoHostAuthenticationForLocalhost,omitempty"`
	NumberOfPasswordPrompts          string                    `yaml:"numberofpasswordprompts,omitempty,flow" json:"NumberOfPasswordPrompts,omitempty"`
	PasswordAuthentication           string                    `yaml:"passwordauthentication,omitempty,flow" json:"PasswordAuthentication,omitempty"`
	PermitLocalCommand               string                    `yaml:"permitlocalcommand,omitempty,flow" json:"PermitLocalCommand,omitempty"`
	PKCS11Provider                   string                    `yaml:"pkcs11provider,omitempty,flow" json:"PKCS11Provider,omitempty"`
	Port                             string                    `yaml:"port,omitempty,flow" json:"Port,omitempty"`
	PreferredAuthentications         string                    `yaml:"preferredauthentications,omitempty,flow" json:"PreferredAuthentications,omitempty"`
	Protocol                         composeyaml.Stringorslice `yaml:"protocol,omitempty,flow" json:"Protocol,omitempty"`
	ProxyJump                        string                    `yaml:"proxyjump,omitempty,flow" json:"ProxyJump,omitempty"`
	ProxyUseFdpass                   string                    `yaml:"proxyusefdpass,omitempty,flow" json:"ProxyUseFdpass,omitempty"`
	PubkeyAcceptedAlgorithms         string                    `yaml:"pubkeyacceptedalgorithms,omitempty,flow" json:"PubkeyAcceptedAlgorithms,omitempty"`
	PubkeyAcceptedKeyTypes           string                    `yaml:"pubkeyacceptedkeytypes,omitempty,flow" json:"PubkeyAcceptedKeyTypes,omitempty"`
	PubkeyAuthentication             string                    `yaml:"pubkeyauthentication,omitempty,flow" json:"PubkeyAuthentication,omitempty"`
	RekeyLimit                       string                    `yaml:"rekeylimit,omitempty,flow" json:"RekeyLimit,omitempty"`
	RemoteForward                    composeyaml.Stringorslice `yaml:"remoteforward,omitempty,flow" json:"RemoteForward,omitempty"`
	RequestTTY                       string                    `yaml:"requesttty,omitempty,flow" json:"RequestTTY,omitempty"`
	RevokedHostKeys                  string                    `yaml:"revokedhostkeys,omitempty,flow" json:"RevokedHostKeys,omitempty"`
	RhostsRSAAuthentication          string                    `yaml:"rhostsrsaauthentication,omitempty,flow" json:"RhostsRSAAuthentication,omitempty"`
	RSAAuthentication                string                    `yaml:"rsaauthentication,omitempty,flow" json:"RSAAuthentication,omitempty"`
	SendEnv                          composeyaml.Stringorslice `yaml:"sendenv,omitempty,flow" json:"SendEnv,omitempty"`
	SetEnv                           composeyaml.Stringorslice `yaml:"setenv,omitempty,flow" json:"SetEnv,omitempty"`
	ServerAliveCountMax              int                       `yaml:"serveralivecountmax,omitempty,flow" json:"ServerAliveCountMax,omitempty"`
	ServerAliveInterval              int                       `yaml:"serveraliveinterval,omitempty,flow" json:"ServerAliveInterval,omitempty"`
	StreamLocalBindMask              string                    `yaml:"streamlocalbindmask,omitempty,flow" json:"StreamLocalBindMask,omitempty"`
	StreamLocalBindUnlink            string                    `yaml:"streamlocalbindunlink,omitempty,flow" json:"StreamLocalBindUnlink,omitempty"`
	StrictHostKeyChecking            string                    `yaml:"stricthostkeychecking,omitempty,flow" json:"StrictHostKeyChecking,omitempty"`
	TCPKeepAlive                     string                    `yaml:"tcpkeepalive,omitempty,flow" json:"TCPKeepAlive,omitempty"`
	Tunnel                           string                    `yaml:"tunnel,omitempty,flow" json:"Tunnel,omitempty"`
	TunnelDevice                     string                    `yaml:"tunneldevice,omitempty,flow" json:"TunnelDevice,omitempty"`
	UpdateHostKeys                   string                    `yaml:"updatehostkeys,omitempty,flow" json:"UpdateHostKeys,omitempty"`
	UseKeychain                      string                    `yaml:"usekeychain,omitempty,flow" json:"UseKeychain,omitempty"`
	UsePrivilegedPort                string                    `yaml:"useprivilegedport,omitempty,flow" json:"UsePrivilegedPort,omitempty"`
	User                             string                    `yaml:"user,omitempty,flow" json:"User,omitempty"`
	UserKnownHostsFile               composeyaml.Stringorslice `yaml:"userknownhostsfile,omitempty,flow" json:"UserKnownHostsFile,omitempty"`
	VerifyHostKeyDNS                 string                    `yaml:"verifyhostkeydns,omitempty,flow" json:"VerifyHostKeyDNS,omitempty"`
	VisualHostKey                    string                    `yaml:"visualhostkey,omitempty,flow" json:"VisualHostKey,omitempty"`
	XAuthLocation                    string                    `yaml:"xauthlocation,omitempty,flow" json:"XAuthLocation,omitempty"`

	// ssh-config fields with a different behavior
	HostName     string `yaml:"hostname,omitempty,flow" json:"HostName,omitempty"`
	ProxyCommand string `yaml:"proxycommand,omitempty,flow" json:"ProxyCommand,omitempty"`

	// exposed assh fields
	Inherits              composeyaml.Stringorslice `yaml:"inherits,omitempty,flow" json:"Inherits,omitempty"`
	Gateways              composeyaml.Stringorslice `yaml:"gateways,omitempty,flow" json:"Gateways,omitempty"`
	ResolveNameservers    composeyaml.Stringorslice `yaml:"resolvenameservers,omitempty,flow" json:"ResolveNameservers,omitempty"`
	ResolveCommand        string                    `yaml:"resolvecommand,omitempty,flow" json:"ResolveCommand,omitempty"`
	ControlMasterMkdir    string                    `yaml:"controlmastermkdir,omitempty,flow" json:"ControlMasterMkdir,omitempty"`
	Aliases               composeyaml.Stringorslice `yaml:"aliases,omitempty,flow" json:"Aliases,omitempty"`
	Hooks                 *HostHooks                `yaml:"hooks,omitempty,flow" json:"Hooks,omitempty"`
	Comment               composeyaml.Stringorslice `yaml:"comment,omitempty,flow" json:"Comment,omitempty"`
	RateLimit             string                    `yaml:"ratelimit,omitempty,flow" json:"RateLimit,omitempty"`
	GatewayConnectTimeout int                       `yaml:"gatewayconnecttimeout,omitempty,flow" json:"GatewayConnectTimeout,omitempty"`

	// private assh fields
	noAutomaticRewrite bool
	knownHosts         []string
	pattern            string
	name               string
	inputName          string
	isDefault          bool
	isTemplate         bool
	inherited          map[string]bool
}

// NewHost returns a host with name
func NewHost(name string) *Host { _ = "STUB: not implemented"; return nil }

// Validate checks for values errors
func (h *Host) Validate() []error { _ = "STUB: not implemented"; return nil }

// String returns the JSON output
func (h *Host) String() string { _ = "STUB: not implemented"; return "" }

// Prototype returns a prototype representation of the host, used in listings
func (h *Host) Prototype() string { _ = "STUB: not implemented"; return "" }

// Name returns the name of a host
func (h *Host) Name() string {
	_ = "STUB: not implemented"

	// RawName returns the raw name of a host without pattern computing
	return ""
}

func (h *Host) RawName() string {
	_ = "STUB: not implemented"

	// Clone returns a copy of an existing Host
	return ""
}

func (h *Host) Clone() *Host { _ = "STUB: not implemented"; return nil }

// Matches returns true if the host matches a given string
func (h *Host) Matches(needle string) bool { _ = "STUB: not implemented"; return false }

// Options returns a map of set options
// nolint:gocyclo
func (h *Host) Options() OptionsList { _ = "STUB: not implemented"; return *new(OptionsList) }

// ssh-config fields

// ssh-config fields with a different behavior
// HostName
// ProxyCommand

// exposed assh fields
// Inherits
// Gateways
// ResolveNameservers
// ResolveCommand
// ControlMasterMkdir
// Aliases
// Comment
// Hooks

// private assh fields
// knownHosts
// pattern
// name
// inputName
// isDefault
// isTemplate
// inherited

// minimal host preparation after loading
func (h *Host) prepare() { _ = "STUB: not implemented"; return }

// ApplyDefaults ensures a Host is valid by filling the missing fields with defaults
// nolint:gocyclo
func (h *Host) ApplyDefaults(defaults *Host) {
	_ = "STUB: not implemented"
	// ssh-config fields
	return
}

// h.CompressionLevel = utils.ExpandField(h.CompressionLevel)

// h.ConnectTimeout = utils.ExpandField(h.ConnectTimeout)

// h.ForwardX11Timeout = utils.ExpandField(h.ForwardX11Timeout)

// h.ServerAliveCountMax = utils.ExpandField(h.ServerAliveCountMax)

// h.ServerAliveInterval = utils.ExpandField(h.ServerAliveInterval)

// ssh-config fields with a different behavior

// exposed assh fields

// h.ResolveNameservers = utils.ExpandField(h.ResolveNameservers)

// h.Gateways = utils.ExpandField(h.Gateways)

// h.Inherits = utils.ExpandField(h.Inherits)

// private assh fields
// h.inherited = make(map[string]bool, 0)

// Extra defaults

// AddKnownHost append target to the host' known hosts list
func (h *Host) AddKnownHost(target string) { _ = "STUB: not implemented"; return }

// WriteSSHConfigTo writes an ~/.ssh/config file compatible host definition to a writable stream
// nolint:gocyclo
func (h *Host) WriteSSHConfigTo(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// FIXME: skip complex patterns

// ssh-config fields

// ssh-config fields with a different behavior

// assh fields

// ExpandString replaces elements in a format string with host variables.
func (h *Host) ExpandString(input string, gateway string) string {
	_ = "STUB: not implemented"

	// name of the host in config
	return ""
}

// original target host name specified on the command line

// target host name

// port

// gateway

// FIXME: add
//   %L -> first component of the local host name
//   %l -> local host name
//   %r -> remote login username
//   %u -> username of the user running assh
//   %r -> remote login username
