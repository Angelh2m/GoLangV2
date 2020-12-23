package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
    // Declare http client
    client := &http.Client{}

    // Declare post data
    // PostData := strings.NewReader("useId=5&age=12")

    // Declare HTTP Method and Url
    req, err := http.NewRequest("GET", "", nil)

    // Set cookie
    req.Header.Set("cookie", "")
    req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36")
    req.Header.Set("referer", "")
    req.Header.Set("authority", "")
	resp, err := client.Do(req)
	
	for _, cookie := range resp.Cookies() {
		fmt.Println("Found a cookie named:", cookie)
	  }
    // Read response
    data, err := ioutil.ReadAll(resp.Body)

    // error handle
    if err != nil {
        fmt.Printf("error = %s \n", err);
	}   
	
	_ = data
	
    // Print response
    // fmt.Printf("Response = %s", string(data));
}    