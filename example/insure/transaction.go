package insure

import (
	"fmt"
	"time"
)

type Transactions []*Transaction

type Transaction struct {
	Id             string
	TypeCode       string
	ImageType      string
	Status         string
	State          string
	EffectiveDate  time.Time
	CreatedDate    time.Time
	ExpirationDate time.Time
	IssuedDate     time.Time
}

func NewTransaction(id string) *Transaction {
	return &Transaction{Id: id}
}

func (tx Transaction) StringWithTabs(tabs string) string {
	return fmt.Sprintf("%sTransaction ID: %s", tabs, tx.Id)
}
