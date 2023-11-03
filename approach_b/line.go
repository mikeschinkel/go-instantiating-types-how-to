package insure

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type Lines []*Line

type Line struct {
	Id                  string
	TypeLOB             int
	coverages           Coverages
	TermPremium         decimal.Decimal
	PriorTermPremium    decimal.Decimal
	ChangePremium       decimal.Decimal
	WrittenPremium      decimal.Decimal
	PriorWrittenPremium decimal.Decimal
	risks               Risks
	locations           Locations
}

type LineOptions func(*Line)

func NewLine(Id string, opts ...LineOptions) *Line {
	l := &Line{
		Id:        Id,
		risks:     make([]*Risk, 0),
		coverages: make([]*Coverage, 0),
		locations: make([]*Location, 0),
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

func (LineOptions) SetTypeLOB(TypeLOB int) LineOptions {
	return func(l *Line) {
		l.TypeLOB = TypeLOB
	}
}

func (LineOptions) SetTermPremium(TermPremium decimal.Decimal) LineOptions {
	return func(l *Line) {
		l.TermPremium = TermPremium
	}
}

func (LineOptions) SetPriorTermPremium(PriorTermPremium decimal.Decimal) LineOptions {
	return func(l *Line) {
		l.PriorTermPremium = PriorTermPremium
	}
}

func (LineOptions) SetChangePremium(ChangePremium decimal.Decimal) LineOptions {
	return func(l *Line) {
		l.ChangePremium = ChangePremium
	}
}

func (LineOptions) SetWrittenPremium(WrittenPremium decimal.Decimal) LineOptions {
	return func(l *Line) {
		l.WrittenPremium = WrittenPremium
	}
}

func (LineOptions) SetPriorWrittenPremium(PriorWrittenPremium decimal.Decimal) LineOptions {
	return func(l *Line) {
		l.PriorWrittenPremium = PriorWrittenPremium
	}
}

func (LineOptions) AddCoverage(coverage *Coverage) LineOptions {
	return func(l *Line) {
		l.coverages = append(l.coverages, coverage)
	}
}

func (LineOptions) AddRisk(risk *Risk) LineOptions {
	return func(l *Line) {
		l.risks = append(l.risks, risk)
	}
}

func (LineOptions) AddLocation(location *Location) LineOptions {
	return func(l *Line) {
		l.locations = append(l.locations, location)
	}
}

func (l Line) StringWithTabs(tabs string) (s string) {
	return fmt.Sprintf("%sLine: %s\n"+
		"%sCoverages:\n%s"+
		"%sRisks:\n%s"+
		"%sLocations:\n%s",
		tabs, l.Id,
		tabs+tab, stringWithTabs(l.coverages, tabs+tab),
		tabs+tab, stringWithTabs(l.risks, tabs+tab),
		tabs+tab, stringWithTabs(l.locations, tabs+tab),
	)
}
