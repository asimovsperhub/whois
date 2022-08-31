package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestGetEnsAddr_domain(t *testing.T) {
	client := &http.Client{}
	addr := strings.ToLower("0xb6E040C9ECAaE172a89bD561c5F73e1C48d28cd9")
	var data = strings.NewReader(fmt.Sprintf(`{"operationName":"getNamesFromSubgraph","variables":{"address":"%s"},"query":"query getNamesFromSubgraph($address: String!) {\n  domains(first: 1000, where: {resolvedAddress: $address}) {\n    name\n    __typename\n  }\n}\n"}`, addr))
	req, err := http.NewRequest("POST", "https://api.thegraph.com/subgraphs/name/ensdomains/ens", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.thegraph.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "https://app.ens.domains")
	req.Header.Set("referer", "https://app.ens.domains/")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	ens := make(map[string]map[string][]map[string]string)
	json.Unmarshal(bodyText, &ens)
	domainList := ens["data"]["domains"]
	doLi := []string{}
	for _, v := range domainList {
		doLi = append(doLi, v["name"])
	}
	fmt.Printf("%s\n", doLi)
}

func Test(t *testing.T) {
	client := &http.Client{}
	addr := strings.ToLower("0xAC9ba72fb61aA7c31A95df0A8b6ebA6f41EF875e")
	var data = strings.NewReader(fmt.Sprintf(`{"operationName":"getRegistrations","variables":{"id":"%s","expiryDate":0,"skip":0,"first":1000},"query":"query getRegistrations($id: ID!, $first: Int, $skip: Int, $orderBy: Registration_orderBy, $orderDirection: OrderDirection, $expiryDate: Int) {\n  account(id: $id) {\n    registrations(first: $first, skip: $skip, orderBy: $orderBy, orderDirection: $orderDirection, where: {expiryDate_gt: $expiryDate}) {\n      expiryDate\n      domain {\n        id\n        labelName\n        labelhash\n        name\n        isMigrated\n        parent {\n          name\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n"}`, addr))
	req, err := http.NewRequest("POST", "https://api.thegraph.com/subgraphs/name/ensdomains/ens", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.thegraph.com")
	req.Header.Set("accept", "*/*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("origin", "https://app.ens.domains")
	req.Header.Set("referer", "https://app.ens.domains/")
	req.Header.Set("sec-ch-ua", `" Not A;Brand";v="99", "Chromium";v="101", "Google Chrome";v="101"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	addrList := []string{}
	ens := make(map[string]map[string]map[string][]map[string]map[string]string)
	json.Unmarshal(bodyText, &ens)
	registrations := ens["data"]["account"]["registrations"]
	for _, v := range registrations {
		fmt.Println(v["domain"]["name"])
		addrList = append(addrList, v["domain"]["name"])
	}
	//return addrList
	fmt.Printf("%s\n", addrList)
}
