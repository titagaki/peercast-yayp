package peercast

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

type StatXML struct {
	ChannelsFound ChannelsFound `xml:"channels_found"`
}

type ChannelsFound struct {
	Channel []Channel `xml:"channel"`
}

type Channel struct {
	Name       string `xml:"name,attr"`
	ID         string `xml:"id,attr"`
	Bitrate    int    `xml:"bitrate,attr"`
	Type       string `xml:"type,attr"`
	Genre      string `xml:"genre,attr"`
	Desc       string `xml:"desc,attr"`
	Url        string `xml:"url,attr"`
	Uptime     uint   `xml:"uptime,attr"`
	Skip       bool   `xml:"skip,attr"`
	Comment    string `xml:"comment,attr"`
	Age        int    `xml:"age,attr"`
	Bcflags    bool   `xml:"bcflags,attr"`
	ChanHitStat ChanHitStat `xml:"hits"`
	ChannelTrack ChannelTrack `xml:"track"`
}

type ChanHitStat struct {
	Hosts      int    `xml:"Hosts,attr"`
	Listeners  int    `xml:"listeners,attr"`
	Relays     int    `xml:"relays,attr"`
	Firewalled int    `xml:"firewalled,attr"`
	Closest    int    `xml:"closest,attr"`
	Furthest   int    `xml:"furthest,attr"`
	Newest     int    `xml:"newest,attr"`
	ChanHit    []ChanHit `xml:"host"`
}

type ChanHit struct {
	Ip         string `xml:"ip,attr"`
	Hops       int    `xml:"hops,attr"`
	Listeners  int    `xml:"listeners,attr"`
	Relays     int    `xml:"relays,attr"`
	Uptime     uint   `xml:"uptime,attr"`
	Push       bool   `xml:"push,attr"`
	Relay      bool   `xml:"relay,attr"`
	Direct     bool   `xml:"direct,attr"`
	Cin        bool   `xml:"cin,attr"`
	Stable     bool   `xml:"stable,attr"`
	Version    uint   `xml:"version,attr"`
	Update     string `xml:"update,attr"`
	Tracker    string `xml:"tracker,attr"`
}

type ChannelTrack struct {
	Title      string `xml:"title,attr"`
	Artist     string `xml:"artist,attr"`
	Album      string `xml:"album,attr"`
	Genre      string `xml:"genre,attr"`
	Contact    string `xml:"contact,attr"`
}

func GetStatXML() (*StatXML, error) {
	resp, err := requestViewStatXML()
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	var data StatXML
	if err := xml.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, errors.Wrap(err,"XML Unmarshal error")
	}

	return &data, nil
}