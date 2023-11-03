package insure

import (
	"fmt"
	"regexp"
	"time"
)

type Policy struct {
	number         string
	effectiveDate  time.Time
	expirationDate time.Time
	lines          Lines
	transactions   Transactions
}

type PolicyOpts struct {
	EffectiveDate  time.Time
	ExpirationDate time.Time
}

func NewPolicy(number string, opts *PolicyOpts) *Policy {
	return &Policy{
		number:         number,
		effectiveDate:  opts.EffectiveDate,
		expirationDate: opts.ExpirationDate,
		lines:          make([]*Line, 0),
		transactions:   make([]*Transaction, 0),
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

func (p *Policy) AddLine(line *Line) *Policy {
	p.lines = append(p.lines, line)
	return p
}

func (p *Policy) AddTransaction(tx *Transaction) *Policy {
	p.transactions = append(p.transactions, tx)
	return p
}
