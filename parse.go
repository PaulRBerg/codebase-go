// Parsing arbitrary JSON using interfaces in Go
// Demonstrates how to parse JSON with abritrary key names
// See https://blog.golang.org/json-and-go for more info on generic JSON parsing

package main

import (
	"encoding/json"
	"fmt"
)

var jsonBytes = []byte(`
{
	"key1":{
		"Item1": "Value1",
		"Item2": 1},
	"key2":{
		"Item1": "Value2",
		"Item2": 2},
	"key3":{
		"Item1": "Value3",
		"Item2": 3}
}`)

// Item struct; we want to create these from the JSON above
type Item struct {
	Item1 string
	Item2 int
}

// Implement the String interface for pretty printing Items
func (i Item) String() string {
	return fmt.Sprintf("Item: %s, %d", i.Item1, i.Item2)
}

func mainparse() {

	// Unmarshal using a generic interface
	var f interface{}
	err := json.Unmarshal(jsonBytes, &f)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}

	// JSON object parses into a map with string keys
	itemsMap := f.(map[string]interface{})

	// Loop through the Items; we're not interested in the key, just the values
	for _, v := range itemsMap {
		// Use type assertions to ensure that the value's a JSON object
		switch jsonObj := v.(type) {
		// The value is an Item, represented as a generic interface
		case interface{}:
			var item Item
			// Access the values in the JSON object and place them in an Item
			for itemKey, itemValue := range jsonObj.(map[string]interface{}) {
				switch itemKey {
				case "Item1":
					// Make sure that Item1 is a string
					switch itemValue := itemValue.(type) {
					case string:
						item.Item1 = itemValue
					default:
						fmt.Println("Incorrect type for", itemKey)
					}
				case "Item2":
					// Make sure that Item2 is a number; all numbers are transformed to float64
					switch itemValue := itemValue.(type) {
					case float64:
						item.Item2 = int(itemValue)
					default:
						fmt.Println("Incorrect type for", itemKey)
					}
				default:
					fmt.Println("Unknown key for Item found in JSON")
				}
			}
			fmt.Println(item)
			// Not a JSON object; handle the error
		default:
			fmt.Println("Expecting a JSON object; got something else")
		}
	}

}
