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
	ID          string
	Type        string
	Description string
	URL         string
}

// simple JSON-LD doc for some early testing   Will be removed later as we use simpleServer
const jsonldtest = `{
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

	frameresult := frameForDataCatalog(jsonldtest)
	fmt.Println(frameresult)
}

func frameForDataCatalog(jsonld string) []string {
	var results []string

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

	fmt.Println(graph) // If I don't do @graph the @context part is HUGE..  why is that?  Is the code going out on the net and getting something?

	// test getting a single result into a struct..  later make an []DataSet

	// ref https://stackoverflow.com/questions/38185916/convert-interface-to-map-in-golang
	// ds := DataSetStruct{}
	v, ok := graph.([]map[string]string) // this looks like it should assert..  based on the print statement above..
	if !ok {
		// Can't assert, handle error.
		fmt.Println("Assert failed")
	}
	for _, s := range v {
		fmt.Printf("Value: %v\n", s)
	}

	// fake results for now to satisfy the func return type
	results = append(results, "test string")
	return results
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

// registerURL take a URL and places it into the bolt KV store.
// While doing so it first ensures that the URL has not already been placed into the KV store
// regardless of whether the URL has been marked as read.
// These URLs come from a framing call onto the JSON-LD for a particular @type
func registerURL(urlstring string) {

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
