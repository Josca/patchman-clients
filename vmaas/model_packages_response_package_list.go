/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.3.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vmaas
// PackagesResponsePackageList struct for PackagesResponsePackageList
type PackagesResponsePackageList struct {
	Summary string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
	SourcePackage string `json:"source_package,omitempty"`
	PackageList []string `json:"package_list,omitempty"`
	Repositories []PackagesResponseRepositories `json:"repositories,omitempty"`
}