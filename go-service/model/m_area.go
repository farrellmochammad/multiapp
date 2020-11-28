package models

type Area struct {
	Uuid         string `json:"uuid"`
	Komoditas    string `json:"komoditas"`
	AreaProvinsi string `json:"area_provinsi"`
	AreaKota     string `json:"area_kota"`
	Size         string `json:"size"`
	Price        string `json:"price"`
	UsdPrice     string `json:"usd_price"`
	TglParsed    string `json:"tgl_parsed"`
	Timestamp    string `json:"timestamp"`
}

type Converter struct {
	UsdIdr float64 `json:"USD_IDR"`
}

type Fetch struct {
	AreaProvinsi string `form:"area_provinsi",json:"area_provinsi"`
	Week         string `form:"week", json:"week"`
}
