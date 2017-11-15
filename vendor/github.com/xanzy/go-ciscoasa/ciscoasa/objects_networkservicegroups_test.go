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

func TestCreateNetworkServiceGroup(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/objects/networkservicegroups", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testJSONBody(t, r, values{
			"kind": "object#NetworkServiceGroup",
			"name": "28ec2bb3",
			"members": []interface{}{
				map[string]interface{}{
					"kind":  "TcpUdpService",
					"value": "tcp/2121",
				},
				map[string]interface{}{
					"kind":  "TcpUdpService",
					"value": "udp/123",
				},
				map[string]interface{}{
					"kind":  "TcpUdpService",
					"value": "tcp/https",
				},
			},
		})

		fmt.Fprint(w, ``)
	})

	mux.HandleFunc("/api/objects/networkservicegroups/28ec2bb3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
  			"kind": "object#NetworkServiceGroup",
  			"selfLink": "https://localhost/api/objects/networkservicegroups/28ec2bb3",
  			"name": "28ec2bb3",
  			"members": [
  			{
  				"kind": "TcpUdpService",
  				"value": "tcp/2121"
  			},
  			{
  				"kind": "TcpUdpService",
  				"value": "udp/123"
  			},
  			{
  				"kind": "TcpUdpService",
  				"value": "tcp/https"
  			}],
			"objectId": "28ec2bb3"
		}`)
	})

	_, err := client.Objects.CreateNetworkServiceGroup("28ec2bb3", "", []string{"tcp/2121", "udp/123", "tcp/https"})
	if err != nil {
		t.Errorf("Failed to create NetworkServiceGroup 28ec2bb3: %s", err)
	}
}

func TestGetNetworkServiceGroup(t *testing.T) {
	mux, server, client := setup()
	defer teardown(server)

	mux.HandleFunc("/api/objects/networkservicegroups/28ec2bb3", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{
  			"kind": "object#NetworkServiceGroup",
  			"selfLink": "https://localhost/api/objects/networkservicegroups/28ec2bb3",
  			"name": "28ec2bb3",
  			"members": [
  			{
  				"kind": "TcpUdpService",
  				"value": "tcp/2121"
  			},
  			{
  				"kind": "TcpUdpService",
  				"value": "udp/123"
  			},
  			{
  				"kind": "TcpUdpService",
  				"value": "tcp/https"
  			}],
			"objectId": "28ec2bb3"
		}`)
	})

	o, err := client.Objects.GetNetworkServiceGroup("28ec2bb3")
	if err != nil {
		t.Fatalf("GetNetworkServiceGroup got NetworkServiceGroup 28ec2bb3: %s", err)
	}

	if o.Members[0].Value != "tcp/2121" {
		t.Errorf("Failed on GetNetworkServiceGroup 28ec2bb3, expected value tcp/2121, got %s", o.Members[0].Value)
	}
	if o.Members[1].Value != "udp/123" {
		t.Errorf("Failed on GetNetworkServiceGroup 28ec2bb3, expected value udp/123, got %s", o.Members[1].Value)
	}
	if o.Members[2].Value != "tcp/https" {
		t.Errorf("Failed on GetNetworkServiceGroup 28ec2bb3, expected value tcp/https, got %s", o.Members[2].Value)
	}
}
