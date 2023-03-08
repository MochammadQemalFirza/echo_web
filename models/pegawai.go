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
	res.Message = "Success"
	res.Data = arrobj
	return res, nil
}

func StorePegawai(name string, age int, position string) (Response, error) {
	var res Response
	con := config.CreateCon()
	sqlstatement := "INSERT INTO pegawai (name, age, position) VALUES($1,$2,$3)"

	stmt, err := con.Prepare(sqlstatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(name, age, position)
	if err != nil {
		return res, err
	}
	// lastInsertedId, err := result.LastInsertId()
	// if err != nil {
	// 	return res, err
	// }
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = result
	return res, nil
}

func UpdatePegawai(id int, name string, age int, position string) (Response, error) {
	var res Response
	con := config.CreateCon()
	sqlstatement := "UPDATE pegawai SET name = $1, age = $2, position = $3 WHERE id =$4"
	stmt, err := con.Prepare(sqlstatement)
	if err != nil {
		return res, err
	}
	result, err := stmt.Exec(name, age, position, id)
	if err != nil {
		return res, err
	}
	// rowsAffected, err := result.RowsAffected()
	// if err != nil{
	// 	return res, err
	// }
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = result
	return res, nil
}
