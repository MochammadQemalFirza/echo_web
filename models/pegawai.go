package models

import (
	"net/http"

	"github.com/MochammadQemalFirza/echo_web/config"
)

type Pegawai struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Position string `json:"position"`
}

func FetchPegawai() (Response, error) {
	var peg Pegawai
	var arrobj []Pegawai
	var res Response
	con := config.CreateCon()

	sqlstatement := "SELECT * FROM pegawai"

	rows, err := con.Query(sqlstatement)
	defer rows.Close()
	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&peg.Id, &peg.Name, &peg.Age, &peg.Position)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, peg)
	}
	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrobj
	return res, nil
}
