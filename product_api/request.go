/*
This code is to pull data from URL using REST API
Author: Meera <meeradevirj2000@gmail.com>
Date: 23/Sep/2022
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func main() {

	url := "https://world.openfoodfacts.org/api/v0/product/737628064502.json"

	req, _ := http.NewRequest("GET", url, nil)

	//req.Header.Add("X-API-Key", "XXX")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	var data = string(body)

	fmt.Println(data)
	fmt.Println("Type of data = ", reflect.TypeOf(data))

}
