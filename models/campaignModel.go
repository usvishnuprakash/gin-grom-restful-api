package models

import "gorm.io/gorm"

type Campaign struct {
	gorm.Model
	Id                int    `json:"id" grom:"primary_key"`
	AirByteUniqueKey  string `json:"_airbyte_unique_key"`
	MetricConversion  string `json:"metrics_conversions`
	CampaignTargetCpm string `json:"campaign_target_cpm" grom:"null"`
}
