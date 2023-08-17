package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Pro
//const (
//    AppCode   = "fTf4pGcqYgzjp2JlgTJIfagoEk2byHGW"
//    AppToken  = "e4i5ylf6-id2r-itg8-x9nx-crhsbczo44b8"
//    URL       = "http://api.onelight.woa.com/api/uquery/sql/parser_table_lineage"
//)

// Test
const (
	AppCode  = "9gEbwOFXfUktCm1rsAhHETbYWGsWVBVs"
	AppToken = "28qqw351-w92v-dy6g-4ghi-muocd2rzc6ln"
	URL      = "http://apigateway.gzp-in.woa.com/api/uquery/lineage/ast_column"
)

func main() {
	for i := 0; i < 2; i++ {
		testCurl()
	}
}

func testCurl() {
	t := fmt.Sprintf("%d", time.Now().Unix())
	destination, source, timestamp, nonce, secret := "apigateway", AppCode, t, "1234567", AppToken
	signStr := strings.Join([]string{destination, source, timestamp, nonce, secret}, ",")
	sign := md5.Sum([]byte(signStr))
	signStr = hex.EncodeToString(sign[:])

	headers := http.Header{}
	headers.Add("X-Ntm-Destination-Service", destination)
	headers.Add("X-Ntm-Source-Service", source)
	headers.Add("X-Ntm-Timestamp", timestamp)
	headers.Add("X-Ntm-Nonce", nonce)
	headers.Add("X-Ntm-Signature", signStr)
	headers.Add("X-Onelight-User", "dinozhong")
	headers.Add("Content-Type", "application/json")

	data := `{
    "user_name": "",
    "sql_texts": [
        "use ieg_bdp",
        "select dteventtime,item_name,count(item_name) as nums,max(item_name) as max_nums from (select dteventtime,item_name from ft_dsl_itemflow_fht0) group by item_name,dteventtime"
]
}`

	req, err := http.NewRequest("POST", URL, strings.NewReader(data))
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}
	req.Header = headers

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		responseBody, _ := ioutil.ReadAll(response.Body)
		fmt.Println("Error response: ", string(responseBody))
		return
	}

	responseBody, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response: ", string(responseBody))

}
