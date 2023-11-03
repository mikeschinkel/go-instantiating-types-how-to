package insure

type Insured struct {
	IsPrimaryApplicant     bool
	IsCoApplicant          bool
	IsAdditionalInterest   bool
	IsPerson               bool
	IsOrg                  bool
	IsPrimaryDriver        bool
	IsOccupantDriver       bool
	AdditionalInterestType string
	Id                     string
	AssociatedLocationID   string
}
