package main

import (
	"fmt"

	"github.com/grokify/spectrum/openapi3/openapi3postman2"
	"github.com/grokify/spectrum/postman2"
)

func main() {
	swagSpecFilepath := "scim-swagger-20170712.json"
	pmanBaseFilepath := "ringcentral.postman2.base.json"
	pmanSpecFilepath := "scim-postman-20170712.json"

	cfg := openapi3postman2.Configuration{
		PostmanURLHostname: "{{RINGCENTRAL_SERVER_HOSTNAME}}",
		PostmanHeaders: []postman2.Header{{
			Key:   "Authorization",
			Value: "Bearer {{my_access_token}}"}}}

	conv := openapi3postman2.NewConverter(cfg)

	merge := true
	var err error

	if merge {
		err = conv.MergeConvert(swagSpecFilepath, pmanBaseFilepath, pmanSpecFilepath)
	} else {
		err = conv.ConvertFile(swagSpecFilepath, pmanSpecFilepath)
	}

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Wrote %v\n", pmanSpecFilepath)
	}

	fmt.Println("DONE")
}
