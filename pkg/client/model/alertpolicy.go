package model

import "encoding/xml"

type AlertPolicy struct {
	XMLName              xml.Name `xml:"alert_policy"`
	PolicyName           string   `json:"policyName" xml:"policyName"`
	MetricType           string   `json:"metricType" xml:"metricType"`
	MetricName           string   `json:"metricName" xml:"metricName"`
	CreatedBy            string   `json:"createdBy" xml:"createdBy"`
	IsEnabled            string   `json:"isEnabled" xml:"isEnabled"`
	IsPerInstanceMetric  string   `json:"isPerInstanceMetric" xml:"isPerInstanceMetric"`
	Period               string   `json:"period" xml:"period"`
	PeriodUnits          string   `json:"periodUnits" xml:"periodUnits"`
	DatapointsToConsider string   `json:"datapointsToConsider" xml:"datapointsToConsider"`
	DatapointsToAlert    string   `json:"datapointsToAlert" xml:"datapointsToAlert"`
	Statistic            string   `json:"statistic" xml:"statistic"`
	Operator             string   `json:"operator" xml:"operator"`
	Condition            struct {
		ThresholdUnits string `json:"thresholdUnits" xml:"thresholdUnits"`
		ThresholdValue string `json:"thresholdValue" xml:"thresholdValue"`
		SeverityType   string `json:"severityType" xml:"severityType"`
	} `json:"condition" xml:"condition"`
}

// AlertPolicies is a list of alert policies
type AlertPolicies struct {
	// Items is the list of alert policies
	Items []AlertPolicy `json:"alert_policies" xml:"alert_policies"`

	// MaxBuckets is the maximum number of alert policies requested in the listing
	MaxPolicies int `json:"MaxPolicies,omitempty" xml:"MaxPolicies"`

	// NextMarker is a reference object to receive the next set of alert policies
	NextMarker string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`

	// Filter is a string query used to limit the returned alert policies in the
	// listing
	Filter string `json:"Filter,omitempty" xml:"Filter,omitempty"`

	// NextPageLink is a hyperlink to the next page in the alert policy listing
	NextPageLink string `json:"next_page_link,omitempty" xml:"next_page_link,omitempty"`
}
