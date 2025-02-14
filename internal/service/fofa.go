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

type HostDtlReq struct {
	Ip string `json:"ip"`
}
type HostDtlResp struct {
	Error           bool     `json:"error"`
	Host            string   `json:"host"`
	ConsumedFpoint  int      `json:"consumed_fpoint"`
	RequiredFpoints int      `json:"required_fpoints"`
	IP              string   `json:"ip"`
	Asn             int      `json:"asn"`
	Org             string   `json:"org"`
	CountryName     string   `json:"country_name"`
	CountryCode     string   `json:"country_code"`
	Domain          []string `json:"domain"`
	Ports           []Ports  `json:"ports"`
	UpdateTime      string   `json:"update_time"`
}
type Products struct {
	Product      string `json:"product"`
	Category     string `json:"category"`
	Level        int    `json:"level"`
	SortHardCode int    `json:"sort_hard_code"`
	Company      string `json:"company"`
}
type Ports struct {
	Port       int        `json:"port"`
	UpdateTime string     `json:"update_time"`
	Protocol   string     `json:"protocol"`
	Products   []Products `json:"products,omitempty"`
}
