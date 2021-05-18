package cli

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	api "github.com/atricore/josso-api-go"
)

func LocationToStr(l *api.LocationDTO) string {

	url := strings.ToLower(*l.Protocol) + "://" + *l.Host

	if *l.Port != 0 && *l.Port != 80 && *l.Port != 443 {
		url = fmt.Sprintf("%s:%d", url, *l.Port)
	}

	if *l.Context != "" {
		url += "/" + *l.Context
	}

	if *l.Uri != "" {
		url += "/" + *l.Uri
	}

	return url
}

func StrToLocation(v string) (*api.LocationDTO, error) {
	// Parse URL
	u, err := url.Parse(v)
	if err != nil {
		return nil, err
	}

	location := api.NewLocationDTO()
	location.Protocol = &u.Scheme

	// Strip port from host
	h := u.Hostname()
	location.Host = &h

	// Get Port
	location.Port, err = StrToPort(u.Port())

	// Default ports
	if *location.Port == 0 {
		switch u.Scheme {
		case "https":
			var p int32 = 443
			location.Port = &p
		default:
			var p int32 = 80
			location.Port = &p
		}

	}

	s := strings.SplitN(u.Path, "/", 3)

	if len(s) > 1 {
		location.Context = &s[1]
	} else {
		v = ""
		location.Context = &v
	}

	if len(s) > 2 {
		location.Uri = &s[2]
	} else {
		v = ""
		location.Uri = &v
	}

	return location, err
}

func StrToPort(v string) (*int32, error) {

	var port int32 = 0
	if v == "" {
		return &port, nil
	}
	i, err := strconv.Atoi(v)
	y := int32(i)
	return &y, err
}
