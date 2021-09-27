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
	"fmt"
	"net/http"
	"testing"
)

func TestCreateNetworkObject(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/objects/networkobjects", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testJSONBody(t, r, values{
			"kind": "object#NetworkObj",
			"name": "test_obj",
			"host": map[string]interface{}{
				"kind":  "IPv4Address",
				"value": "192.168.10.15",
			},
		})

		fmt.Fprint(w, ``)
	})

	mux.HandleFunc("/api/objects/networkobjects/test_obj", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
 			"kind": "object#NetworkObj",
			"selfLink": "https://localhost/api/objects/networkobjects/test_obj",
			"name": "test_obj",
			"host": {
				"kind": "IPv4Network",
				"value": "192.168.10.0/24"
			},
			"objectId": "test_obj"
		}`)
	})

	_, err := client.Objects.CreateNetworkObject("test_obj", "", "192.168.10.15")
	if err != nil {
		t.Errorf("Failed to create NetworkObject test_obj: %s", err)
	}
}

func TestGetNetworkObject(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/objects/networkobjects/test_obj", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
 			"kind": "object#NetworkObj",
			"selfLink": "https://localhost/api/objects/networkobjects/test_obj",
			"name": "test_obj",
			"host": {
				"kind": "IPv4Network",
				"value": "192.168.10.0/24"
			},
			"objectId": "test_obj"
		}`)
	})

	o, err := client.Objects.GetNetworkObject("test_obj")
	if err != nil {
		t.Errorf("Failed to get NetworkObject test_obj: %s", err)
	} else {
		if o.Host.Value != "192.168.10.0/24" {
			t.Errorf("Failed to get NetworkObject test_obj, expected value 192.168.10.0/24, got %s", o.Host.Value)
		}
	}
}
