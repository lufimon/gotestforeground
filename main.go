package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	router := gin.Default()
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/training")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	group := router.Group("/test")
	{
		group.GET("/getAll", func(context *gin.Context) {
			getAll, err := db.Query("SELECT * FROM test")
			if err != nil {
				panic(err.Error())
			}
			defer getAll.Close()
			var results []string
			for getAll.Next() {
				var str string
				var device string
				errRow := getAll.Scan(&str, &device)
				if errRow != nil {
					panic(errRow.Error())
				}
				results = append(results, str+"/"+device)
			}
			if results != nil {
				context.JSON(http.StatusOK, results)
			} else {
				context.JSON(http.StatusInternalServerError, "ERROR")
			}
		})

		group.GET("/getAll/:device", func(context *gin.Context) {
			device := context.Param("device")
			getAll, err := db.Query("SELECT * FROM test WHERE device = ?", device)
			if err != nil {
				panic(err.Error())
			}
			defer getAll.Close()
			var results []string
			for getAll.Next() {
				var str string
				var device string
				errRow := getAll.Scan(&str, &device)
				if errRow != nil {
					panic(errRow.Error())
				}
				results = append(results, str+"/"+device)
			}
			if results != nil {
				context.JSON(http.StatusOK, results)
			} else {
				context.JSON(http.StatusInternalServerError, "ERROR")
			}
		})

		group.GET("/add/:param/:device", func(context *gin.Context) {
			param := context.Param("param")
			device := context.Param("device")
			insert, err := db.Query("INSERT INTO test VALUES(?, ?)", param, device)

			if err != nil {
				panic(err.Error())
			}

			defer insert.Close()
		})

		group.DELETE("/delete", func(context *gin.Context) {
			delete, err := db.Query("DELETE FROM test")
			if err != nil {
				panic(err.Error())
			}

			defer delete.Close()
		})
	}
	router.Run(":8888").Error()
}
