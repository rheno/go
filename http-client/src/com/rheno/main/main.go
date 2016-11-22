package main

import (

"fmt"
"com/rheno/util"

)

func main() {

	/* Set Header */
	var header = make(map[string]string)

	header["Content-Type"] = "application/json"

	url := "https://feeds.citibikenyc.com/stations/stations.json"


	/* Call The Request Function and get result */
	util.Request(func(s util.Success){
		fmt.Println(s.Result)
	},
	func(e util.Error){
		fmt.Println(e.Result)
	}, url, "GET", header, "")


}
