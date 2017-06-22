package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/deiu/rdf2go"
	"github.com/kazarena/json-gold/ld"
)

// DataSetStruct is struct to hold information about schema.org/DataSet
type DataSetStruct struct {
	Description string
	ID          string
	Type        string
	URL         string
}

// simple JSON-LD doc for some early testing   Will be removed later as we use simpleServer
const bodyTest = `{
    "@context": "http://schema.org",
    "@type": "DataCatalog",
    "@id": "http://opencoredata.org/catalogs",
    "url": "http://opencoredata.org/catalogs",
    "description": "Can I use this approach to reference this catalog from type WebSite",
    "dataset": [{
            "@type": "Dataset",
            "description": "An example dataset 1",
            "url": "http://opencoredata.org/id/rdf/geolink1.ttl"
        },
        {
            "@type": "Dataset",
            "description": "An example dataset 2",
            "url": "http://opencoredata.org/id/rdf/cruises.ttl"
        }
    ]
}
`

// A simple crawler to go through a given web site (single domain) and starting at
// a JSON-LD document, walk through the tree of documents.  Driven by either hydra or
// by JSON-LD framing.  We wil try both
func main() {
	fmt.Println("Simple crawler")

	// Take a seed domain
	// read int he JSON-LD  (validate it, then frame it against a given frame...  place results into a struct)
	// Load the URLs discovered to crawl to a boltdb system (where I can check if they have been already crawled)
	// Store the JSON-LD to a graph as triples.  (also store the original JSON-LD to a bolt table.)
	// In the end we have the triples of the site...

	body := getDoc("http://opencoredata.org")
	body = []byte(bodyTest) // replace with test block above....   getDoc is []byte, frameForDataCatalog is string  (review)
	// TODO  Do we need to load this up in the URL KV with a status visited?
	// TODO  Do we need to store the docs we get associated with a URL as well?  simple to do

	frameresult := frameForDataCatalog(string(body))

	// for each in struct, pull out the URL's and pass to
	for k, v := range frameresult {
		log.Printf("Item %d with URL: %v   \n", k, v.URL)
		// TODO Register the URL in a KV with status set to unvisited
		registerURL(v.URL)
	}
}

// getDoc simply takes a URL and return the contents of the response body to a byte array
func getDoc(urlstring string) []byte {

	u, err := url.Parse(urlstring)
	if err != nil {
		log.Println(err)
	}

	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Accept", "application/json") // oddly the content-type is ignored for the accept header...
	req.Header.Set("Cache-Control", "no-cache")
	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer res.Body.Close()

	// secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(res.Body)
	return body
}

// frameForDataCatalog take string and JSON-LD and uses a frame call to extract
// only type DataSet.  This is then marshalled to a struct...
func frameForDataCatalog(jsonld string) []DataSetStruct {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	frame := map[string]interface{}{
		"@context": "http://schema.org",
		"@type":    "Dataset",
	}

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
	}

	framedDoc, err := proc.Frame(myInterface, frame, options) // do I need the options set in order to avoid the large context that seems to be generated?
	if err != nil {
		log.Println("Error when trying to frame document", err)
	}

	graph := framedDoc["@graph"]
	// ld.PrintDocument("JSON-LD graph section", graph)  // debug print....
	jsonm, err := json.MarshalIndent(graph, "", " ")
	if err != nil {
		log.Println("Error trying to marshal data", err)
	}

	dss := make([]DataSetStruct, 0)
	err = json.Unmarshal(jsonm, &dss)
	if err != nil {
		log.Println("Error trying to unmarshal data to struct", err)
	}
	// log.Printf("%v\n", dss)
	return dss
}

// registerURL take a URL and places it into the bolt KV store.
// While doing so it first ensures that the URL has not already been placed into the KV store
// regardless of whether the URL has been marked as read.
// These URLs come from a framing call onto the JSON-LD for a particular @type
func registerURL(urlstring string) {

	// check to see if we have been there before..  if not, load and set status unvisited
	// If it's in the KV already ignore..  this is just a register system

}

// getURLToVisit just looks into the KV store and looks for a URL to visit...
func getURLToVisit() string {

}

// processJSONLD takes the JSONLD document (as a byte array) and processes it to ensure
// it is valid.  It then
func graphJSONLD(jsonld string) {

	baseURI := "https://earthcube.org/cdf/"

	// Create a new graph
	g := rdf2go.NewGraph(baseURI)
	g.Parse(strings.NewReader(jsonld), "application/ld+json")

	// if err != nil {
	// 	// deal with err
	// }
}

// jsonLDToRDF take a JSON-LD string and convert it to n-triples and returns it.
func jsonLDToRDF(jsonld string) string {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")
	options.Format = "application/nquads"

	var myInterface interface{}
	err := json.Unmarshal([]byte(jsonld), &myInterface)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to interface:", err)
	}

	triples, err := proc.ToRDF(myInterface, options)
	if err != nil {
		log.Println("Error when transforming JSON-LD document to RDF:", err)
		return err.Error()
	}

	return triples.(string)
}
