package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

	func main() {
		client, err := alidns.NewClientWithAccessKey("cn-hangzhou", "LTAI4FkqAiunRYKeHgtrHike", "YebnUl8wLSiCCBJova0oG3YGRA1pbv")

		request := alidns.CreateAddDomainRecordRequest()
		request.Scheme = "https"

		request.Value = "139.9.214.171"
		request.Type = "A"
		request.RR = "s"
		request.DomainName = "letskr.com"

		response, err := client.AddDomainRecord(request)
		if err != nil {
			fmt.Print(err.Error())
		}
		fmt.Printf("response is %#v\n", response)
	}