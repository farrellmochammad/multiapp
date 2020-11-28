package repository

import (
	"bytes"
	"encoding/json"
	model "go-service/model"
	"net/http"
)

func ReadArea() ([]model.Area, error) {
	resp, err := http.Get("https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var areaList []model.Area
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, &areaList); err != nil {
		return nil, err
	}

	return areaList, nil
}

func ReadConverter() (*model.Converter, error) {
	resp, err := http.Get("https://free.currconv.com/api/v7/convert?q=USD_IDR&compact=ultra&apiKey=971e1c9087790b741a01")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var converter model.Converter
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()
	if err := json.Unmarshal(respByte, (&converter)); err != nil {
		return nil, err
	}

	return &converter, nil
}
