package ciscoasa

var u = "/api/failover/setup"

// UpdateFailoverSetup updates a Failover Setup.
func (s *failoverService) UpdateFailoverSetup(setup *FailoverSetup) error {

	req, err := s.newRequest("PUT", u, setup)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}

// GetFailoverSetup retrieves a Failover Setup.
func (s *failoverService) GetFailoverSetup() (*FailoverSetup, error) {

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &FailoverSetup{}
	_, err = s.do(req, r)

	return r, err
}
