package config

// Option is an host option
type Option struct {
	Name  string
	Value string
}

// OptionsList is a list of options
type OptionsList []Option

func (o *Option) String() string { _ = "STUB: not implemented"; return "" }

// Get returns the option value matching the name or "" if the key is not found
func (ol *OptionsList) Get(name string) string { _ = "STUB: not implemented"; return "" }

// ToStringList returns a list of string with the following format: `key=value`
func (ol *OptionsList) ToStringList() []string { _ = "STUB: not implemented"; return nil }

// Remove removes an option from the list based on its key
func (ol *OptionsList) Remove(key string) { _ = "STUB: not implemented"; return }
