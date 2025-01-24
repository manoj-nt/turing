package main

import (
	"fmt"
	"modela/paramserializer" // Your package's path.
)

func main() {
	rawQuery := "user_id=123&name=JohnDoe&age=30" +
		"&address[city]=NewYork&address[state]=NY" +
		"&address[coordinates][lat]=40.7128" +
		"&address[coordinates][lng]=-74.0060" +
		"&tags[]=go&tags[]=backend" +
		"&metadata[key1]=value1&metadata[key2]=value2"

	user, err := paramserializer.SerializeQueryParams(rawQuery)
	if err != nil {
		fmt.Printf("Error parsing query: %v\n", err)
		return
	}

	if err := paramserializer.ValidateUser(user); err != nil {
		fmt.Printf("Validation failed: %v\n", err)
		return
	}

	fmt.Printf("Parsed User: %+v\n", user)
}
