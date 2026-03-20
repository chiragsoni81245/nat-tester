package detector

import (
	"github.com/chiragsoni81245/nat-tester/internal/types"
)

func Detect(results []types.Result) types.NATType {
	if len(results) < 2 {
		return types.Unknown
	}

	first := results[0].Addr

	sameIP := true
	samePort := true

	for _, r := range results[1:] {
		if !r.Addr.IP.Equal(first.IP) {
			sameIP = false
		}
		if r.Addr.Port != first.Port {
			samePort = false
		}
	}

	if !samePort {
		return types.SymmetricNAT
	}

	if sameIP && samePort {
		return types.RestrictedCone
	}

	return types.Unknown
}
