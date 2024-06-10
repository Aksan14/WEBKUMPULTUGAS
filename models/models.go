package models

import (
	"tesgolang/config"
	"tesgolang/entities"
)

func See() []entities.User {
	query, err := config.DB.Query("SELECT * FROM  tugas1")

	if err != nil {
		panic(err)
	}

	defer query.Close()

	var category []entities.User

	for query.Next() {
		var categorys entities.User
		if err := query.Scan(&categorys.ID, &categorys.NAMA, &categorys.LINKTUGAS, &categorys.
			WAKTUKUMPUL); err != nil {
			panic(err)
		}

		category = append(category, categorys)
	}

	return category
}

func Create(category entities.User) bool {
	res, err := config.DB.Exec("INSERT INTO tugas1(NAMA, LINKTUGAS, WAKTUKUMPUL) VALUES (?,?,?)",
		category.NAMA,
		category.LINKTUGAS,
		category.WAKTUKUMPUL,
	)

	if err != nil {
		panic(err)
	}

	lastinsertid, err := res.LastInsertId()
	if err != nil {
		panic(err)
	}

	return lastinsertid > 0
}

func Delete(id int) error  {
	_, err := config.DB.Exec("DELETE FROM tugas1 WHERE ID = ? " ,id)

	return err
}