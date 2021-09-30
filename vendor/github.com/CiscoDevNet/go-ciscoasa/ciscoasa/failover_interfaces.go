package ciscoasa

import (
	"fmt"
)

// ListFailoverInterfaces returns a collection of Failover Interfaces.
func (s *failoverService) ListFailoverInterfaces() (*FailoverInterfacesCollection, error) {
	u := "/api/failover/interfaces"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &FailoverInterfacesCollection{}
	_, err = s.do(req, r)

	return r, err
}

// UpdateFailoverInterface updates a Failover Interface.
func (s *failoverService) UpdateFailoverInterface(iface *FailoverInterface) error {
	u := fmt.Sprintf("/api/failover/interfaces/%s", iface.ObjectId)

	req, err := s.newRequest("PUT", u, iface)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}

// GetFailoverInterface retrieves a Failover Interface.
func (s *failoverService) GetFailoverInterface(id string) (*FailoverInterface, error) {
	u := fmt.Sprintf("/api/failover/interfaces/%s", id)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &FailoverInterface{}
	_, err = s.do(req, r)

	return r, err
}
