/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// HandlersAnnouncedProducerListItem struct for HandlersAnnouncedProducerListItem
type HandlersAnnouncedProducerListItem struct {
	// ACTION have 2 states: \"ADD\" means Producer is normal, \"REMOVE\" means Producer has announced to leave the group by itself
	Action          string `json:"action"`
	AnnouncedPubkey string `json:"announcedPubkey"`
	AnnouncerSign   string `json:"announcerSign"`
	Memo            string `json:"memo"`
	Result          string `json:"result"`
	TimeStamp       int64  `json:"timeStamp"`
}
