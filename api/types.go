package api

import (
	"net"
)

type Writeable interface {
	Strings() []string
}

type JSONInput struct {
	Frontends Frontends `json:"frontends"`
	Backends  Backends  `json:"backends"`
}

type Frontends []Frontend

type Frontend struct {
	Name         string       `json:"name"`
	IP           net.IP       `json:"ip"`
	Port         uint16       `json:"port"`
	BackendRules BackendRules `json:"backend_rules"`
}

type BackendRules []RuleUseBackend

type RuleUseBackend struct {
	Name string `json:"name"`
}

type Backends []Backend

type Backend struct {
	Name        string      `json:"name"`
	BalanceType BalanceType `json:"balance_type"`
	Options     Options     `json:"options"`
	Mode        Mode        `json:"mode"`
	Servers     Servers     `json:"servers"`
}

// BalanceType defines how a backend should balance connections across it's servers
type BalanceType string

const (
	// BalanceTypeRoundRobin balances connections by iterating over the backend servers
	BalanceTypeRoundRobin BalanceType = "roundrobin"
	// BalanceTypeLeastConn chooses the backend server with the least connection
	BalanceTypeLeastConn BalanceType = "leastconn"
)

type Options []Option

// Option is an option for an HAProxy backend
type Option string

const (
	OptionForwardFor Option = "forwardfor"
)

type Servers []Server

// Server represents a HAProxy backend server
type Server struct {
	Name string `json:"name"`
	IP   net.IP `json:"ip"`
	Port uint16 `json:"port"`
}

// Mode is the type of loadbalancer
type Mode string

const (
	ModeHTTP Mode = "http"
	ModeTCP  Mode = "tcp"
)
