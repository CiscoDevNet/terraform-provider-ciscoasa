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

import "fmt"

// ListAccessOutRules returns a collection of access control element objects.
func (s *accessService) ListAccessOutRules(iface string) (*ExtendedACEObjectCollection, error) {
	result := &ExtendedACEObjectCollection{}
	page := 0

	for {
		offset := page * s.pageLimit
		u := fmt.Sprintf("/api/access/out/%s/rules?limit=%d&offset=%d", iface, s.pageLimit, offset)

		req, err := s.newRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}

		e := &ExtendedACEObjectCollection{}
		_, err = s.do(req, e)
		if err != nil {
			return nil, err
		}

		result.RangeInfo = e.RangeInfo
		result.Items = append(result.Items, e.Items...)
		result.Kind = e.Kind
		result.SelfLink = e.SelfLink

		if e.RangeInfo.Offset+e.RangeInfo.Limit == e.RangeInfo.Total {
			break
		}
		page++
	}

	return result, nil
}

// CreateAccessOutRule creates an access control element.
func (s *accessService) CreateAccessOutRule(iface, src, srcService, dst, dstService string, active, permit bool) (string, error) {
	u := fmt.Sprintf("/api/access/out/%s/rules", iface)

	e := &ExtendedACEObject{
		Active: active,
		Permit: permit,
		Kind:   "object#ExtendedACE",
	}

	var err error
	if e.SrcAddress, err = s.Objects.objectFromAddress(src); err != nil {
		return "", err
	}
	if e.SrcService, err = s.Objects.objectFromService(srcService); err != nil {
		return "", err
	}
	if e.DstAddress, err = s.Objects.objectFromAddress(dst); err != nil {
		return "", err
	}
	if e.DstService, err = s.Objects.objectFromService(dstService); err != nil {
		return "", err
	}

	req, err := s.newRequest("POST", u, e)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)
	if err != nil {
		return "", err
	}

	return idFromResponse(resp)
}

// GetAccessOutRule retrieves an access control element.
func (s *accessService) GetAccessOutRule(iface string, ruleID string) (*ExtendedACEObject, error) {
	u := fmt.Sprintf("/api/access/out/%s/rules/%s", iface, ruleID)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	e := &ExtendedACEObject{}
	_, err = s.do(req, e)

	return e, err
}

// UpdateAccessOutRule updates an access control element.
func (s *accessService) UpdateAccessOutRule(iface, ruleID, src, srcService, dst, dstService string, active, permit bool) (string, error) {
	u := fmt.Sprintf("/api/access/out/%s/rules/%s", iface, ruleID)

	e := &ExtendedACEObject{
		Active: active,
		Permit: permit,
		Kind:   "object#ExtendedACE",
	}

	var err error
	if e.SrcAddress, err = s.Objects.objectFromAddress(src); err != nil {
		return "", err
	}
	if e.SrcService, err = s.Objects.objectFromService(src); err != nil {
		return "", err
	}
	if e.DstAddress, err = s.Objects.objectFromAddress(dst); err != nil {
		return "", err
	}
	if e.DstService, err = s.Objects.objectFromService(dstService); err != nil {
		return "", err
	}

	req, err := s.newRequest("POST", u, e)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)
	if err != nil {
		return "", err
	}

	return idFromResponse(resp)
}

// DeleteAccessOutRule deletes an access control element.
func (s *accessService) DeleteAccessOutRule(iface string, ruleID string) error {
	u := fmt.Sprintf("/api/access/out/%s/rules/%s", iface, ruleID)

	req, err := s.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}
