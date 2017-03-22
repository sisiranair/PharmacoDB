package main

import (
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

// GetDatasets handles GET requests for /datasets endpoint.
func GetDatasets(c *gin.Context) {
	queryStr := "select dataset_id, dataset_name from datasets;"
	desc := "List of all datasets in pharmacodb"
	getDataTypes(c, desc, queryStr)
}

// GetDatasetStats handles GET requests for /datasets/stats
func GetDatasetStats(c *gin.Context) {
	var (
		cstat DatasetStat
		tstat DatasetStat
		dstat DatasetStat
		estat DatasetStat

		cstats []DatasetStat
		tstats []DatasetStat
		dstats []DatasetStat
		estats []DatasetStat
	)

	db, err := initDB()
	defer db.Close()
	if err != nil {
		handleError(c, nil, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	rows, err := db.Query("select dataset_id, cell_lines, tissues, drugs, experiments from dataset_statistics;")
	defer rows.Close()
	if err != nil {
		handleError(c, err, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	for rows.Next() {
		err = rows.Scan(&cstat.Dataset, &cstat.Count, &tstat.Count, &dstat.Count, &estat.Count)
		if err != nil {
			handleError(c, err, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		tstat.Dataset = cstat.Dataset
		dstat.Dataset = cstat.Dataset
		estat.Dataset = cstat.Dataset
		cstats = append(cstats, cstat)
		tstats = append(tstats, tstat)
		dstats = append(dstats, dstat)
		estats = append(estats, estat)
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"description": "Number of items tested in each dataset per datatype, as well as number of experiments carried out in each dataset",
		"data": gin.H{
			"cell_lines":  cstats,
			"tissues":     tstats,
			"drugs":       dstats,
			"experiments": estats,
		},
	})
}