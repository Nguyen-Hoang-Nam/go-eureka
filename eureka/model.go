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
	HostName                      string                 `json:"hostName"`
	HomePageURL                   string                 `json:"homePageUrl,omitempty"`
	StatusPageURL                 string                 `json:"statusPageUrl"`
	HealthCheckURL                string                 `json:"healthCheckUrl,omitempty"`
	App                           string                 `json:"app"`
	IPAddr                        string                 `json:"ipAddr"`
	VipAddress                    string                 `json:"vipAddress"`
	SecureVipAddress              string                 `json:"secureVipAddress,omitempty"`
	Status                        string                 `json:"status"`
	Port                          Port                   `json:"port,omitempty"`
	SecurePort                    Port                   `json:"securePort,omitempty"`
	DataCenterInfo                *DataCenterInfo        `json:"dataCenterInfo"`
	LeaseInfo                     *LeaseInfo             `json:"leaseInfo,omitempty"`
	Metadata                      map[string]interface{} `json:"metadata,omitempty"`
	IsCoordinatingDiscoveryServer string                 `json:"isCoordinatingDiscoveryServer,omitempty"`
	LastUpdatedTimestamp          string                 `json:"lastUpdatedTimestamp,omitempty"`
	LastDirtyTimestamp            string                 `json:"lastDirtyTimestamp,omitempty"`
	ActionType                    string                 `json:"actionType,omitempty"`
	OverriddenStatus              string                 `json:"overriddenstatus,omitempty"`
	CountryID                     int                    `json:"countryId,omitempty"`
	InstanceId                    string                 `json:"instanceId,omitempty"`
}

type NewInstance struct {
	Instance Instance `json:"instance"`
}

type Port struct {
	Port    int    `json:"$"`
	Enabled string `json:"@enabled"`
}

type DataCenterInfo struct {
	Name     string              `json:"name"`
	Class    string              `json:"@class"`
	Metadata *DataCenterMetadata `json:"metadata,omitempty"`
}

type DataCenterMetadata struct {
	AmiLaunchIndex   string `json:"ami-launch-index,omitempty"`
	LocalHostname    string `json:"local-hostname,omitempty"`
	AvailabilityZone string `json:"availability-zone,omitempty"`
	InstanceID       string `json:"instance-id,omitempty"`
	PublicIpv4       string `json:"public-ipv4,omitempty"`
	PublicHostname   string `json:"public-hostname,omitempty"`
	AmiManifestPath  string `json:"ami-manifest-path,omitempty"`
	LocalIpv4        string `json:"local-ipv4,omitempty"`
	Hostname         string `json:"hostname,omitempty"`
	AmiID            string `json:"ami-id,omitempty"`
	InstanceType     string `json:"instance-type,omitempty"`
}

type LeaseInfo struct {
	RenewalIntervalInSecs int `json:"renewalIntervalInSecs,omitempty"`
	DurationInSecs        int `json:"durationInSecs,omitempty"`
}
