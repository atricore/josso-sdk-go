package cli

import (
	"testing"

	api "github.com/atricore/josso-api-go"
)

type stringToLocationTest struct {
	name    string
	value   string
	scheme  string
	port    int32
	host    string
	context string
	uri     string
}

type locationToStringTest struct {
	name  string
	value api.LocationDTO
	url   string
}

var strToLocTests = []stringToLocationTest{
	{
		name:    "StrToLoc-1",
		value:   "http://localhost",
		scheme:  "http",
		host:    "localhost",
		port:    80,
		context: "",
		uri:     "",
	},
	{
		name:    "StrToLoc-2",
		value:   "https://localhost/IDBUS/IDA-A",
		scheme:  "https",
		host:    "localhost",
		port:    443,
		context: "IDBUS",
		uri:     "IDA-A",
	},
	{
		name:    "StrToLoc-3",
		value:   "https://localhost/IDBUS/IDA-A/IDP-1",
		scheme:  "https",
		host:    "localhost",
		port:    443,
		context: "IDBUS",
		uri:     "IDA-A/IDP-1",
	},
	{
		name:    "StrToLoc-4",
		value:   "HTTPS://mycompany.com:8443/IDBUS/IDA-A/IDP-1",
		scheme:  "https",
		host:    "mycompany.com",
		port:    8443,
		context: "IDBUS",
		uri:     "IDA-A/IDP-1",
	},
	{
		name:    "StrToLoc-5",
		value:   "HTTP://localhost:8081",
		scheme:  "http",
		host:    "localhost",
		port:    8081,
		context: "",
		uri:     "",
	},
}

var locToStrTests = []locationToStringTest{
	{
		name: "LocToStr-1",
		value: api.LocationDTO{
			Protocol: api.PtrString("http"),
			Host:     api.PtrString("localhost"),
			Port:     api.PtrInt32(80),
			Context:  api.PtrString(""),
			Uri:      api.PtrString(""),
		},
		url: "http://localhost",
	},
	{
		name: "LocToStr-2",
		value: api.LocationDTO{
			Protocol: api.PtrString("http"),
			Host:     api.PtrString("localhost"),
			Port:     api.PtrInt32(81),
			Context:  api.PtrString(""),
			Uri:      api.PtrString(""),
		},
		url: "http://localhost:81",
	},
	{
		name: "LocToStr-3",
		value: api.LocationDTO{
			Protocol: api.PtrString("http"),
			Host:     api.PtrString("localhost"),
			Port:     api.PtrInt32(80),
			Context:  api.PtrString("IDBUS"),
			Uri:      api.PtrString(""),
		},
		url: "http://localhost/IDBUS",
	},
	{
		name: "LocToStr-4",
		value: api.LocationDTO{
			Protocol: api.PtrString("https"),
			Host:     api.PtrString("localhost"),
			Port:     api.PtrInt32(443),
			Context:  api.PtrString(""),
			Uri:      api.PtrString(""),
		},
		url: "https://localhost",
	},
	{
		name: "LocToStr-5",
		value: api.LocationDTO{
			Protocol: api.PtrString("https"),
			Host:     api.PtrString("localhost"),
			Port:     api.PtrInt32(443),
			Context:  api.PtrString("IDBUS"),
			Uri:      api.PtrString("IDA-1/IDP-1"),
		},
		url: "https://localhost/IDBUS/IDA-1/IDP-1",
	},
	{
		name: "LocToStr-6",
		value: api.LocationDTO{
			Protocol: api.PtrString("https"),
			Host:     api.PtrString("localhost"),
			Port:     api.PtrInt32(8443),
			Context:  api.PtrString("IDBUS"),
			Uri:      api.PtrString("IDA-1/IDP-1"),
		},
		url: "https://localhost:8443/IDBUS/IDA-1/IDP-1",
	},
}

func TestStringToLocation(t *testing.T) {

	for _, lt := range strToLocTests {
		t.Run(
			lt.name,
			func(t *testing.T) {
				l, err := StrToLocation(lt.value)
				if err != nil {
					t.Errorf("Error in location %s, [%v]", lt.value, err)
					return
				}
				if l == nil {
					t.Errorf("Received nil location for %s", lt.value)
					return
				}

				if *l.Protocol != lt.scheme {
					t.Errorf("Invalid scheme %s, expected %s", *l.Protocol, lt.scheme)
				}
				if *l.Port != lt.port {
					t.Errorf("Invalid port %d, port %d", *l.Port, lt.port)
				}
				if *l.Host != lt.host {
					t.Errorf("Invalid host %s, expected %s", *l.Host, lt.host)
				}
				if *l.Context != lt.context {
					t.Errorf("Invalid context %s, expected %s", *l.Context, lt.context)
				}
				if *l.Uri != lt.uri {
					t.Errorf("Invalid path %s, expected %s", *l.Uri, lt.uri)
				}
			},
		)
	}

}

func TestLocationToString(t *testing.T) {
	for _, lt := range locToStrTests {
		t.Run(
			lt.name,
			func(t *testing.T) {
				l := LocationToStr(&lt.value)

				if l != lt.url {
					t.Errorf("Invalid URL %s, expected %s", l, lt.url)
				}
			},
		)
	}
}
