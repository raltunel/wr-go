package tables

import (
	"database/sql"
	"encoding/json"
	"log"
)

type CreatedContractTable struct{}

func (tbl CreatedContractTable) GetID(r CreatedContract) string {
	return r.ID
}

func (tbl CreatedContractTable) GetTime(r CreatedContract) int {
	return r.Time
}

func (tbl CreatedContractTable) GetBlock(r CreatedContract) int {
	return r.Block
}

type CreatedContract struct {
	ID      string `db:"id"`
	Network string `db:"network"`
	Tx      string `db:"tx"`
	Block   int    `db:"block"`
	Time    int    `db:"time"`
	User    string `db:"user"`
	Token   string `db:"token"`
}

type CreatedContractSubGraph struct {
	ID              string `json:"id"`
	TransactionHash string `json:"transactionHash"`
	Block           string `json:"block"`
	Time            string `json:"time"`
	User            string `json:"user"`
	Token           string `json:"token"`
}

type CreatedContractSubGrapData struct {
	CreatedContracts []CreatedContractSubGraph `json:"createdContracts"`
}

type CreatedContractSubGraphResp struct {
	Data CreatedContractSubGrapData `json:"data"`
}

func (tbl CreatedContractTable) ConvertSubGraphRow(r CreatedContractSubGraph, network string) CreatedContract {
	return CreatedContract{
		ID:      network + r.ID,
		Network: network,
		Tx:      r.TransactionHash,
		Block:   parseInt(r.Block),
		Time:    parseInt(r.Time),
		User:    translateUser(r.User),
		Token:   r.Token,
	}
}

func (tbl CreatedContractTable) SqlTableName() string { return "createdContracts" }

func (tbl CreatedContract) ReadSqlRow(rows *sql.Rows) CreatedContract {
	var createdContract CreatedContract
	err := rows.Scan(
		&createdContract.ID,
		&createdContract.Network,
		&createdContract.Tx,
		&createdContract.Block,
		&createdContract.Time,
		&createdContract.User,
		&createdContract.Token,
	)
	if err != nil {
		log.Fatal(err)
	}
	return createdContract
}

func (tbl CreatedContractTable) ParseSubGraphResp(body []byte) ([]CreatedContractSubGraph, error) {
	var parsed CreatedContractSubGraphResp

	err := json.Unmarshal(body, &parsed)
	if err != nil {
		return nil, err
	}

	ret := make([]CreatedContractSubGraph, 0)
	for _, entry := range parsed.Data.CreatedContracts {
		ret = append(ret, entry)
	}
	return ret, nil
}
