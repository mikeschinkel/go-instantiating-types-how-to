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
	var po insure.PolicyOptions
	var lo insure.LineOptions
	var ro insure.RiskOptions

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
	policy := insure.NewPolicy("Policy1",
		po.SetEffectiveDate(now),
		po.SetExpirationDate(addYear(now)),
		po.AddTransaction(insure.NewTransaction("Transaction1")),
		po.AddLine(
			insure.NewLine("Line1",
				lo.SetTypeLOB(insure.AutoLOBType),
				lo.SetTermPremium(decimal.NewFromInt(120)),
				lo.SetPriorTermPremium(decimal.NewFromInt(110)),
				lo.AddCoverage(insure.NewCoverage(true)),
				lo.AddRisk(insure.NewRisk("Risk1",
					ro.AddCoverage(insure.NewCoverage(true)),
				)),
				lo.AddRisk(insure.NewRisk("Risk2",
					ro.AddCoverage(insure.NewCoverage(true)),
				)),
				lo.AddLocation(insure.NewLocation("Addr1")),
			),
		),
		po.AddLine(insure.NewLine("Line2")),
	)

	fmt.Printf("%s\n", policy)

}

func addYear(t time.Time) time.Time {
	return t.AddDate(1, 0, 0)
}
