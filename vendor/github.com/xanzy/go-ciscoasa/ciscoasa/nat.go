package ciscoasa

import (
	"encoding/json"
	"fmt"
)

type natService struct {
	*Client
}

type NatCollection struct {
	Items     []*Nat    `json:"items"`
	Kind      string    `json:"kind"`
	RangeInfo RangeInfo `json:"rangeInfo"`
	SelfLink  string    `json:"selfLink"`
}

type TranslatedOriginalObj struct {
	Kind     string `json:"kind"`
	ObjectId string `json:"objectId,omitempty"`
	Value    string `json:"value,omitempty"`
}

type TranslatedOriginalInterface struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type Nat struct {
	Active                      bool                         `json:"active"`
	BlockAllocation             bool                         `json:"blockAllocation,omitempty"`
	Description                 string                       `json:"description,omitempty"`
	Extended                    bool                         `json:"extended,omitempty"`
	Flat                        bool                         `json:"flat,omitempty"`
	IncludeReserve              bool                         `json:"includeReserve,omitempty"`
	IsDNS                       bool                         `json:"isDNS,omitempty"`
	IsInterfacePAT              bool                         `json:"isInterfacePAT,omitempty"`
	IsNetToNet                  bool                         `json:"isNetToNet,omitempty"`
	IsNoProxyArp                bool                         `json:"isNoProxyArp,omitempty"`
	IsPatPool                   bool                         `json:"isPatPool,omitempty"`
	IsRoundRobin                bool                         `json:"isRoundRobin,omitempty"`
	IsRouteLookup               bool                         `json:"isRouteLookup,omitempty"`
	IsUnidirectional            bool                         `json:"isUnidirectional,omitempty"`
	Kind                        string                       `json:"kind,omitempty"`
	Mode                        string                       `json:"mode"`
	ObjectID                    string                       `json:"objectId,omitempty"`
	OriginalDestination         *TranslatedOriginalObj       `json:"originalDestination,omitempty"`
	OriginalInterface           *TranslatedOriginalInterface `json:"originalInterface,omitempty"`
	OriginalService             *TranslatedOriginalObj       `json:"originalService,omitempty"`
	OriginalSource              *TranslatedOriginalObj       `json:"originalSource,omitempty"`
	Position                    int                          `json:"position,omitempty"`
	SelfLink                    string                       `json:"selfLink,omitempty"`
	TranslatedDestination       *TranslatedOriginalObj       `json:"translatedDestination,omitempty"`
	TranslatedInterface         *TranslatedOriginalInterface `json:"translatedInterface,omitempty"`
	TranslatedService           *TranslatedOriginalObj       `json:"translatedService,omitempty"`
	TranslatedSource            *TranslatedOriginalObj       `json:"translatedSource,omitempty"`
	TranslatedSourcePatPool     *TranslatedOriginalObj       `json:"translatedSourcePatPool,omitempty"`
	UseDestinationInterfaceIPv6 bool                         `json:"useDestinationInterfaceIPv6,omitempty"`
	UseSourceInterfaceIPv6      bool                         `json:"useSourceInterfaceIPv6,omitempty"`
}

func (iface *TranslatedOriginalInterface) UnmarshalJSON(b []byte) error {
	type alias TranslatedOriginalInterface
	if err := json.Unmarshal(b, (*alias)(iface)); err != nil {
		iface = nil
	}
	return nil
}

func (translated *TranslatedOriginalObj) UnmarshalJSON(b []byte) error {
	type alias TranslatedOriginalObj
	if err := json.Unmarshal(b, (*alias)(translated)); err != nil {
		translated = nil
	}
	return nil
}

// ListNats returns a collection of NATs.
func (s *natService) ListNats(section string) (*NatCollection, error) {
	u := fmt.Sprintf("/api/nat/%s", section)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &NatCollection{}
	_, err = s.do(req, r)

	return r, err
}

// CreateNat creates a NAT.
func (s *natService) CreateNat(section string, nat *Nat) (string, error) {
	u := fmt.Sprintf("/api/nat/%s", section)

	req, err := s.newRequest("POST", u, nat)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)
	if err != nil {
		return "", err
	}

	return idFromResponse(resp)
}

// GetNat retrieves a NAT.
func (s *natService) GetNat(section, id string) (*Nat, error) {
	u := fmt.Sprintf("/api/nat/%s/%s", section, id)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &Nat{}
	_, err = s.do(req, r)

	return r, err
}

// UpdateNat updates a NAT.
func (s *natService) UpdateNat(section, id string, nat *Nat) (string, error) {
	u := fmt.Sprintf("/api/nat/%s/%s", section, id)

	req, err := s.newRequest("PUT", u, nat)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)
	if err != nil {
		return "", err
	}

	return idFromResponse(resp)
}

// DeleteNat deletes a NAT.
func (s *natService) DeleteNat(section, id string) error {
	u := fmt.Sprintf("/api/nat/%s/%s", section, id)

	req, err := s.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}
