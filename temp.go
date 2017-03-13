package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gin-gonic/gin.v1"
)

type Cell struct {
	Id        int            `json:"id"`
	Accession sql.NullString `json:"accession"`
	Name      string         `json:"name"`
	Tissue    sql.NullString `json:"tissue"`
}

var (
	cell   Cell
	result gin.H
)

func getCell(c *gin.Context) {
	id := c.Param("id")
	row := db.QueryRow("select cell_id, accession_id, cell_name, tissue_name from cells inner join tissues on cells.tissue_id = tissues.tissue_id where cells.cell_id = ?;", id)
	err = row.Scan(&cell.Id, &cell.Accession, &cell.Name, &cell.Tissue)
	if err != nil {
		row := db.QueryRow("select cell_id, accession_id, cell_name, tissue_name from cells inner join tissues on cells.tissue_id = tissues.tissue_id where cells.cell_id = ?;", id)
		err = row.Scan(&cell.Id, &cell.Accession, &cell.Name, &cell.Tissue)
		if err != nil {
			row := db.QueryRow("select cell_id, accession_id, cell_name, tissue_name from cells inner join tissues on cells.tissue_id = tissues.tissue_id where cells.cell_id = ?;", id)
			err = row.Scan(&cell.Id, &cell.Accession, &cell.Name, &cell.Tissue)
			if err != nil {
				result = gin.H{
					"category": "cell line",
					"count":    0,
					"data":     nil,
				}
			} else {
				result = gin.H{
					"category": "cell line",
					"count":    1,
					"data":     cell,
				}
			}
		} else {
			result = gin.H{
				"category": "cell line",
				"count":    1,
				"data":     cell,
			}
		}
	} else {
		result = gin.H{
			"category": "cell line",
			"count":    1,
			"data":     cell,
		}
	}
	c.JSON(http.StatusOK, result)
}

func getCells(c *gin.Context) {
	var (
		cell   Cell
		result gin.H
	)
	id := c.Param("id")
	row := db.QueryRow("select cell_id, accession_id, cell_name, tissue_name from cells inner join tissues on cells.tissue_id = tissues.tissue_id where cells.cell_id = ?;", id)
	err = row.Scan(&cell.Cell_Id, &cell.Accession_Id, &cell.Cell_Name)
	if err != nil {
		row := db.QueryRow("select cell_id, accession_id, cell_name, tissue_name from cells inner join tissues on cells.tissue_id = tissues.tissue_id where cells.cell_id = ?;", id)
		err = row.Scan(&cell.Cell_Id, &cell.Accession_Id, &cell.Cell_Name)
		if err != nil {
			row := db.QueryRow("select cell_id, accession_id, cell_name, tissue_name from cells inner join tissues on cells.tissue_id = tissues.tissue_id where cells.cell_id = ?;", id)
			err = row.Scan(&cell.Cell_Id, &cell.Accession_Id, &cell.Cell_Name)
			if err != nil {
				result = gin.H{
					"category": "cell line",
					"count":    0,
					"data":     nil,
				}
			} else {
				result = gin.H{
					"category": "cell line",
					"count":    1,
					"data":     cell,
				}
			}
		} else {
			result = gin.H{
				"category": "cell line",
				"count":    1,
				"data":     cell,
			}
		}
	} else {
		result = gin.H{
			"category": "cell line",
			"count":    1,
			"data":     cell,
		}
	}
	c.JSON(http.StatusOK, result)
}

func main() {
	dbname := os.Getenv("pharmacodb_api_dbname")
	passwd := os.Getenv("local_mysql_passwd")

	cred := "root:" + passwd + "@tcp(127.0.0.1:3306)/" + dbname

	// prepare database abstraction for later use
	db, err := sql.Open("mysql", cred)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// check that a network connection can be established and login
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	v1 := router.Group("v1")
	{
		// v1.GET("/cell_lines")
		v1.GET("/cell_lines/name/:name", CellByName)
		v1.GET("/cell_lines/id/:id", CellByID)
		// v1.GET("/cell_lines/accession/:accession", CellByACC)
		// v1.GET("/cell_lines/:name/synonyms")
		// v1.GET("/cell_lines/:name/drugs")
		// v1.GET("/cell_lines/:name/drugs_stat")
	}

	router.Run(":3000")
}

// GET cell line request handler (param: id)
func CellByID(c *gin.Context) {
	var (
		cell   Cell
		result gin.H
	)
	db := InitDb()
	defer db.Close()
	id := c.Param("id")
	row := db.QueryRow("select cell_id, accession_id, cell_name, tissue_name from cells inner join tissues on cells.tissue_id = tissues.tissue_id where cells.cell_id = ?;", id)
	err := row.Scan(&cell.ID, &cell.Accession, &cell.Name, &cell.Tissue)
	if err != nil {
		result = gin.H{
			"status":  http.StatusNotFound,
			"message": fmt.Sprintf("cell line with id - %s - not found in database", id),
		}
		c.JSON(http.StatusNotFound, result)
	} else {
		result = gin.H{
			"category": "cell line",
			"count":    1,
			"data":     cell,
		}
		c.JSON(http.StatusOK, result)
	}
}