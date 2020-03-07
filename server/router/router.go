package router

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const port = ":8000"
const apiKey = "c3GiN7UKkiD6fo_UaPlEveSoz2alFdimqPHpbTI4cXSLBp7cgvrParSkyF3WNlWz9jlRCD4DlSMRuVtCLx5wOBiXRj-rlrDWgZYvZM5SlwtYgC70Ids1gZIbmwNfXnYx"

func fetchYelpBusinesses(params map[string]string) map[string]interface{} {
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(params).
		SetHeader("Accept", "application/json").
		SetAuthToken(apiKey).
		Get("https://api.yelp.com/v3/businesses/search")

	if err != nil {
		print("there was an error.")
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(resp.String()), &result)

	return result
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	yelpParams := map[string]string{
		"latitude":  "39.945800",
		"longitude": "-82.990910",
	}
	yelpTotalResponse := fetchYelpBusinesses(yelpParams)
	total := int(yelpTotalResponse["total"].(float64))
	offset := rand.Intn(total)
	yelpParams["offset"] = strconv.Itoa(offset)
	yelpParams["limit"] = "1"
	yelpBusinessResponse := fetchYelpBusinesses(yelpParams)
	json.NewEncoder(w).Encode(yelpBusinessResponse)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Result for %v\n", vars["id"])
}

/* Listen starts server on provided port */
func Listen() {
	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/random", randomHandler)
	r.HandleFunc("/result/{id}", resultHandler)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(port, loggedRouter))
}
