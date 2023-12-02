/*
 * rum
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

// HandlersAddPeerResult struct for HandlersAddPeerResult
type HandlersAddPeerResult struct {
	ErrCount int64 `json:"err_count,omitempty"`
	// Example: {\"/ip4/132.145.109.63/tcp/10666/p2p/16Uiu2HAmTovb8kAJiYK8saskzz7cRQhb45NRK5AsbtdmYsLfD3RM\": \"error info\"}
	Error     map[string]string `json:"error,omitempty"`
	SuccCount int64             `json:"succ_count,omitempty"`
}
