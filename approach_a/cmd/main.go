package main

import (
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"insure"
)

// shortFormTime is a format.
// To understand why 2006-01-02, see https://stackoverflow.com/a/52966197/102699
const shortFormTime = "2006-01-02"

func main() {
	now := time.Now()

	/*
		var P Policy
		P.Number = "Policy1"
		P.line[0].ID = "Line1"
		P.line[1].ID = "Line2"
		P.transaction[0].ID = "Transaction1"
		P.line[0].coverages[0].Indicator = true
		P.line[0].Risks[0].ID = "Risk1"
		P.line[0].Risks[1].ID = "Risk1"
		P.line[0].Risks[0].coverages[0].Indicator = true
		P.line[0].Risks[1].coverages[0].Indicator = true
		P.line[0].loc[0].Address1 = "Addr1"
	*/
	policy := insure.NewPolicy("Policy1", &insure.PolicyOpts{
		EffectiveDate:  now,
		ExpirationDate: addYear(now),
	}).
		AddTransaction(insure.NewTransaction("Transaction1")).
		AddLine(insure.NewLine("Line1", &insure.LineOpts{
			TypeLOB:          insure.AutoLOBType,
			TermPremium:      decimal.NewFromInt(120),
			PriorTermPremium: decimal.NewFromInt(110),
		}).
			AddCoverage(insure.NewCoverage(true)).
			AddRisk(insure.NewRisk("Risk1", &insure.RiskOpts{
				Included: true,
			}).AddCoverage(insure.NewCoverage(true))).
			AddRisk(insure.NewRisk("Risk2", &insure.RiskOpts{
				Included: true,
			}).AddCoverage(insure.NewCoverage(true))).
			AddLocation(insure.NewLocation("Addr1")),
		).
		AddLine(insure.NewLine("Line2", &insure.LineOpts{}))

	fmt.Printf("%s\n", policy)

}

func addYear(t time.Time) time.Time {
	return t.AddDate(1, 0, 0)
}
