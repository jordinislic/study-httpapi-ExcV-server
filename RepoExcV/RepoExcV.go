package RepoExcV

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type ExcValue struct {
	CurlFrom  string
	CurlTo    string
	Value     float64
	CreatedOn string
}

type Value struct {
	Disclaimer string
	License    string
	Timestamp  int
	Base       string
	Rates      rate
}

type rate struct {
	EUR float64
}

type Repo struct {
	db *gorm.DB
}

func New(host string, port int, user string, password string, dbname string) Repo {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return Repo{
		db: DB,
	}
}

func (r Repo) ExchangeValue(jsonResp []byte) {

	var value Value
	err := json.Unmarshal(jsonResp, &value)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
	var v ExcValue
	v.CurlFrom = "USD"
	v.CurlTo = "EUR"
	v.Value = value.Rates.EUR
	ntime := int64(value.Timestamp)
	v.CreatedOn = time.Unix(ntime, 10000).Format("2006-01-02 15:04:05")
	fmt.Println("sono qui", v)
	r.db.Create(v)

}