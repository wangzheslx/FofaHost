package service

type HostReq struct {
	Ip string `json:"ip"`
}

type HostResp struct {
	Host        string   `json:"host"`
	IP          string   `json:"ip"`
	Asn         int      `json:"asn"`
	Org         string   `json:"org"`
	CountryName string   `json:"country_name"`
	CountryCode string   `json:"country_code"`
	Protocol    []string `json:"protocol"`
	Port        []int    `json:"port"`
	Category    []string `json:"category"`
	Product     []string `json:"product"`
	UpdateTime  string   `json:"update_time"`
}
