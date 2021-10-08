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
	"errors"
	"fmt"
	"net"
)

type routingService struct {
	*Client
}

// RoutingObjectCollection represents a collection of routing objects.
type RoutingObjectCollection struct {
	RangeInfo RangeInfo        `json:"rangeInfo"`
	Items     []*RoutingObject `json:"items"`
	Kind      string           `json:"kind"`
	SelfLink  string           `json:"selfLink"`
}

// RoutingObject represents a routing object.
type RoutingObject struct {
	DistanceMetric int            `json:"distanceMetric"`
	Gateway        *AddressObject `json:"gateway"`
	Interface      struct {
		Name     string `json:"name"`
		Kind     string `json:"kind"`
		ObjectID string `json:"objectId,omitempty"`
	} `json:"interface"`
	Network  *AddressObject `json:"network"`
	Tracked  bool           `json:"tracked,omitempty"`
	Tunneled bool           `json:"tunneled,omitempty"`
	Kind     string         `json:"kind"`
	ObjectID string         `json:"objectId,omitempty"`
	SelfLink string         `json:"selfLink,omitempty"`
}

// ListStaticRoutes returns a collection of static routes.
func (s *routingService) ListStaticRoutes() (*RoutingObjectCollection, error) {
	u := "/api/routing/static"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &RoutingObjectCollection{}
	_, err = s.do(req, r)

	return r, err
}

// CreateStaticRoute creates a static route.
func (s *routingService) CreateStaticRoute(iface, network, gateway string, metric int, tracked, tunneled bool) (string, error) {
	u := "/api/routing/static"

	if metric < 1 || metric > 255 {
		return "", errors.New("metric must be between 0 and 255")
	}

	// Confirm that the gateway specifies an IP address before trying to
	// create an AddressObject from the given value.
	addr := net.ParseIP(gateway)
	if addr == nil {
		return "", fmt.Errorf("gateway must be an IP address")
	}

	gatewayObject, err := s.Objects.objectFromAddress(gateway)
	if err != nil {
		return "", err
	}

	networkObject, err := s.Objects.objectFromAddress(network)
	if err != nil {
		return "", err
	}

	r := &RoutingObject{
		DistanceMetric: metric,
		Gateway:        gatewayObject,
		Network:        networkObject,
		Tracked:        tracked,
		Tunneled:       tunneled,
		Kind:           fmt.Sprintf("object#%sRoute", networkObject.Kind[:4]),
	}

	r.Interface.Name = iface
	r.Interface.Kind = "objectRef#Interface"

	req, err := s.newRequest("POST", u, r)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)
	if err != nil {
		return "", err
	}

	return idFromResponse(resp)
}

// GetNetworkObject retrieves a network object.
func (s *routingService) GetStaticRoute(routeID string) (*RoutingObject, error) {
	u := fmt.Sprintf("/api/routing/static/%s", routeID)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &RoutingObject{}
	_, err = s.do(req, r)

	return r, err
}

// UpdateStaticRoute updates a static route.
func (s *routingService) UpdateStaticRoute(routeID, iface, network, gateway string, metric int, tracked, tunneled bool) (string, error) {
	u := fmt.Sprintf("/api/routing/static/%s", routeID)

	if metric < 1 || metric > 255 {
		return "", errors.New("metric must be between 0 and 255")
	}

	// Confirm that the gateway specifies an IP address before trying to
	// create an AddressObject from the given value.
	addr := net.ParseIP(gateway)
	if addr == nil {
		return "", fmt.Errorf("gateway must be an IP address")
	}

	gatewayObject, err := s.Objects.objectFromAddress(gateway)
	if err != nil {
		return "", err
	}

	networkObject, err := s.Objects.objectFromAddress(network)
	if err != nil {
		return "", err
	}

	r := &RoutingObject{
		DistanceMetric: metric,
		Gateway:        gatewayObject,
		Network:        networkObject,
		Tracked:        tracked,
		Tunneled:       tunneled,
		Kind:           fmt.Sprintf("object#%sRoute", networkObject.Kind[:4]),
	}

	r.Interface.Name = iface
	r.Interface.Kind = "objectRef#Interface"

	req, err := s.newRequest("PUT", u, r)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)
	if err != nil {
		return "", err
	}

	return idFromResponse(resp)
}

// DeleteStaticRoute deletes a static route.
func (s *routingService) DeleteStaticRoute(routeID string) error {
	u := fmt.Sprintf("/api/routing/static/%s", routeID)

	req, err := s.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}
