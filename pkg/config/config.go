package config

import (
	"io"
)

var asshBinaryPath = "assh"

const defaultSSHConfigPath = "~/.ssh/config"

// Config contains a list of Hosts sections and a Defaults section representing a configuration file
type Config struct {
	Hosts             HostsMap `yaml:"hosts,omitempty,flow" json:"hosts"`
	Templates         HostsMap `yaml:"templates,omitempty,flow" json:"templates"`
	Defaults          Host     `yaml:"defaults,omitempty,flow" json:"defaults,omitempty"`
	Includes          []string `yaml:"includes,omitempty,flow" json:"includes,omitempty"`
	ASSHKnownHostFile string   `yaml:"asshknownhostfile,omitempty,flow" json:"asshknownhostfile,omitempty"`
	ASSHBinaryPath    string   `yaml:"asshbinarypath,omitempty,flow" json:"asshbinarypath,omitempty"`

	includedFiles map[string]bool
	sshConfigPath string
}

// DisableAutomaticRewrite will configure the ~/.ssh/config file to not automatically rewrite the configuration file
func (c *Config) DisableAutomaticRewrite() { _ = "STUB: not implemented"; return }

// SetASSHBinaryPath sets the default assh binary path
// this value may be overwritten in the assh.yml file using the asshbinarypath variable
func SetASSHBinaryPath(path string) { _ = "STUB: not implemented"; return }

// String returns the JSON output
func (c *Config) String() string { _ = "STUB: not implemented"; return "" }

// SaveNewKnownHost registers the target as a new known host and save the full known hosts list on disk
func (c *Config) SaveNewKnownHost(target string) { _ = "STUB: not implemented"; return }

func (c *Config) addKnownHost(target string) { _ = "STUB: not implemented"; return }

// KnownHostsFileExists returns nil if it the file exists and an error if it doesn't
func (c *Config) KnownHostsFileExists() error { _ = "STUB: not implemented"; return nil }

// LoadKnownHosts loads known hosts list from disk
func (c *Config) LoadKnownHosts() error { _ = "STUB: not implemented"; return nil }

// IncludedFiles returns the list of the included files
func (c *Config) IncludedFiles() []string { _ = "STUB: not implemented"; return nil }

// JSONString returns a string representing the JSON of a Config object
func (c *Config) JSONString() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// computeHost returns a copy of the host with applied defaults, resolved inheritances and configured internal fields
func computeHost(host *Host, config *Config, name string, fullCompute bool) (*Host, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// name internal field

// self is already inherited

// Inheritance
// FIXME: allow deeper inheritance:
//     currently not resolving inherited hosts
//     we should resolve all inherited hosts and pass the
//     currently resolved hosts to avoid computing an host twice

// fullCompute applies config.Defaults
// config.Defaults should be applied when proxying
// but should not when exporting .ssh/config file

// apply defaults based on "Host *"

// expands variables in host
// i.e: %h.some.zone -> {name}.some.zone

// ssh resolve '%h' in hostnames
// -> we bypass the string expansion if the input matches
//    an already resolved hostname
// See https://github.com/moul/assh/issues/103

func (c *Config) getHostByName(name string, safe bool, compute bool, allowTemplate bool) (*Host, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Config) getHostByPath(path string, safe bool, compute bool, allowTemplate bool) (*Host, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetGatewaySafe returns gateway Host configuration, a gateway is like a Host, except, the host path is not resolved
func (c *Config) GetGatewaySafe(name string) *Host { _ = "STUB: not implemented"; return nil }

// FIXME: fullCompute for gateway ?

// GetHost returns a matching host form Config hosts list
func (c *Config) GetHost(name string) (*Host, error) { _ = "STUB: not implemented"; return nil, nil }

// GetHostSafe won't fail, in case the host is not found, it will returns a virtual host matching the pattern
func (c *Config) GetHostSafe(name string) *Host { _ = "STUB: not implemented"; return nil }

// isSSHConfigOutdated returns true if assh.yml or an included file has a
// modification date more recent than .ssh/config
func (c *Config) isSSHConfigOutdated() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// IsConfigOutdated returns true if .ssh/config needs to be rebuild.
// The reason may be:
// - assh.yml (or an included file) was updated recently
// - <target> matches a regex and was never seen before (not present in known-hosts file)
func (c *Config) IsConfigOutdated(target string) (bool, error) {
	_ = "STUB: not implemented"
	// check if the target is a regex and if the pattern
	// was never matched before (not in known hosts)
	return false, nil
}

// check if the ~/.ssh/config file is older than assh.yml or any included file

// needsARebuildForTarget returns true if the .ssh/config file needs to be rebuild for a specific target
func (c *Config) needsARebuildForTarget(target string) bool {
	_ = "STUB: not implemented"
	return false
}

// compute lists

// check for direct hostname matching

// check for direct alias matching

// check for pattern matching

// LoadConfig loads the content of an io.Reader source
func (c *Config) LoadConfig(source io.Reader) error { _ = "STUB: not implemented"; return nil }

func (c *Config) mergeWildCardEntries() { _ = "STUB: not implemented"; return }

// if * is in the middle

// if the wildcard matches

// if the wildcard matches

func (c *Config) applyMissingNames() { _ = "STUB: not implemented"; return }

// should be removed

// should be removed

// SaveSSHConfig saves the configuration to ~/.ssh/config
func (c *Config) SaveSSHConfig() error { _ = "STUB: not implemented"; return nil }

// validate hosts

// LoadFile loads the content of a configuration file in the Config object
func (c *Config) LoadFile(filename string) error { _ = "STUB: not implemented"; return nil }

// Resolve '~' and '$HOME'

// Anti-loop protection

// Read file

// Load config stream

// Successful loading

// Handling includes

// LoadFiles will try to glob the pattern and load each matching entries
func (c *Config) LoadFiles(pattern string) error {
	_ = "STUB: not implemented"
	// Resolve '~' and '$HOME'
	return nil
}

// Globbing

// Load files iteratively

// sortedNames returns the host names sorted alphabetically
func (c *Config) sortedNames() []string { _ = "STUB: not implemented"; return nil }

// Validate checks for values errors
func (c *Config) Validate() []error { _ = "STUB: not implemented"; return nil }

// ValidateSummary summaries Validate() errors slice
func (c *Config) ValidateSummary() error { _ = "STUB: not implemented"; return nil }

// WriteSSHConfigTo returns a .ssh/config valid file containing assh configuration
func (c *Config) WriteSSHConfigTo(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// FIXME: add version

// SSHConfigPath returns the ~/.ssh/config file path
func (c *Config) SSHConfigPath() string { _ = "STUB: not implemented"; return "" }

// New returns an instantiated Config object
func New() *Config { _ = "STUB: not implemented"; return nil }

// Open parses a configuration file and returns a *Config object
func Open(path string) (*Config, error) { _ = "STUB: not implemented"; return nil, nil }
