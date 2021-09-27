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

func TestCreateNetworkObjectGroup(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/objects/networkobjects/192.168.10.15", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, ``)
	})

	mux.HandleFunc("/api/objects/networkobjectgroups/192.168.10.15", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
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

	mux.HandleFunc("/api/objects/networkobjectgroups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testJSONBody(t, r, values{
			"kind": "object#NetworkObjGroup",
			"name": "test_objgrp",
			"members": []interface{}{
				map[string]interface{}{
					"kind":  "IPv4Address",
					"value": "192.168.10.15",
				},
				map[string]interface{}{
					"kind":     "objectRef#NetworkObj",
					"objectId": "test_obj",
				},
			},
		})

		fmt.Fprint(w, ``)
	})

	mux.HandleFunc("/api/objects/networkobjectgroups/test_objgrp", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
  			"kind": "object#NetworkObjGroup",
  			"selfLink": "https://localhost/api/objects/networkobjectgroups/test_objgrp",
  			"name": "test_objgrp",
  			"members": [{
				"kind": "IPv4Address",
				"value": "192.168.10.10"
			}],
			"objectId": "test_objgrp"
		}`)
	})

	_, err := client.Objects.CreateNetworkObjectGroup("test_objgrp", "", []string{"192.168.10.15", "test_obj"})
	if err != nil {
		t.Errorf("Failed to create NetworkObject test_obj: %s", err)
	}
}

func TestGetNetworkObjectGroup(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/objects/networkobjectgroups/test_objgrp", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
  			"kind": "object#NetworkObjGroup",
  			"selfLink": "https://localhost/api/objects/networkobjectgroups/test_objgrp",
  			"name": "test_objgrp",
  			"members": [{
				"kind": "IPv4Address",
				"value": "192.168.10.10"
			}],
			"objectId": "test_objgrp"
		}`)
	})

	o, err := client.Objects.GetNetworkObjectGroup("test_objgrp")
	if err != nil {
		t.Fatalf("GetNetworkObjectGroup got NetworkObject test_obj: %s", err)
	}

	if o.Members[0].Value != "192.168.10.10" {
		t.Errorf("Failed on GetNetworkObjectGroup test_obj, expected value 192.168.10.10, got %s", o.Members[0].Value)
	}
}
