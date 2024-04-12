package main

import (
	"bin-checker/structs"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
	"log"
)

type storage struct {
	db *goqu.Database
}

func newDb(dbConn *string) (storage, error) {
	postgres, err := sql.Open("postgres", *dbConn)
	if err != nil {
		return storage{}, err
	}

	return storage{goqu.New("postgres", postgres)}, nil
}

func (s *storage) getBin(binId string) (structs.SaveBinData, error) {
	var binData structs.SaveBinData

	found, err := s.db.
		Select(
			goqu.C("bin-id"),
			goqu.L("coalesce(brand, '')").As("brand"),
			goqu.L("coalesce(type, '')").As("type"),
			goqu.L("coalesce(category, '')").As("category"),
			goqu.L("coalesce(issuer, '')").As("issuer"),
			goqu.L("coalesce(alpha_2, '')").As("alpha_2"),
			goqu.L("coalesce(alpha_3, '')").As("alpha_3"),
			goqu.L("coalesce(country, '')").As("country")).
		From("bin_data").
		Where(goqu.Ex{"bin-id": binId}).
		ScanStruct(&binData)
	if err != nil {
		log.Fatal(err)
	}

	if !found {
		return structs.SaveBinData{}, err
	}

	return binData, nil
}

func (s *storage) saveBin(binData structs.SaveBinData) error {
	_, err := s.db.
		From("bin_data").
		Insert().Rows(binData).
		Executor().
		Exec()
	if err != nil {
		return err
	}

	return nil
}

func (s *storage) getAllBinsFromPostgres() ([]structs.SaveBinData, error) {
	var bins []structs.SaveBinData

	err := s.db.Select(
		goqu.C("bin-id"),
		goqu.L("coalesce(brand, '')").As("brand"),
		goqu.L("coalesce(type, '')").As("type"),
		goqu.L("coalesce(category, '')").As("category"),
		goqu.L("coalesce(issuer, '')").As("issuer"),
		goqu.L("coalesce(alpha_2, '')").As("alpha_2"),
		goqu.L("coalesce(alpha_3, '')").As("alpha_3"),
		goqu.L("coalesce(country, '')").As("country")).
		From("bin_data").ScanStructs(&bins)
	if err != nil {
		return nil, err
	}

	return bins, nil
}
