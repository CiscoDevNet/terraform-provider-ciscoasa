package ciscoasa

import (
	"fmt"
)

// ListDhcpRelayLocals returns a collection of DHCP relay interface servers.
func (s *dhcpService) ListDhcpRelayLocals() (*DhcpRelayLocalCollection, error) {
	u := "/api/dhcp/relay/servers/local"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &DhcpRelayLocalCollection{}
	_, err = s.do(req, r)

	return r, err
}

// CreateDhcpRelayLocal creates a DHCP relay interface server.
func (s *dhcpService) CreateDhcpRelayLocal(iface string, servers []string) error {
	u := "/api/dhcp/relay/servers/local"

	r := &DhcpRelayLocal{
		Interface: iface,
		Servers:   servers,
		Kind:      "object#DHCPRelayInterfaceServer",
	}

	req, err := s.newRequest("POST", u, r)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}

// GetDhcpRelayLocal retrieves a DHCP relay interface server.
func (s *dhcpService) GetDhcpRelayLocal(iface string) (*DhcpRelayLocal, error) {
	u := fmt.Sprintf("/api/dhcp/relay/servers/local/%s", iface)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &DhcpRelayLocal{}
	_, err = s.do(req, r)

	return r, err
}

// UpdateDhcpRelayLocal updates a DHCP relay interface server.
func (s *dhcpService) UpdateDhcpRelayLocal(iface string, servers []string) error {
	u := fmt.Sprintf("/api/dhcp/relay/servers/local/%s", iface)

	r := &DhcpRelayLocal{
		Interface: iface,
		Servers:   servers,
		Kind:      "object#DHCPRelayInterfaceServer",
	}

	req, err := s.newRequest("PUT", u, r)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}

// DeleteDhcpRelayLocal deletes a DHCP relay interface server.
func (s *dhcpService) DeleteDhcpRelayLocal(iface string) error {
	u := fmt.Sprintf("/api/dhcp/relay/servers/local/%s", iface)

	req, err := s.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}
