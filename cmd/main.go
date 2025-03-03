package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ughvj/takaaki/config"
	"github.com/ughvj/takaaki/drivers"
)

type NyutaikunApi struct {
	accessToken string
	client      *http.Client
}

type NyutaikunApiStudentsResult struct {
	Success bool                  `json:"success"`
	Total   int                   `json:"total"`
	Data    []NyutaikunApiStudent `json:"data"`
}

type NyutaikunApiStudent struct {
	NyutaikunUserId int    `json:"id"`
	Name            string `json:"name"`
}

func NewNyutaikunApi(accessToken string) *NyutaikunApi {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &NyutaikunApi{
		accessToken: accessToken,
		client:      client,
	}
}

func main() {
	conf := config.Loader.Get()
	fmt.Printf("%+v\n", conf)

	//
	req, _ := http.NewRequest(
		"GET",
		"https://site1.nyutai.com/api/chief/v1/students",
		nil,
	)

	req.Header.Add("Api-Token", conf.NyutaikunAccessToken)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("リクエストエラー: %s\n", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("レスポンス読み取りエラー: %s\n", err)
	}

	var result NyutaikunApiStudentsResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("JSONデコードエラー: %s\n", err)
	}

	//
	db, err := drivers.NewMysqlDriver()
	if err != nil {
		fmt.Printf("db connecting error")
	}

	for _, st := range result.Data {
		rows, _ := db.Use().Query("select * from students where nyutaikun_user_id = ?", st.NyutaikunUserId)
		if rows.Next() {
			fmt.Printf("NyutaikunUserId が既に存在しているため処理しませんでした: %s/%d\n", st.Name, st.NyutaikunUserId)
			continue
		}
		_, err = db.Use().Exec("insert into students (name, nyutaikun_user_id) values (?, ?)", st.Name, st.NyutaikunUserId)
		if err != nil {
			fmt.Printf("insertion error")
		}
		fmt.Printf("登録: %s/%d\n", st.Name, st.NyutaikunUserId)
	}
}
