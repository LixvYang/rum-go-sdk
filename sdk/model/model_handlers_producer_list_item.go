/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// HandlersProducerListItem struct for HandlersProducerListItem
type HandlersProducerListItem struct {
	BlockWithness  int64  `json:"blockWithness,omitempty"`
	OwnerPubkey    string `json:"ownerPubkey,omitempty"`
	OwnerSign      string `json:"ownerSign,omitempty"`
	ProducerPubkey string `json:"producerPubkey,omitempty"`
	TimeStamp      int64  `json:"timeStamp,omitempty"`
}
