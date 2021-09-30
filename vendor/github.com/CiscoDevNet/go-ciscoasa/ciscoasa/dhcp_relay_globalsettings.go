package ciscoasa

// UpdateDhcpRelayGlobalsettings updates a DHCP Relay Global Settings.
func (s *dhcpService) UpdateDhcpRelayGlobalsettings(gs *DhcpRelayGS) error {
	u := "/api/dhcp/relay/servers/globalsettings"

	req, err := s.newRequest("PUT", u, gs)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}

// GetDhcpRelayGlobalsettings retrieves a DHCP Relay Global Settings.
func (s *dhcpService) GetDhcpRelayGlobalsettings() (*DhcpRelayGS, error) {
	u := "/api/dhcp/relay/servers/globalsettings"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &DhcpRelayGS{}
	_, err = s.do(req, r)

	return r, err
}
