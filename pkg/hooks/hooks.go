package hooks

import (
	composeyaml "github.com/docker/libcompose/yaml"
)

// Hooks represents a slice of Hook
type Hooks composeyaml.Stringorslice

// HookDriver represents a hook driver
type HookDriver interface {
	Run(RunArgs) error
	Close() error
}

// HookDrivers represents a slice of HookDriver
type HookDrivers []HookDriver

// RunArgs is a map of interface{}
type RunArgs interface{}

// InvokeAll calls all hooks
func (h *Hooks) InvokeAll(args RunArgs) (HookDrivers, error) {
	_ = "STUB: not implemented"
	return *new(HookDrivers), nil
}

// Close closes all hook drivers and returns a slice of errs
func (hd *HookDrivers) Close() []error { _ = "STUB: not implemented"; return nil }

// New returns an HookDriver instance
func New(expr string) (HookDriver, error) { _ = "STUB: not implemented"; return *new(HookDriver), nil }
