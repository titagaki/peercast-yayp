package peercast

import (
	"strings"
)

var ypPrefix = "sp"

const (
	RestrictionTypeNone int = iota
	RestrictionTypeFirewalled
	RestrictionTypeBandWidth
	RestrictionTypeHighBandWidth
)

type StreamOptions struct {
	Prefix          string
	NameSpace       string
	HiddenListeners bool
	RestrictionType int
	Genre           string
}

func ParseGenre(s string) (o StreamOptions, ok bool) {
	if !strings.HasPrefix(s, ypPrefix) {
		return o, false
	}

	o.Prefix = ypPrefix
	pos := len(ypPrefix)

Loop:
	for i := pos; i < len(s); i++ {
		switch c := s[i]; c {
		case ':':
			o.NameSpace = s[pos:i]
			pos = i + 1
		case '?':
			if pos == i {
				o.HiddenListeners = true
				pos++
			}
		case '@':
			if pos == i {
				if o.RestrictionType < RestrictionTypeHighBandWidth {
					o.RestrictionType++
				}
				pos++
			}
		default:
			if !((c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')) {
				break Loop
			}
		}
	}

	o.Genre = s[pos:]

	return o, true
}
