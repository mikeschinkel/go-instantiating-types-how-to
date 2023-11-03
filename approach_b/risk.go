package insure

import (
	"fmt"
	"strings"
)

type Risks []*Risk
type Risk struct {
	Id                   string
	Indicator            bool
	Included             bool
	Deleted              bool
	GUID                 string
	AssociatedLocationID string
	StringBuilderRisk    strings.Builder
	coverages            []*Coverage
}

type RiskOptions func(*Risk)

func (ro RiskOptions) SetIndicator(indicator bool) RiskOptions {
	return func(r *Risk) {
		r.Indicator = indicator
	}
}

func (RiskOptions) SetIncluded(included bool) RiskOptions {
	return func(r *Risk) {
		r.Included = included
	}
}

func (RiskOptions) SetDeleted(deleted bool) RiskOptions {
	return func(r *Risk) {
		r.Deleted = deleted
	}
}

func (RiskOptions) SetGuid(guid string) RiskOptions {
	return func(r *Risk) {
		r.GUID = guid
	}
}

func (RiskOptions) SetAssociatedLocationId(associatedLocationId string) RiskOptions {
	return func(r *Risk) {
		r.AssociatedLocationID = associatedLocationId
	}
}

func (RiskOptions) SetStringBuilderRisk(stringBuilderRisk strings.Builder) RiskOptions {
	return func(r *Risk) {
		r.StringBuilderRisk = stringBuilderRisk
	}
}

func (RiskOptions) AddCoverage(coverage *Coverage) RiskOptions {
	return func(r *Risk) {
		r.coverages = append(r.coverages, coverage)
	}
}

func NewRisk(id string, opts ...RiskOptions) *Risk {
	r := &Risk{Id: id}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (r Risk) StringWithTabs(tabs string) (s string) {
	s = fmt.Sprintf("%sRisk ID: %s\n", tabs+tab, r.Id)
	s += fmt.Sprintf("%sCoverages:\n%s", tabs+tab+tab, stringWithTabs(r.coverages, tabs+tab+tab))
	return s
}
