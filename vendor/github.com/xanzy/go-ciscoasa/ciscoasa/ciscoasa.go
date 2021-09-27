//
// Copyright 2017, Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ciscoasa

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Predefine standard errors
var (
	ErrInternalServer = errors.New("ciscoasa: Internal Server error")
	ErrUnknown        = errors.New("ciscoasa: Unknown error")
)

// Client represents a Cisco ASA client
type Client struct {
	client    *http.Client
	baseURL   *url.URL
	username  string
	password  string
	pageLimit int

	Access      *accessService
	Interfaces  *interfaceService
	Objects     *objectsService
	Routing     *routingService
	DeviceSetup *devicesetupService
	Dhcp        *dhcpService
	Nat         *natService
	Failover    *failoverService
	Licensing   *licenseService
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Messages []*ErrorMessage `json:"messages"`
}

func (e *ErrorResponse) Error() string {
	var errs []string
	for _, m := range e.Messages {
		errs = append(errs, m.Details)
	}
	return strings.Join(errs, "\n")
}

// ErrorMessage represents a single API error
type ErrorMessage struct {
	Level   string `json:"level"`
	Code    string `json:"code"`
	Context string `json:"context,omitempty"`
	Details string `json:"details"`
}

type protocolDefinition struct {
	Prefix    string
	GroupName string
	Kind      string
}

// RangeInfo common data type amongst object types
type RangeInfo struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
	Total  int `json:"total"`
}

// NewClient creates a new client for communicating with a Cisco ASA
func NewClient(apiURL, username, password string, sslNoVerify bool) (*Client, error) {
	baseURL, err := url.Parse(apiURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client: &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				Dial: (&net.Dialer{
					Timeout:   300 * time.Second,
					KeepAlive: 30 * time.Second,
				}).Dial,
				TLSClientConfig:     &tls.Config{InsecureSkipVerify: sslNoVerify},
				TLSHandshakeTimeout: 180 * time.Second,
			},
			Timeout: 300 * time.Second,
		},
		baseURL:   baseURL,
		username:  username,
		password:  password,
		pageLimit: 100,
	}

	c.Access = &accessService{c}
	c.Interfaces = &interfaceService{c}
	c.Objects = &objectsService{c}
	c.Routing = &routingService{c}
	c.DeviceSetup = &devicesetupService{c}
	c.Dhcp = &dhcpService{c}
	c.Nat = &natService{c}
	c.Failover = &failoverService{c}
	c.Licensing = &licenseService{c}

	return c, nil
}

// newRequest creates a HTTP request of given method and data.
func (c *Client) newRequest(method string, api string, v interface{}) (*http.Request, error) {
	var body io.Reader

	if v != nil {
		data, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(data)
	}

	u, err := url.Parse(api)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.baseURL.ResolveReference(u).String(), body)
	if err != nil {
		return nil, err
	}

	req.Close = true
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("User-Agent", "REST API Agent")
	req.SetBasicAuth(c.username, c.password)

	return req, nil
}

// do makes the actual API request.
func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = checkResponse(resp)
	if err != nil {
		return resp, err
	}

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
		}
	}

	return resp, err
}

// CheckResponse checks the API response for errors, and returns them if present.
func checkResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204:
		return nil
	}

	errorResponse := &ErrorResponse{}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		if err := json.Unmarshal(data, errorResponse); err != nil {
			errorResponse.Messages = append(errorResponse.Messages, &ErrorMessage{
				Details: "failed to parse unknown error format",
				Level:   "Error",
			})
		}
	}

	return errorResponse
}

// kindFromValue returns the object kind of the given value.
func kindFromValue(value string) (string, error) {
	// Test if the value references a reserved keyword.
	switch strings.ToLower(value) {
	case "any":
		return "AnyIPAddress", nil
	case "any4":
		return "AnyIPAddress", nil
	case "any6":
		return "AnyIPv6Address", nil
	}

	// Test if the value specifies a range.
	if strings.Contains(value, "-") {
		parts := strings.Split(value, "-")
		from := net.ParseIP(parts[0])
		to := net.ParseIP(parts[1])

		if from.To4() != nil && to.To4() != nil {
			return "IPv4Range", nil
		}
		if from.To16() != nil && to.To16() != nil {
			return "IPv6Range", nil
		}

		return "", fmt.Errorf("value is not a valid range %q", value)
	}

	// Test if the value specifies a CIDR.
	if strings.Contains(value, "/") {
		addr, _, err := net.ParseCIDR(value)
		if err != nil {
			return "", fmt.Errorf("value is not a valid CIDR: %v", err)
		}

		if addr.To4() != nil {
			return "IPv4Network", nil
		}
		if addr.To16() != nil {
			return "IPv6Network", nil
		}

		return "", fmt.Errorf("value is not a valid CIDR %q", value)
	}

	// Test if the value specifies an IP address.
	if addr := net.ParseIP(value); addr != nil {
		if addr.To4() != nil {
			return "IPv4Address", nil
		}
		if addr.To16() != nil {
			return "IPv6Address", nil
		}
	}

	return "", fmt.Errorf("failed to infer kind from value %q", value)
}

var protocolDefinitions = map[string]*protocolDefinition{
	"tcp":             {"tcp", "Tcp", "TcpUdpService"},
	"udp":             {"udp", "Udp", "TcpUdpService"},
	"tcpudp":          {"tcp-udp", "TcpUdp", "TcpUdpService"},
	"icmp":            {"icmp", "Icmp", "ICMPService"},
	"icmp6":           {"icmp", "Icmp6", "ICMP6Service"},
	"ip":              {"ip", "Ip", "IPService"},
	"networkprotocol": {"networkprotocol", "Protocol", "NetworkProtocol"},
}

func protocolFromValue(value string) (*protocolDefinition, error) {
	if p, ok := protocolDefinitions[strings.ToLower(value)]; ok {
		return p, nil
	}

	return nil, fmt.Errorf("protocol %q is not a known protocol", value)
}

func idFromResponse(resp *http.Response) (string, error) {
	parts := strings.Split(resp.Header.Get("Location"), "/")

	loc := parts[len(parts)-1]
	if loc == "" {
		return "", errors.New("failed to retrieve ID from response")
	}

	return loc, nil
}

// Backup represents a backup.
type Backup struct {
	Context    string `json:"context,omitempty"`
	Location   string `json:"location"`
	Passphrase string `json:"passphrase,omitempty"`
}

// CreateBackup creates a backup.
func (c *Client) CreateBackup(context, location, passphrase string) error {
	u := "/api/backup"

	r := &Backup{
		Context:    context,
		Location:   location,
		Passphrase: passphrase,
	}

	req, err := c.newRequest("POST", u, r)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)

	return err
}
