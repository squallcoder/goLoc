package main

import (
	"fmt"
	"io"
	"net/http"
)

func main()  {
	resp, err := http.Get("https://www.loc.gov/?fo=json&at=trending_content")
	if err != nil {
	
	}
	
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	fmt.Printf("%s\n",body)

	
}
