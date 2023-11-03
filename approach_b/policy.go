package insure

import (
	"fmt"
	"regexp"
	"time"
)

type Policy struct {
	number         string //GUID
	effectiveDate  time.Time
	expirationDate time.Time
	lines          Lines
	transactions   Transactions
}

type PolicyOptions func(*Policy)

func (PolicyOptions) SetEffectiveDate(d time.Time) PolicyOptions {
	return func(p *Policy) {
		p.effectiveDate = d
	}
}

func (PolicyOptions) SetExpirationDate(d time.Time) PolicyOptions {
	return func(p *Policy) {
		p.expirationDate = d
	}
}

func NewPolicy(number string, opts ...PolicyOptions) *Policy {
	p := &Policy{
		number:       number,
		lines:        make([]*Line, 0),
		transactions: make([]*Transaction, 0),
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (PolicyOptions) AddLine(line *Line) PolicyOptions {
	return func(p *Policy) {
		p.lines = append(p.lines, line)
	}
}

func (PolicyOptions) AddTransaction(tx *Transaction) PolicyOptions {
	return func(p *Policy) {
		p.transactions = append(p.transactions, tx)
	}
}

var reCompressNewLines = regexp.MustCompile(`\n+`)

func (p Policy) String() (s string) {
	s = fmt.Sprintf(
		"Policy:\n"+
			"%sNumber: %s\n"+
			"%sEffective Date: %s\n"+
			"%sExpiration Date: %s\n"+
			"%sLines:\n%s"+
			"%sTransactions:\n%s",
		tab, p.number,
		tab, p.effectiveDate.Format(shortFormTime),
		tab, p.expirationDate.Format(shortFormTime),
		tab, stringWithTabs(p.lines, tab+tab),
		tab, stringWithTabs(p.transactions, tab+tab),
	)
	return reCompressNewLines.ReplaceAllLiteralString(s, "\n")
}
