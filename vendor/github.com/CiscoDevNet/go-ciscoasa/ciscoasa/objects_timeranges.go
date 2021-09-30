package ciscoasa

import (
	"fmt"
)

// TimeRangeCollection represents a collection of time ranges.
type TimeRangeCollection struct {
	RangeInfo RangeInfo    `json:"rangeInfo"`
	Items     []*TimeRange `json:"items"`
	Kind      string       `json:"kind"`
	SelfLink  string       `json:"selfLink"`
}

type TRValue struct {
	Start    string      `json:"start"`
	End      string      `json:"end"`
	Periodic []*Periodic `json:"periodic"`
}

// TimeRange represents a time range.
type TimeRange struct {
	Name     string   `json:"name,omitempty"`
	Value    *TRValue `json:"value,omitempty"`
	Kind     string   `json:"kind"`
	ObjectID string   `json:"objectId,omitempty"`
	SelfLink string   `json:"selfLink,omitempty"`
}

// ListTimeRange returns a collection of time ranges.
func (s *objectsService) ListTimeRange() (*TimeRangeCollection, error) {
	result := &TimeRangeCollection{}
	page := 0

	for {
		offset := page * s.pageLimit
		u := fmt.Sprintf("/api/objects/timeranges?limit=%d&offset=%d", s.pageLimit, offset)

		req, err := s.newRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}

		n := &TimeRangeCollection{}
		_, err = s.do(req, n)
		if err != nil {
			return nil, err
		}

		result.RangeInfo = n.RangeInfo
		result.Items = append(result.Items, n.Items...)
		result.Kind = n.Kind
		result.SelfLink = n.SelfLink

		if n.RangeInfo.Offset+n.RangeInfo.Limit == n.RangeInfo.Total {
			break
		}
		page++
	}
	return result, nil
}

// CreateTimeRange creates a new time range.
func (s *objectsService) CreateTimeRange(name, start, end string, periodic []*Periodic) (*TimeRange, error) {
	u := "/api/objects/timeranges"

	// o, err := s.objectFromService(service)
	// if err != nil {
	// 	return nil, err
	// }
	value := &TRValue{
		Start:    start,
		End:      end,
		Periodic: periodic,
	}

	n := &TimeRange{
		Name:  name,
		Value: value,
		Kind:  "object#TimeRange",
	}

	req, err := s.newRequest("POST", u, n)
	if err != nil {
		return nil, err
	}

	_, err = s.do(req, nil)
	if err != nil {
		return nil, err
	}

	return s.GetTimeRange(name)
}

// GetTimeRange retrieves a time range.
func (s *objectsService) GetTimeRange(name string) (*TimeRange, error) {
	u := fmt.Sprintf("/api/objects/timeranges/%s", name)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	n := &TimeRange{}
	_, err = s.do(req, n)

	return n, err
}

// UpdateTimeRange updates a time range.
func (s *objectsService) UpdateTimeRange(name, start, end string, periodic []*Periodic) (*TimeRange, error) {
	u := fmt.Sprintf("/api/objects/timeranges/%s", name)

	value := &TRValue{
		Start:    start,
		End:      end,
		Periodic: periodic,
	}

	n := &TimeRange{
		Name:  name,
		Value: value,
		Kind:  "object#TimeRange",
	}

	req, err := s.newRequest("PUT", u, n)
	if err != nil {
		return nil, err
	}

	_, err = s.do(req, nil)
	if err != nil {
		return nil, err
	}

	return s.GetTimeRange(name)
}

// DeleteTimeRange deletes a time range.
func (s *objectsService) DeleteTimeRange(name string) error {
	u := fmt.Sprintf("/api/objects/timeranges/%s", name)

	req, err := s.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}
