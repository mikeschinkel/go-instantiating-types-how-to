package insure

import (
	"fmt"
	"strings"
)

type Risks []*Risk
type Risk struct {
	id                   string
	guid                 string
	indicator            bool
	included             bool
	deleted              bool
	limit                int
	associatedLocationID string
	stringBuilderRisk    strings.Builder
	coverages            Coverages
}

type RiskOpts struct {
	ID                   string
	GUID                 string
	Indicator            bool
	Include              bool
	Deleted              bool
	Limit                int
	AssociatedLocationID string
	StringBuilderRisk    strings.Builder
}

func (r *Risk) AddCoverage(coverage *Coverage) *Risk {
	r.coverages = append(r.coverages, coverage)
	return r
}

func NewRisk(id string, opts *RiskOpts) *Risk {
	return &Risk{
		id:                   id,
		guid:                 opts.GUID,
		indicator:            opts.Indicator,
		included:             opts.Include,
		deleted:              opts.Deleted,
		limit:                opts.Limit,
		associatedLocationID: opts.AssociatedLocationID,
		stringBuilderRisk:    opts.StringBuilderRisk,
		coverages:            nil,
	}
}

func (r Risk) StringWithTabs(tabs string) (s string) {
	s = fmt.Sprintf("%sRisk ID: %s\n", tabs+tab, r.id)
	s += fmt.Sprintf("%sCoverages:\n%s", tabs+tab+tab, stringWithTabs(r.coverages, tabs+tab+tab))
	return s
}
