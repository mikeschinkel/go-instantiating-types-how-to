package insure

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Lines []*Line

type Line struct {
	id                  string
	typeLOB             int
	coverages           Coverages
	termPremium         decimal.Decimal
	priorTermPremium    decimal.Decimal
	changePremium       decimal.Decimal
	writtenPremium      decimal.Decimal
	priorWrittenPremium decimal.Decimal
	risks               Risks
	locations           Locations
}

type LineOpts struct {
	TypeLOB             int
	TermPremium         decimal.Decimal
	PriorTermPremium    decimal.Decimal
	ChangePremium       decimal.Decimal
	WrittenPremium      decimal.Decimal
	PriorWrittenPremium decimal.Decimal
}

func NewLine(id string, opts *LineOpts) *Line {
	return &Line{
		id:                  id,
		typeLOB:             opts.TypeLOB,
		termPremium:         opts.TermPremium,
		priorTermPremium:    opts.PriorTermPremium,
		changePremium:       opts.ChangePremium,
		writtenPremium:      opts.WrittenPremium,
		priorWrittenPremium: opts.PriorWrittenPremium,
		risks:               make([]*Risk, 0),
		coverages:           make([]*Coverage, 0),
		locations:           make([]*Location, 0),
	}
}
func (l *Line) AddCoverage(coverage *Coverage) *Line {
	l.coverages = append(l.coverages, coverage)
	return l
}

func (l *Line) AddRisk(risk *Risk) *Line {
	l.risks = append(l.risks, risk)
	return l
}

func (l *Line) Risks() Risks {
	return l.risks
}

func (l *Line) AddLocation(location *Location) *Line {
	l.locations = append(l.locations, location)
	return l
}

func (l Line) StringWithTabs(tabs string) (s string) {
	return fmt.Sprintf("%sLine: %s\n"+
		"%sCoverages:\n%s"+
		"%sRisks:\n%s"+
		"%sLocations:\n%s",
		tabs, l.id,
		tabs+tab, stringWithTabs(l.coverages, tabs+tab),
		tabs+tab, stringWithTabs(l.risks, tabs+tab),
		tabs+tab, stringWithTabs(l.locations, tabs+tab),
	)
}
