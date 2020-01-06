package models

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jinzhu/gorm"
	//"github.com/lib/pq"
)

type Test struct {
	Adata int
	Bdata string
}

// Users
type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"not null"`
	Name     string
	Role     string `gorm:"size:255"`
	KpLog    string `gorm:"size:2048"`
	Tests    []byte
	Stickers *MyJson
	Mytodo   *Test
	//Test 		 pq.StringArray `gorm:"type:varchar(100)[]"`
}

type MyJson struct {
	Data string
}

// Scan interface
func (ls *MyJson) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	t := MyJson{}
	if e := json.Unmarshal(value.([]byte), &t); e != nil {
		return e
	}
	*ls = t
	return nil
}

// Value interface
func (ls *MyJson) Value() (driver.Value, error) {
	if ls == nil {
		return nil, nil
	}
	b, e := json.Marshal(*ls)
	return b, e
}

// Scan interface
func (ls *Test) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	t := Test{}
	if e := json.Unmarshal(value.([]byte), &t); e != nil {
		return e
	}
	*ls = t
	return nil
}

// Value interface
func (ls *Test) Value() (driver.Value, error) {
	if ls == nil {
		return nil, nil
	}
	b, e := json.Marshal(*ls)
	return b, e
}
