/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// PbTrx struct for PbTrx
type PbTrx struct {
	Data         []int64 `json:"Data,omitempty"`
	Expired      int64   `json:"Expired,omitempty"`
	GroupId      string  `json:"GroupId,omitempty"`
	ResendCount  int64   `json:"ResendCount,omitempty"`
	SenderPubkey string  `json:"SenderPubkey,omitempty"`
	SenderSign   []int64 `json:"SenderSign,omitempty"`
	StorageType  int64   `json:"StorageType,omitempty"`
	TimeStamp    string  `json:"TimeStamp,omitempty"`
	TrxId        string  `json:"TrxId,omitempty"`
	Type         int64   `json:"Type,omitempty"`
	Version      string  `json:"Version,omitempty"`
}
