package fritzbox

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

type device struct {
	url                 *url.URL
	username            string
	password            string
	iface               string
	challengeToken      string
	authenticationToken string
	sid                 string
}

func NewDevice(rawurl string, username string, password string, iface string) (*device, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return nil, err
	}

	return &device{url, username, password, iface, "", "", ""}, nil
}

func (d *device) Authenticate() error {
	if err := d.requestChallengeToken(); err != nil {
		return err
	}

	d.authenticationToken = createAuthenticationToken(d.challengeToken, d.password)

	if err := d.requestSID(); err != nil {
		return err
	}

	return nil
}

func (d *device) StartCapturing() (io.ReadCloser, error) {
	pathParams := fmt.Sprintf("cgi-bin/capture_notimeout?ifaceorminor=%s&snaplen=&capture=Start&sid=%s", d.iface, d.sid)
	path := path.Join(d.url.Path, pathParams)
	uri := fmt.Sprintf("%s://%s/%s", d.url.Scheme, d.url.Host, path)

	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (d *device) Username() string {
	return d.username
}

func (d *device) Interface() string {
	return d.iface
}

func (d *device) requestChallengeToken() error {
	path := path.Join(d.url.Path, "login_sid.lua")
	uri := fmt.Sprintf("%s://%s/%s", d.url.Scheme, d.url.Host, path)

	xmlBody, err := getUriAndReadBody(uri)
	if err != nil {
		return err
	}

	d.challengeToken, err = parseChallengeToken(xmlBody)
	if err != nil {
		return err
	}

	return nil
}

func (d *device) requestSID() error {
	path := path.Join(d.url.Path, "login_sid.lua")
	uri := fmt.Sprintf("%s://%s/%s", d.url.Scheme, d.url.Host, path)

	response := fmt.Sprintf("%s-%s", d.challengeToken, d.authenticationToken)

	xmlBody, err := postUriAndReadBody(uri, response, d.username)
	if err != nil {
		return err
	}

	d.sid, err = parseSID(xmlBody)
	if err != nil {
		return err
	}

	if d.sid == "0000000000000000" {
		return fmt.Errorf("Authentication for username %s failed", d.username)
	}

	return nil
}
