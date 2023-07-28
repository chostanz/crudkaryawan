package main

import (
	"fmt"
	_ "log"
	"net/http"
	"strconv"

	sqlx "github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
)

// User
// gaboleh declare di main
type Karyawan struct {
	Id     int    `json:"id" db:"id"`
	Nama   string `json:"name" db:"name"`
	No_hp  string `json:"phone" db:"phone"`
	Alamat string `json:"address" db:"address"`
}

type Respons struct {
	Message string
	Status  bool
	//Data    []Karyawan
}

func main() {

	//membuat koneksi database, mengambil objek di db untuk menjalankan berbagai method query
	db, err := sqlx.Connect("postgres", "user=postgres password=12345 dbname= db_users sslmode=disable")

	if err = db.Ping(); err != nil {
		fmt.Println(err)
	}

	//mengisi struct Respons, disimpan di respon
	respon := Respons{
		Message: "Sukses menjalankan query",
		Status:  true,
	}

	e := echo.New()

	e.Use(middleware.CORS()) //untuk allow web server

	// get method karyawan -> menampilkan data karyawan
	e.GET("/users", func(c echo.Context) error {

		//menjalankan query di simpan di var rows
		//parameter dan argumen tu sama, ada di dalam function
		rows, _ := db.Queryx("select * from users")

		var users []Karyawan //membuat slice

		for rows.Next() {
			//membuat var bru dari struct Karyawan yg nilai awalnya kosong
			place := Karyawan{}
			rows.StructScan(&place)
			users = append(users, place) //menambah elemen baru menggunakan append ke users
		}
		return c.JSON(http.StatusOK, users)
	})

	// post method karyawan insert data ->
	e.POST("/users", func(c echo.Context) error {
		reqBody := Karyawan{}

		c.Bind(&reqBody)

		// query insert pakek insert into tapi value pakek dari reqbody
		db.NamedExec("insert into users(name, phone, address) values (:name, :phone, :address)", reqBody)

		return c.JSON(http.StatusOK, respon)
	})

	e.PUT("/users/update/:id", func(c echo.Context) error {
		reqBody := Karyawan{}

		c.Bind(&reqBody)
		//menggunakan parameter untuk menghapus data dgn nilai dari parameter :id, bertipe string
		id := c.Param("id")
		parsedID, err := strconv.Atoi(id) // Konversi id dari string ke int
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		}
		//mengembalikan nilai id ke reqBody.Id (sturct ReqBody Id)
		reqBody.Id = parsedID

		// query update value pakek dari reqbody
		db.NamedExec("update users SET name= :name, phone= :phone, address= :address WHERE id= :id", reqBody)

		return c.JSON(http.StatusOK, respon)
	})

	e.DELETE("/users/delete/:id", func(c echo.Context) error {
		// reqBody := Karyawan{}
		// c.Bind(&reqBody)
		// db.NamedExec("DELETE from users WHERE id= :id", reqBody)
		// return c.JSON(http.StatusOK, respon)

		id := c.Param("id")
		parsedID, err := strconv.Atoi(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
		}

		// Lakukan DELETE pada database menggunakan ID yang bertipe int
		// hapus data dengan id tertentu
		if err := deleteDataByID(parsedID); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete data"})
		}

		return c.NoContent(http.StatusNoContent) // Mengembalikan status 204 No Content
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func deleteDataByID(id int) error {
	db, err := sqlx.Connect("postgres", "user=postgres password=12345 dbname= db_users sslmode=disable")

	db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
