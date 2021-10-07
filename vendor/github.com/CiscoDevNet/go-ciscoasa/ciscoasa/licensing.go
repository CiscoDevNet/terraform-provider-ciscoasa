package ciscoasa

type licenseService struct {
	*Client
}

type SmartLicenseConfig struct {
	Kind             string `json:"kind,omitempty"`
	LicenseServerURL string `json:"licenseServerURL"`
	TransportURL     bool   `json:"transportURL"`
	PrivacyHostName  bool   `json:"privacyHostName"`
	PrivacyVersion   bool   `json:"privacyVersion"`
	Throughput       string `json:"throughput,omitempty"`
	FeatureTier      string `json:"featureTier,omitempty"`
}

// UpdateLicenseConfig updates a Smart License Config.
func (s *licenseService) UpdateLicenseConfig(lic *SmartLicenseConfig) error {
	u := "/api/licensing/smart/asav"

	req, err := s.newRequest("PUT", u, lic)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}

// GetLicense retrieves a Smart License Config.
func (s *licenseService) GetLicenseConfig() (*SmartLicenseConfig, error) {
	u := "/api/licensing/smart/asav"

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &SmartLicenseConfig{}
	_, err = s.do(req, r)

	return r, err
}

// RegisterLicense registers the Smart License.
func (c *licenseService) RegisterLicense(token string, force bool) error {
	u := "/api/licensing/smart/asav/register"

	r := struct {
		Kind    string `json:"kind"`
		IdToken string `json:"idToken"`
		Force   bool   `json:"force"`
	}{
		Kind:    "object#SmartLicenseRegId",
		IdToken: token,
		Force:   force,
	}

	req, err := c.newRequest("POST", u, r)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)

	return err
}

// DeregisterLicense removes the Smart License.
func (c *licenseService) DeregisterLicense() error {
	u := "/api/licensing/smart/asav/deregister"

	req, err := c.newRequest("POST", u, nil)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)

	return err
}

// RenewIdLicense renews the Smart License entitlement.
func (c *licenseService) RenewAuthLicense() error {
	u := "/api/licensing/smart/asav/renewauth"

	req, err := c.newRequest("POST", u, nil)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)

	return err
}

// RenewIdLicense renews the Smart License ID certificate.
func (c *licenseService) RenewIdLicense() error {
	u := "/api/licensing/smart/asav/renewid"

	req, err := c.newRequest("POST", u, nil)
	if err != nil {
		return err
	}

	_, err = c.do(req, nil)

	return err
}
