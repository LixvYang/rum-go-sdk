/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// ApiGroupInfo struct for ApiGroupInfo
type ApiGroupInfo struct {
	AppKey          string           `json:"app_key"`
	CipherKey       string           `json:"cipher_key"`
	ConsensusType   string           `json:"consensus_type"`
	CurrtEpoch      int64            `json:"currt_epoch"`
	CurrtTopBlock   int64            `json:"currt_top_block"`
	EncryptionType  string           `json:"encryption_type"`
	GroupId         string           `json:"group_id"`
	GroupName       string           `json:"group_name"`
	LastUpdated     int64            `json:"last_updated"`
	OwnerPubkey     string           `json:"owner_pubkey"`
	Peers           []string         `json:"peers"`
	RexSyncerResult DefRexSyncResult `json:"rex_Syncer_result"`
	RexSyncerStatus string           `json:"rex_syncer_status"`
	UserEthAddr     string           `json:"user_eth_addr"`
	UserPubkey      string           `json:"user_pubkey"`
}