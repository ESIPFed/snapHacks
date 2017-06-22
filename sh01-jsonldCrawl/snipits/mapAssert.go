package main

import "encoding/json"
import "fmt"

type MapContainer struct {
	M map[string]string
}

func main() {
	j := []byte(`{"a": "b", "c": "d"}`)

	var interfaceVal interface{}
	json.Unmarshal(j, &interfaceVal)

	fmt.Printf("interfaceVal's type: %T\n", interfaceVal)

	strings := interfaceVal.(map[string]interface{})

	valueA := strings["a"]

	fmt.Printf(valueA.(string) + "\n")

	var stringMap map[string]string
	json.Unmarshal(j, &stringMap)
	fmt.Printf("String map: %#v\n", stringMap)

	var mapContainer MapContainer
	json.Unmarshal(j, &mapContainer.M)
	fmt.Printf("mapContainer: %#v\n", mapContainer)

	var mapContainer2 MapContainer
	json.Unmarshal([]byte(`{"m": {"a": "b", "c": "d"}}`), &mapContainer2)
	fmt.Printf("mapContainer2: %#v\n", mapContainer2)

	var mapContainer3 MapContainer
	json.Unmarshal([]byte(`{"m": {"a": "b", "c": "d", "e": 3.2}}`), &mapContainer3)
	fmt.Printf("mapContainer3: %#v\n", mapContainer3)
}
