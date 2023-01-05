package go_xorm

import (
	"errors"
	"time"
)

type Contract struct {
	ID        int64 `xorm:"pk autoincr"`
	UID       int64
	OrgOrUser string
	Cluster   string
	Product   string
	Features  []string `xorm:"jsonb"`

	Emails ContractEmail `xorm:"jsonb"`

	StartUnix    int64
	EndUnix      int64
	DocumentPath string

	Revoked bool
}

type ContractEmail struct {
	ContactPerson    []string `json:"contactPerson"`
	SigningAuthority []string `json:"signingAuthority"`
	LicenseUsers     []string `json:"licenseUsers"`
	FinanceTeam      []string `json:"financeTeam"`
}

func ListActiveContractsForCluster(uid int64, cluster string) ([]Contract, error) {
	cds := make([]Contract, 0)
	if err := en.Desc("end_unix").Where("`end_unix` > ?", time.Now().UTC().Unix()).Find(&cds, Contract{
		UID:     uid,
		Cluster: cluster,
	}); err != nil {
		return nil, err
	}

	return cds, nil
}

func ListContracts(uid int64) ([]Contract, error) {
	cds := make([]Contract, 0)

	if err := en.Find(&cds, Contract{UID: uid}); err != nil {
		return nil, err
	}

	return cds, nil
}

func GetContract(id int64, uid int64) (*Contract, error) {
	cd := Contract{ID: id, UID: uid}

	has, err := en.Get(&cd)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, errors.New("not found")
	}

	return &cd, nil
}

func AddContract(cd Contract) (*Contract, error) {
	if _, err := en.Insert(&cd); err != nil {
		return nil, err
	}

	return &cd, nil
}
