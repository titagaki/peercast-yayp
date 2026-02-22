package peercast

import "strings"

// ジャンル文字列中の制限タイプ。
const (
	RestrictionNone = iota
	RestrictionFirewalled
	RestrictionBandwidth
	RestrictionHighBandwidth
)

// StreamOptions はジャンル文字列から解析されたストリーミングオプション。
type StreamOptions struct {
	Prefix          string
	Namespace       string
	HiddenListeners bool
	Restriction     int
	Genre           string
}

// ParseGenre はジャンル文字列をパースしてストリーミングオプションを返す。
// ypPrefix で始まらない場合は ok=false を返す。
func ParseGenre(s string, ypPrefix string) (StreamOptions, bool) {
	if !strings.HasPrefix(s, ypPrefix) {
		return StreamOptions{}, false
	}

	o := StreamOptions{Prefix: ypPrefix}
	pos := len(ypPrefix)

loop:
	for i := pos; i < len(s); i++ {
		c := s[i]
		switch {
		case c == ':':
			o.Namespace = s[pos:i]
			pos = i + 1
		case c == '?' && pos == i:
			o.HiddenListeners = true
			pos++
		case c == '@' && pos == i:
			if o.Restriction < RestrictionHighBandwidth {
				o.Restriction++
			}
			pos++
		case (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z'):
			// 有効な文字、続行
		default:
			break loop
		}
	}

	o.Genre = s[pos:]
	return o, true
}
