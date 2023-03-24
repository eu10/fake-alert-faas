package function

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"

	"github.com/Jeffail/gabs"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	const endpoint = "https://api.bitcore.io/api/BTC/mainnet/block/tip"
	const filePath = "temp.json"
	var body []byte

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		body = Fetch(endpoint)
		_ = ioutil.WriteFile(filePath, body, 0644)
	} else {
		n := rand.Intn(100)
		if n < 60 {
			jsonFile, err := os.Open(filePath)
			if err != nil {
				fmt.Println(err)
			}
			defer jsonFile.Close()

			body, _ = ioutil.ReadAll(jsonFile)
		} else {
			body = Fetch(endpoint)
			_ = ioutil.WriteFile(filePath, body, 0644)
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(string(body))
	w.Header().Set("Content-Type", "application/json")
	jsonRes, _ := json.Marshal(body)
	w.Write(jsonRes)
}

func Fetch(endpoint string) []byte {
	var body []byte
	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return body
	}
	if resp.Body != nil {
		defer resp.Body.Close()

		body, _ = io.ReadAll(resp.Body)
	}
	return body
}

func Test2(data []byte) string {
	jsonParsed, err := gabs.ParseJSON(data)
	if err != nil {
		panic(err)
	}
	height := jsonParsed.Path("height").Data()
	fmt.Println(jsonParsed)
	fmt.Println(height)
	jsonParsed1, _ := gabs.ParseJSON([]byte(`{"height":123}`))
	jsonParsed.Delete("height")
	jsonParsed.Merge(jsonParsed1)
	fmt.Println(jsonParsed)
	return jsonParsed.String()
}

func Test1() {
	data := []byte(`{
		"employees":{
		   "protected":false,
		   "address":{
			  "street":"22 Saint-Lazare",
			  "postalCode":"75003",
			  "city":"Paris",
			  "countryCode":"FRA",
			  "country":"France"
		   },
		   "employee":[
			  {
				 "id":1,
				 "first_name":"Jeanette",
				 "last_name":"Penddreth"
			  },
			  {
				 "id":2,
				 "firstName":"Giavani",
				 "lastName":"Frediani"
			  }
		   ]
		}
	 }`)
	jsonParsed, err := gabs.ParseJSON(data)
	if err != nil {
		panic(err)
	}

	// Search JSON
	fmt.Println("Get value of Protected:\t", jsonParsed.Path("employees.protected").Data())
	fmt.Println("Get value of Country:\t", jsonParsed.Search("employees", "address", "country").Data())
	fmt.Println("ID of first employee:\t", jsonParsed.Path("employees.employee.0.id").String())
	fmt.Println("Check Country Exists:\t", jsonParsed.Exists("employees", "address", "countryCode"))
}
