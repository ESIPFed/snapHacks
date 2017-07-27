package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/boltdb/bolt"
	"github.com/deiu/rdf2go"
	"github.com/kazarena/json-gold/ld"
)

// DataSetStruct is struct to hold information about schema.org/DataSet
type DataSetStruct struct {
	Description string
	ID          string
	Type        string
	URL         string `json:"schema:url"`
}

// simple JSON-LD doc for some early testing   Will be removed later as we use simpleServer
const bodyTestOLD = `{
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

const bodyTest = ` {
 "@context": {
  "@vocab": "http://schema.org/",
  "re3data": "http://example.org/re3data/0.1/"
 },
 "@id": "http://opencoredata.org/catalogs/geolink",
 "@type": "DataCatalog",
 "dataset": [
  {
   "@type": "Dataset",
   "description": "Collection of cruise data (leg level) for IODP collected for GeoLink",
   "url": "http://opencoredata.org/catalog/geolink/dataset/JRSO_cruises_gl"
  },
  {
   "@type": "Dataset",
   "description": "Collection of Science Party Deployment Information collected for GeoLink",
   "url": "http://opencoredata.org/catalog/geolink/dataset/JRSO_deployments_gl"
  },
  {
   "@type": "Dataset",
   "description": "Collection of cruise data (hole level) for IODP collected for GeoLink",
   "url": "http://opencoredata.org/catalog/geolink/dataset/JRSO_holes_gl"
  }
 ],
 "description": "A catalog of RDF graphs from Open Core Data for GeoLink that align to the GeoLink base ontology",
 "url": "http://opencoredata.org/catalogs/geolink"
}
`

// A simple crawler to go through a given web site (single domain) and starting at
// a JSON-LD document, walk through the tree of documents leveraging JSON-LD framing
func main() {
	fmt.Println("Simple crawler")

	// setup bolt
	SetupBolt()

	// Loop and load the whitelist URLs into the DB to start with
	registerURL("http://opencoredata.org/catalog/geolink")

	URL, count := getURLToVisit() // just grab our initial set of URLs to visit, don't worry about a URL string returned

	for count > 0 {
		URL, count = getURLToVisit()
		fmt.Printf("URL to work with: %s with count %d\n", URL, count)

		if count == 0 {
			break
		}

		body := extractJSON(URL)

		// fmt.Println(string(body))

		// body = []byte(bodyTest) // TEST REPLACE BODY  replace with test block above....   getDoc is []byte, frameForDataCatalog is string  (review)
		frameresult := frameForDataCatalog(string(body))

		for k, v := range frameresult {
			log.Printf("Item %d with URL: %v   \n", k, v.URL)
			// TODO Register the URL in a KV with status set to unvisited
			registerURL(v.URL)
		}

		// TODO   do further processing with the body, like index and to RDF
		// for each in struct, pull out the URL's and pass to

		visitedURL(URL) // TODO make the URL visited
	}

	// Spit out all the content of the KV store just to evaluate what was done with it...
	showAllKV()
}

func showAllKV() {

	db, err := bolt.Open("walker.db", 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("URLBucket"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
		}

		return nil
	})

	db.Close() // explicitly close
}

// getDoc DEPRECATED
// simply takes a URL and return the contents of the response body to a byte array
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

func extractJSON(urlstring string) []byte {
	resp, err := soup.Get(urlstring)
	if err != nil {
		log.Print(err)
	}
	doc := soup.HTMLParse(resp)

	//     <script type="application/ld+json">
	jsonld := doc.Find("script", "type", "application/ld+json").Text()

	return []byte(jsonld)

	// links := doc.Find("div", "id", "comicLinks").FindAll("a")
	// for _, link := range links {
	// 	fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	// }
}

// frameForDataCatalog take string and JSON-LD and uses a frame call to extract
// only type DataSet.  This is then marshalled to a struct...
func frameForDataCatalog(jsonld string) []DataSetStruct {
	proc := ld.NewJsonLdProcessor()
	options := ld.NewJsonLdOptions("")

	frame := map[string]interface{}{
		"@context": "http://schema.org/",
		"@type":    "Dataset",
	}

	// frame := map[string]interface{}{
	// 	"@context": {
	// 		"@vocab":  "http://schema.org/",
	// 		"re3data": "http://example.org/re3data/0.1/",
	// 	},
	// 	"@type": "Dataset",
	// }

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

	log.Printf("This is the dss:  %v\n", dss)
	return dss
}

func visitedURL(urlstring string) {

	// open in write mode
	db, err := bolt.Open("walker.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// What should the key be?  Just a simple UID?
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("URLBucket"))
		err := b.Put([]byte(urlstring), []byte("visited"))
		return err
	})

	db.Close() // explicitly close...

	// look for key and set value to "visited"

}

// registerURL take a URL and places it into the bolt KV store.
// While doing so it first ensures that the URL has not already been placed into the KV store
// regardless of whether the URL has been marked as read.
// These URLs come from a framing call onto the JSON-LD for a particular @type
func registerURL(urlstring string) {
	// open in write mode
	db, err := bolt.Open("walker.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// TODO, check if the key is already in the db
	// db.

	// What should the key be?  Just a simple UID?
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("URLBucket"))
		err := b.Put([]byte(urlstring), []byte("unvisited"))
		return err
	})

	db.Close() // explicitly close...

	// check to see if we have been there before..  if not, load and set status unvisited
	// If it's in the KV already ignore..  this is just a register system

}

// getURLToVisit just looks into the KV store and looks for a URL to visit...
func getURLToVisit() (string, int) {

	//  open in read only mode so at not to block and get the first URL we find that
	// is of value "unvisited"
	db, err := bolt.Open("walker.db", 0600, &bolt.Options{ReadOnly: true})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var uvsite []byte
	count := 0

	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("URLBucket"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			if strings.Compare(string(v), "unvisited") == 0 {
				uvsite = k
				count = count + 1
			}
			// fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})

	return string(uvsite), count

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

func SetupBolt() {

	db, err := bolt.Open("walker.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// You can also create a bucket only if it doesn't exist by using the Tx.CreateBucketIfNotExists()
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("URLBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		log.Printf("Bucket created %v", b.FillPercent)
		return nil
	})
}
