package model

import (
	"github.com/ethereum/go-ethereum/common"
	"time"
)

// search
type SearchIndex struct {
	Details             *Details `json:"details"`
	Total               int      `json:"total"`
	NameList            []string `json:"name_list"`
	Registration_Status int      `json:"registration_status"`
	IsAddress           int      `json:"is_address"`
}
type Details struct {
	Domain          string         `json:"domain"`
	Registrant      common.Address `json:"registrant"`
	Expiration_date time.Time      `json:"expiration_date"`
	Resolver        common.Address `json:"resolver"`
	Addresses       common.Address `json:"addresses"`
	Content         string         `json:"content"`
	Snapshot        string         `json:"snapshot"`
	Url             string         `json:"url"`
	Avatar          string         `json:"avatar"`
	Twitter         string         `json:"twitter"`
	Github          string         `json:"github"`
	Email           string         `json:"email"`
	Description     string         `json:"description"`
	Notice          string         `json:"notice"`
	Keywords        string         `json:"keywords"`
	Discord         string         `json:"discord"`
	Reddit          string         `json:"reddit"`
	Telegram        string         `json:"telegram"`
	Delegate        string         `json:"delegate"`
}
