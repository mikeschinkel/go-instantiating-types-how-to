package insure

import (
	"fmt"
	"strings"
)

type Risks []*Risk
type Risk struct {
	id                   string
	indicator            bool
	included             bool
	deleted              bool
	guid                 string
	associatedLocationID string
	stringBuilderRisk    strings.Builder
	coverages            Coverages
}

type RiskOpts struct {
	ID                   string
	Indicator            bool
	Included             bool
	Deleted              bool
	GUID                 string
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
		indicator:            opts.Indicator,
		included:             opts.Included,
		deleted:              opts.Deleted,
		guid:                 opts.GUID,
		associatedLocationID: opts.AssociatedLocationID,
		stringBuilderRisk:    opts.StringBuilderRisk,
	}
}

func (r Risk) StringWithTabs(tabs string) (s string) {
	s = fmt.Sprintf("%sRisk ID: %s\n", tabs+tab, r.id)
	s += fmt.Sprintf("%sCoverages:\n%s", tabs+tab+tab, stringWithTabs(r.coverages, tabs+tab+tab))
	return s
}
