package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/broemp/broempSignal_Bots/model"
)

func AFK_create(userid int64, username string) (model.Response_afk_create, error) {
	endpoint := config.API_ENDPOINT + "/afk"

	req := model.Request_afk_create{Userid: userid, Username: username}

	jsonValue, err := json.Marshal(req)
	if err != nil {
		return model.Response_afk_create{}, err
	}

	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return model.Response_afk_create{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var afk model.Response_afk_create
	json.Unmarshal(body, &afk)

	return afk, nil
}

func AFK_get_top_list() []model.Response_afk_toplist {
	endpoint := config.API_ENDPOINT + "/afk/list"
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Println("failed to reach endpoint: ", err)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var top_list []model.Response_afk_toplist

	err = json.Unmarshal(body, &top_list)
	if err != nil {
		log.Println("failed to unmarshal api response: ", err)
	}

	return top_list
}

func AFK_get_user(userid int64) []model.Response_afk_user {
	endpoint := config.API_ENDPOINT + "/afk/list/" + strconv.Itoa(int(userid))
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Println("failed to reach endpoint: ", err)
		return nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var afk_list []model.Response_afk_user

	err = json.Unmarshal(body, &afk_list)
	if err != nil {
		log.Println("failed to unmarshal api response: ", err)
	}

	return afk_list
}
