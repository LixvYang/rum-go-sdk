/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// HandlersNodeInfo struct for HandlersNodeInfo
type HandlersNodeInfo struct {
	Mem           HandlersNodeInfoMem `json:"mem,omitempty"`
	NodeId        string              `json:"node_id"`
	NodePublickey string              `json:"node_publickey"`
	NodeStatus    string              `json:"node_status"`
	NodeType      string              `json:"node_type"`
	NodeVersion   string              `json:"node_version"`
	// Example: {\"/quorum/nevis/meshsub/1.1.0\": [\"16Uiu2HAmM4jFjs5EjakvGgJkHS6Lg9jS6miNYPgJ3pMUvXGWXeTc\"]}
	Peers map[string][]string `json:"peers"`
}
