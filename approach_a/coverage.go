package insure

import (
	"fmt"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

type Coverages []*Coverage
type Coverage struct {
	indicator           bool
	included            bool
	deleted             bool
	modified            bool
	deductible          int
	level               string
	coverageType        string
	code                string
	formulae            string
	statCode            string
	currentLimit        string
	priorLimit          string
	baseRate            string
	ilf                 string
	addedDate           time.Time
	effectiveDate       time.Time
	expirationDate      time.Time
	deletedDate         time.Time
	termPremium         decimal.Decimal
	priorTermPremium    decimal.Decimal
	changePremium       decimal.Decimal
	writtenPremium      decimal.Decimal
	priorWrittenPremium decimal.Decimal
	termFactor          decimal.Decimal
	drf                 decimal.Decimal
	stringBuilder       strings.Builder
}

func NewCoverage(indicator bool) *Coverage {
	return &Coverage{indicator: indicator}
}

func (c Coverage) StringWithTabs(tabs string) (s string) {
	return fmt.Sprintf("%sCoverage: %t", tabs+tab, c.indicator)
}
