package config

// HostsMap is a map of **Host).Name -> *Host
type HostsMap map[string]*Host

// HostsList is a list of *Host
type HostsList []*Host

// ToList returns a slice of *Hosts
func (hm *HostsMap) ToList() HostsList { _ = "STUB: not implemented"; return *new(HostsList) }

func (hl HostsList) Len() int           { _ = "STUB: not implemented"; return 0 }
func (hl HostsList) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (hl HostsList) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// SortedList returns a list of hosts sorted by their name
func (hm *HostsMap) SortedList() HostsList { _ = "STUB: not implemented"; return *new(HostsList) }
