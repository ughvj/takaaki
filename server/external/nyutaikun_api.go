package external

import (
	"time"
	"net/http"
	"net/url"
	"fmt"
	"io/ioutil"
	"encoding/json"

	"github.com/ughvj/takaaki/types"
)

type NyutaikunApi struct {
	accessToken string
	client *http.Client 
}

func NewNyutaikunApi(accessToken string) *NyutaikunApi {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &NyutaikunApi{
		accessToken: accessToken,
		client: client,
	}
}

func (this *NyutaikunApi)RequestEntranceAndExit() *types.NyutaikunEntranceAndExitResult {
	baseURL := "https://site1.nyutai.com/api/chief/v1/entrance_and_exits"
	params := url.Values{}
	today := time.Now()
	params.Add("date_from", today.Format("2006-01-02"))
	params.Add("date_to", today.Format("2006-01-02"))
	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		fmt.Printf("リクエスト作成エラー: %s\n", err)
		return nil
	}

	req.Header.Add("Api-Token", this.accessToken)

	resp, err := this.client.Do(req)
	if err != nil {
		fmt.Printf("リクエストエラー: %s\n", err)
		return nil;
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("レスポンス読み取りエラー: %s\n", err)
		return nil
	}

	var result types.NyutaikunEntranceAndExitResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("JSONデコードエラー: %s\n", err)
		return nil
	}

	return &result
}
