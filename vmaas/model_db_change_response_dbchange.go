/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.20.4
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vmaas
// DbChangeResponseDbchange struct for DbChangeResponseDbchange
type DbChangeResponseDbchange struct {
	ErrataChanges string `json:"errata_changes,omitempty"`
	CveChanges string `json:"cve_changes,omitempty"`
	RepositoryChanges string `json:"repository_changes,omitempty"`
	LastChange string `json:"last_change,omitempty"`
	Exported string `json:"exported,omitempty"`
}
