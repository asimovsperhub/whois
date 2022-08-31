package service

import (
	"dewhois/internal/model"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-ens/v3"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type sSearch struct {
}

//func NewDetails() Details {
//
//	return Details{snapshot:"Not set",url: "Not set",avatar: "Not set",twitter: "Not set",github: "Not set",email: "Not set",description: "Not set",notice: "Not set",
//		keywords: "Not set",discord: "Not set",reddit: "Not set",telegram: "Not set",delegate: "Not set"}
//}

func GetENSDomainList(address string) []string {
	client := &http.Client{}
	addr := strings.ToLower(address)
	var data = strings.NewReader(fmt.Sprintf(`{"operationName":"getRegistrations","variables":{"id":"%s","expiryDate":0,"skip":0,"first":1000},"query":"query getRegistrations($id: ID!, $first: Int, $skip: Int, $orderBy: Registration_orderBy, $orderDirection: OrderDirection, $expiryDate: Int) {\n  account(id: $id) {\n    registrations(first: $first, skip: $skip, orderBy: $orderBy, orderDirection: $orderDirection, where: {expiryDate_gt: $expiryDate}) {\n      expiryDate\n      domain {\n        id\n        labelName\n        labelhash\n        name\n        isMigrated\n        parent {\n          name\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n"}`, addr))
	req, err := http.NewRequest("POST", "https://api.thegraph.com/subgraphs/name/ensdomains/ens", data)
	if err != nil {
		log.Println("GetENSDomainList Request failed", err)
		return []string{}
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
		log.Println("GetENSDomainList pars resp failed", err)
		return []string{}
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("GetENSDomainList Read resp.Body failed", err)
		return []string{}
	}
	addrList := []string{}
	ens := make(map[string]map[string]map[string][]map[string]map[string]string)
	json.Unmarshal(bodyText, &ens)
	registrations := ens["data"]["account"]["registrations"]
	for _, v := range registrations {
		addrList = append(addrList, v["domain"]["name"])
	}
	return addrList
	//fmt.Printf("%s\n", addrList)

}

func PramsTools(domain string, client *ethclient.Client) (*model.Details, error) {
	defer client.Close()
	det := new(model.Details)
	name, _ := ens.NewName(client, domain)
	if name == nil {
		return det, nil
	}
	det.Domain = domain
	// 注册人地址
	registrant, _ := name.Registrant()
	det.Registrant = registrant
	// 到期时间
	expiration_date, _ := name.Expires()
	det.Expiration_date = expiration_date
	// 解析器地址
	resolver, _ := name.ResolverAddress()
	det.Resolver = resolver
	addr, _ := ens.NewResolver(client, domain)
	if addr == nil {
		return det, nil
	}
	// 解析记录地址
	addresses, _ := addr.Address()
	det.Addresses = addresses
	// 内容hash
	ContentHash, _ := addr.Contenthash()
	// 解码
	content, _ := ens.ContenthashToString(ContentHash)
	det.Content = content
	// 文本内容
	snapshot, _ := addr.Text("snapshot")
	det.Snapshot = snapshot
	url, _ := addr.Text("url")
	det.Url = url
	avatar, _ := addr.Text("avatar")
	det.Avatar = avatar
	description, _ := addr.Text("description")
	det.Description = description
	notice, _ := addr.Text("notice")
	det.Notice = notice
	keywords, _ := addr.Text("keywords")
	det.Keywords = keywords
	discord, _ := addr.Text("com.discord")
	det.Discord = discord
	github, _ := addr.Text("com.github")
	det.Github = github
	reddit, _ := addr.Text("com.reddit")
	det.Reddit = reddit
	twitter, _ := addr.Text("com.twitter")
	det.Twitter = twitter
	telegram, _ := addr.Text("org.telegram")
	det.Telegram = telegram
	delegate, _ := addr.Text("eth.ens.delegate")
	det.Delegate = delegate
	return det, nil
}

func PramsDomain(query string) *model.SearchIndex {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/5ac1f01fe23347fb9fd7ab0ceb2817cf")
	if err != nil {
		log.Println("ETHClient Dail failed ", err)
		return &model.SearchIndex{
			Details:             &model.Details{},
			Total:               0,
			NameList:            []string{},
			Registration_Status: 0,
			IsAddress:           0,
		}
	}
	// addr
	if len(query) == 42 && !strings.Contains(query, ".") {
		names := GetENSDomainList(query)
		out := &model.SearchIndex{
			&model.Details{},
			len(names),
			names,
			0,
			1,
		}
		return out
		// domain
	} else {
		domain := ""
		if !strings.Contains(query, ".eth") {
			domain = query + ".eth"
		} else {
			domain = query
		}
		data, _ := PramsTools(domain, client)
		if data.Registrant.String() == "0x0000000000000000000000000000000000000000" {
			out := &model.SearchIndex{
				data,
				0,
				[]string{},
				0,
				0,
			}
			return out
		} else {
			names := GetENSDomainList(data.Registrant.String())

			out := &model.SearchIndex{
				data,
				len(names),
				names,
				1,
				0,
			}
			return out
		}
	}
}
func Search() *sSearch {
	return &sSearch{}
}

func (s *sSearch) Search(query string) *model.SearchIndex {
	data := PramsDomain(query)
	return data
}
