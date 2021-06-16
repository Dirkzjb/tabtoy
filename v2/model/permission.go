package model

import "strings"

type Perm uint32

const (
	Perm_Client Perm = 1 << iota
	Perm_Server
	Perm_ClientServer = Perm_Client | Perm_Server
)

func (perm Perm) Match(permstr string) bool {
	switch strings.ToLower(strings.TrimSpace(permstr)) {
	case "cs", "sc":
		return perm&Perm_ClientServer != 0
	case "c":
		return perm&Perm_Client != 0
	case "s":
		return perm&Perm_Server != 0
	case "":
		return true
	}
	return false
}
