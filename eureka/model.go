package eureka

// Credit https://github.com/xuanbo/eureka-client/blob/master/config.go
type Applications struct {
	Applications []Application `json:"application"`
}

type Application struct {
	Name      string     `json:"name"`
	Instances []Instance `json:"instance"`
}

type Instance struct {
	HostName         string                 `json:"hostName"`
	HomePageURL      string                 `json:"homePageUrl"`
	StatusPageURL    string                 `json:"statusPageUrl"`
	App              string                 `json:"app"`
	IPAddr           string                 `json:"ipAddr"`
	VipAddress       string                 `json:"vipAddress"`
	SecureVipAddress string                 `json:"secureVipAddress"`
	Status           string                 `json:"status"`
	Port             Port                   `json:"port"`
	SecurePort       Port                   `json:"securePort"`
	Metadata         map[string]interface{} `json:"metadata"`
	InstanceId       string                 `json:"instanceId"`
}

type NewInstance struct {
	Instance Instance `json:"instance"`
}

type Port struct {
	Port    int    `json:"$"`
	Enabled string `json:"@enabled"`
}
