package serivces

import "github.com/gin-gonic/gin"
import "gopkg.in/couchbase/gocb.v1"
import "fmt"
import "encoding/json"
// import "net/http"

type Company struct {
	ResturantID string `json:"restaurant_id"`
	ResturantName string `json:"restaurant_name"`
	CurrencyCode string `json:"currency_code"`
	BranchID string `json:"branch_id"`
	BranchName string `json:"branch_name"`
	Key string `json:"key"`
}

type CompayList struct {
	Companies []string
}

func test2(c *gin.Context) {
	var n1q1Params []interface{}
	resturant_branch_id := c.Query("id")

	n1q1Params = append(n1q1Params, resturant_branch_id)

	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Theesh",
		Password: "19951122V@n",
	})
	bucketCompany, _ := cluster.OpenBucket("company", "")

	query := gocb.NewN1qlQuery("SELECT company.* FROM company WHERE META().id=$1")

	rows, err := bucketCompany.ExecuteN1qlQuery(query, n1q1Params)
	// var rows interface{}
	var company Company
	rows.One(&company)
	fmt.Print(rows)
	jsonBytes, _ := json.Marshal(company)
	fmt.Println(string(jsonBytes))
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}
	c.JSON(200, gin.H{
			"id": string(jsonBytes),
		})
}

func getCompanyIdList(c *gin.Context) {

	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Theesh",
		Password: "19951122V@n",
	})
	bucketCompany, _ := cluster.OpenBucket("company", "")

	query := gocb.NewN1qlQuery("SELECT company.* FROM company WHERE META().id=$1")

	var comp CompayList
	var outpuData []string
	c.BindJSON(&comp)
	// fmt.Print(comp.Companies)
	for _, s := range comp.Companies {
		var n1q1Params []interface{}
		n1q1Params = append(n1q1Params, s)
		rows, _ := bucketCompany.ExecuteN1qlQuery(query, n1q1Params)
		// var rows interface{}
		var company Company
		rows.One(&company)
		fmt.Print(rows)
		jsonBytes, _ := json.Marshal(company)
		fmt.Println(string(jsonBytes))
		outpuData = append(outpuData, string(jsonBytes))
		
	}
	c.JSON(200, gin.H{
			"id": outpuData,
		})
}

func StartServer(port string) {
	router:= gin.Default()
	router.GET("/test2", test2)
	router.POST("/takeList", getCompanyIdList)
	router.Run(port)
}