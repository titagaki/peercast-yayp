package peercast

import "encoding/xml"

// StatXML はPeerCastの /admin?cmd=viewxml レスポンスのルート要素。
type StatXML struct {
	XMLName       xml.Name `xml:"peercast"`
	ChannelsFound struct {
		Channel []XMLChannel `xml:"channel"`
	} `xml:"channels_found"`
}

// XMLChannel はPeerCastのチャンネルXML要素。
type XMLChannel struct {
	Name    string `xml:"name,attr"`
	ID      string `xml:"id,attr"`
	Bitrate int    `xml:"bitrate,attr"`
	Type    string `xml:"type,attr"`
	Genre   string `xml:"genre,attr"`
	Desc    string `xml:"desc,attr"`
	URL     string `xml:"url,attr"`
	Uptime  uint   `xml:"uptime,attr"`
	Skip    bool   `xml:"skip,attr"`
	Comment string `xml:"comment,attr"`
	Age     uint   `xml:"age,attr"`
	Bcflags bool   `xml:"bcflags,attr"`
	Hits    struct {
		HostCount  int       `xml:"Hosts,attr"`
		Listeners  int       `xml:"listeners,attr"`
		Relays     int       `xml:"relays,attr"`
		Firewalled int       `xml:"firewalled,attr"`
		Closest    int       `xml:"closest,attr"`
		Furthest   int       `xml:"furthest,attr"`
		Newest     int       `xml:"newest,attr"`
		Host       []XMLHost `xml:"host"`
	} `xml:"hits"`
	Track struct {
		Title   string `xml:"title,attr"`
		Artist  string `xml:"artist,attr"`
		Album   string `xml:"album,attr"`
		Genre   string `xml:"genre,attr"`
		Contact string `xml:"contact,attr"`
	} `xml:"track"`
}

// XMLHost はPeerCastのホスト（ChanHit）XML要素。
type XMLHost struct {
	IP        string `xml:"ip,attr"`
	Hops      int    `xml:"hops,attr"`
	Listeners int    `xml:"listeners,attr"`
	Relays    int    `xml:"relays,attr"`
	Uptime    uint   `xml:"uptime,attr"`
	Push      bool   `xml:"push,attr"`
	Relay     bool   `xml:"relay,attr"`
	Direct    bool   `xml:"direct,attr"`
	Cin       bool   `xml:"cin,attr"`
	Stable    bool   `xml:"stable,attr"`
	Version   uint   `xml:"version,attr"`
	Update    uint   `xml:"update,attr"`
	Tracker   bool   `xml:"tracker,attr"`
}
