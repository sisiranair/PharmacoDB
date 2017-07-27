package api

import "github.com/gin-gonic/gin"

// HTTP Request Methods.
const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
	HEAD   string = "HEAD"
	OPTION string = "OPTION"
	PATCH  string = "PATCH"
)

// Route is a routing model.
type Route struct {
	Method   string
	Endpoint string
	Handler  gin.HandlerFunc
}

// Routes is a collection of Route.
type Routes []Route

var routes = Routes{
	Route{GET, "/cell_lines", IndexCell},
	Route{GET, "/cell_lines/:id", ShowCell},
	Route{GET, "/cell_lines/:id/drugs", CellDrugs},

	Route{GET, "/tissues", IndexTissue},
	Route{GET, "/tissues/:id", ShowTissue},
	Route{GET, "/tissues/:id/cell_lines", TissueCells},
	Route{GET, "/tissues/:id/drugs", TissueDrugs},

	Route{GET, "/drugs", IndexDrug},
	Route{GET, "/drugs/:id", ShowDrug},
	Route{GET, "/drugs/:id/cell_lines", DrugCells},
	Route{GET, "/drugs/:id/tissues", DrugTissues},

	Route{GET, "/datasets", IndexDataset},
	Route{GET, "/datasets/:id", ShowDataset},
	Route{GET, "/datasets/:id/cell_lines", DatasetCells},
	Route{GET, "/datasets/:id/tissues", DatasetTissues},
	Route{GET, "/datasets/:id/drugs", DatasetDrugs},

	Route{GET, "/experiments", IndexExperiment},
	Route{GET, "/experiments/:id", ShowExperiment},

	Route{GET, "/intersections", IndexIntersection},
	Route{GET, "/intersections/1/:cell_id/:drug_id", CellDrugIntersection},
	Route{GET, "/intersections/2/:cell_id/:dataset_id", CellDatasetIntersection},

	Route{GET, "/stats/tissue_cells", TissueCellStats},
	Route{GET, "/stats/dataset_drugs", DatasetDrugStats},
}
