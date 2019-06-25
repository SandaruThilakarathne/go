package geo

import "github.com/gin-gonic/gin"
import "gopkg.in/couchbase/gocb.v1"
import "fmt"
import "encoding/json"
// import "net/http"

// var bucket *gocb.Bucket

type Geo struct {
	Lon string `json:"lon"`
	Lat string `json:"lat"`
	Key string `json:"keynew"`
}

func test(c *gin.Context) {
	var n1q1Params []interface{}
	lon := c.Query("lon")
	lat := c.Query("lat")
	fmt.Print(lon)
	fmt.Print(lat)
	// var searchParamerter []interface{}
	n1q1Params = append(n1q1Params, lon)
	n1q1Params = append(n1q1Params, lat)
	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Theesh",
		Password: "19951122V@n",
	})
	bucketGeo, _ := cluster.OpenBucket("geo", "")

	query := gocb.NewN1qlQuery("SELECT geo.* FROM geo WHERE lon=$1 AND lat=$2")

	
	// fmt.Print(query)
	rows, err := bucketGeo.ExecuteN1qlQuery(query, n1q1Params)
	// var rows interface{}
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}
	var geo Geo
	rows.One(&geo)
	fmt.Print(rows)
	jsonBytes, _ := json.Marshal(geo)
	fmt.Println(string(jsonBytes))
	c.JSON(200, gin.H{
			"lon": lon,
			"lat": lat,
			"data": string(jsonBytes),
		})
}


func StartServer(port string) {
	router:= gin.Default()

	router.GET("/test", test)
	router.Run(port)
}