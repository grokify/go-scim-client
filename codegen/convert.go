package main

import (
	"fmt"

	"github.com/grokify/swaggman"
	"github.com/grokify/swaggman/postman2"
)

func main() {
	swagSpecFilepath := "scim-swagger-20170712.json"
	pmanBaseFilepath := "ringcentral.postman2.base.json"
	pmanSpecFilepath := "scim-postman-20170712.json"

	cfg := swaggman.Configuration{
		PostmanURLHostname: "{{RINGCENTRAL_SERVER_HOSTNAME}}",
		PostmanHeaders: []postman2.Header{{
			Key:   "Authorization",
			Value: "Bearer {{my_access_token}}"}}}

	conv := swaggman.NewConverter(cfg)

	merge := true
	var err error

	if merge {
		err = conv.MergeConvert(swagSpecFilepath, pmanBaseFilepath, pmanSpecFilepath)
	} else {
		err = conv.Convert(swagSpecFilepath, pmanSpecFilepath)
	}

	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Wrote %v\n", pmanSpecFilepath)
	}

	fmt.Println("DONE")
}
