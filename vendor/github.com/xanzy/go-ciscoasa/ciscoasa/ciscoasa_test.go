//
// Copyright 2017, Rutger te Nijenhuis & Sander van Harmelen
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
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setup() (*http.ServeMux, *httptest.Server, *Client) {
	// create a HTTP multiplexer to use with the test server
	mux := http.NewServeMux()
	// server is a test HTTP server used to provide mock API responses
	server := httptest.NewServer(mux)
	// client is the CiscoASA client being tested
	client, _ := NewClient(server.URL, "", "", false)

	return mux, server, client
}

func teardown(server *httptest.Server) {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %s, want %s", got, want)
	}
}

type values map[string]interface{}

func testJSONBody(t *testing.T, r *http.Request, want values) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}

	var got values
	json.Unmarshal(b, &got)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Request parameters: %v, want %v", got, want)
	}
}

func TestKindFromValue(t *testing.T) {
	var cases = []struct {
		addr string
		kind string
		err  bool
	}{
		{"192.168.10.15", "IPv4Address", false},
		{"192.168.10.0/24", "IPv4Network", false},
		{"192.168.10.10-192.168.10.15", "IPv4Range", false},
		{"FE80::0202:B3FF:FE1E:8329", "IPv6Address", false},
		{"FE80::0202:B3FF:FE1E:8329/48", "IPv6Network", false},
		{"FE80::0202:B3FF:FE1E:8329-FE80::0202:B3FF:FE1E:8429", "IPv6Range", false},
		{"345.38.10.8", "", true},
	}

	for _, tc := range cases {
		kind, err := kindFromValue(tc.addr)
		if tc.err != (err != nil) {
			t.Fatalf("kindFromValue expected error: %t, got: %v", tc.err, err)
		}
		if kind != tc.kind {
			t.Errorf("expected: %s, got: %s", tc.kind, kind)
		}
	}
}
