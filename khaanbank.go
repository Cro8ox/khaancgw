package khaancgw

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"git.fibo.cloud/itrip/payments/khaancgw/utils"
	"github.com/Cro8ox/khaancgw/structs"
)

type khaanCGW struct {
	username string
	password string
	endpoint string
}

type KhaanCGW interface {
	Statement(params StatementParams) ([]StatementTransaction, error)
}

func New(username, password, endpoint string) KhaanCGW {
	return &khaanCGW{
		username: username,
		password: password,
		endpoint: endpoint,
	}
}

type (
	StatementParams struct {
		StartDate time.Time `json:"start_date" binding:"required"`
		EndDate   time.Time `json:"end_date" binding:"required"`
		AccountNo string    `json:"account_no" binding:"required"`
		Currency  string    `json:"currency" binding:"required"`
		Record    string    `json:"record"`
	}

	StatementTransaction struct {
		ID     string    `json:"id"`
		IsIn   bool      `json:"is_in"`
		Date   time.Time `json:"date"`
		Rate   float64   `json:"rate"`
		Amount float64   `json:"amount"`
		Info   string    `json:"info"`
	}
)

func (k *khaanCGW) Statement(params StatementParams) ([]StatementTransaction, error) {
	path := fmt.Sprintf(Statement.Url, params.AccountNo, params.StartDate.Format("20060102"), params.EndDate.Format("20060102"))
	fmt.Printf("Statement Path: %s\n", path)
	if len(params.Record) > 0 {
		path = path + fmt.Sprintf("&record=%s", params.Record)
	}
	statementByteRes, err := k.HttpRequest(Statement.Method, path, nil)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Statement Response: %s\n", string(statementByteRes))
	var statementResponse structs.StatementResponse
	if err := json.Unmarshal(statementByteRes, &statementResponse); err != nil {
		return nil, err
	}
	var result []StatementTransaction
	for _, trans := range statementResponse.Transactions {
		date, err := utils.ParseDate(trans.TranDate, trans.Time)
		if err != nil {
			return nil, err
		}
		result = append(result, StatementTransaction{
			ID:     strconv.FormatUint(trans.Record, 10),
			IsIn:   utils.IfAssigment(trans.Amount > 0, true, false).(bool),
			Date:   date,
			Rate:   0,
			Amount: trans.Amount,
			Info:   trans.Description,
		})
	}
	return result, nil
}
