package biz

type HostReq struct {
	Ip string `json:"ip"`
}

type HostResp struct {
	Error           bool     `json:"error"`
	Host            string   `json:"host"`
	ConsumedFpoint  int      `json:"consumed_fpoint"`
	RequiredFpoints int      `json:"required_fpoints"`
	IP              string   `json:"ip"`
	Asn             int      `json:"asn"`
	Org             string   `json:"org"`
	CountryName     string   `json:"country_name"`
	CountryCode     string   `json:"country_code"`
	Protocol        []string `json:"protocol"`
	Port            []int    `json:"port"`
	Category        []string `json:"category"`
	Product         []string `json:"product"`
	UpdateTime      string   `json:"update_time"`
}
