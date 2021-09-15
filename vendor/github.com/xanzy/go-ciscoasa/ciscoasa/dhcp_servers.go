package ciscoasa

import (
	"fmt"
)

// ListDhcpServers returns a collection of DHCP servers.
func (s *dhcpService) ListDhcpServers() (*DhcpServerCollection, error) {
	u := "/api/dhcp/servers"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &DhcpServerCollection{}
	_, err = s.do(req, r)

	return r, err
}

// UpdateDhcpServer updates a DHCP server.
func (s *dhcpService) UpdateDhcpServer(server *DhcpServer) (string, error) {
	u := fmt.Sprintf("/api/dhcp/servers/%s", server.ObjectId)

	req, err := s.newRequest("PUT", u, server)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)

	return idFromResponse(resp)
}

// GetDhcpServer retrieves a DHCP server.
func (s *dhcpService) GetDhcpServer(ipAddress string) (*DhcpServer, error) {
	u := fmt.Sprintf("/api/dhcp/servers/%s", ipAddress)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &DhcpServer{}
	_, err = s.do(req, r)

	return r, err
}
