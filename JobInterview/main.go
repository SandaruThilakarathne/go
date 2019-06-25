package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	// "math"
	// "strings"
	// "reflect"
)
import "strings"
import "gopkg.in/couchbase/gocb.v1"


type Company struct {
	ResturantID string `json:"restaurant_id"`
	ResturantName string `json:"restaurant_name"`
	CurrencyCode string `json:"currency_code"`
	BranchID string `json:"branch_id"`
	BranchName string `json:"branch_name"`
	Key string `json:"key"`
}

type Geo struct {
	Lon string `json:"lon"`
	Lat string `json:"lat"`
	Key string `json:"keynew"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text,_ := reader.ReadString('\n')
	
	
	fmt.Println(text)

	file,_ := ioutil.ReadFile("example.json")
	fmt.Println("here")
	fmt.Print(file)
	
	var result map[string]interface{}
	errr2 := json.Unmarshal([]byte(file), &result)
	
	if errr2 != nil {
		fmt.Println("Error json unmashalling")
		fmt.Println(errr2.Error())
	}

	objects:= result["items"].([]interface{})

	cluster, _ := gocb.Connect("couchbase://localhost")
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: "Theesh",
		Password: "19951122V@n",
	})
	bucketCompany, _ := cluster.OpenBucket("company", "")
	bucketGeo, _ := cluster.OpenBucket("geo", "")
	// var company Company
	// var geo Geo
	for i := 0; i < len(objects); i++ {
		first := objects[i].(map[string]interface{})
		fmt.Println(first["restaurant_id"])
		// fmt.Println(first["restaurant_name"])
		// fmt.Println(first["currency_code"])
		// fmt.Println(first["branch_id"])
		// fmt.Println(first["branch_name"])
		resid := fmt.Sprintf("%v", first["restaurant_id"])
		branid := fmt.Sprintf("%v", first["branch_id"])
		idname := []string{resid, branid}
		joined := strings.Join(idname, "-")
		fmt.Println(joined)

		bucketCompany.Upsert(joined, Company{
			ResturantID: resid,
			ResturantName: fmt.Sprintf("%v", first["restaurant_name"]),
			CurrencyCode: fmt.Sprintf("%v", first["currency_code"]),
			BranchID: branid,
			BranchName: fmt.Sprintf("%v", first["branch_name"]),
			Key: joined,
		}, 0)

		bucketGeo.Upsert(joined, Geo{
			Lon: fmt.Sprintf("%v", first["lon"]),
			Lat: fmt.Sprintf("%v", first["lat"]),
			Key: joined,
		}, 0)
		
		// fmt.Println(first["restaurant_id"]+first["branch_id"])

	}

 }
