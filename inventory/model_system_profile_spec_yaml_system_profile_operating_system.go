/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// SystemProfileSpecYamlSystemProfileOperatingSystem Object for OS details
type SystemProfileSpecYamlSystemProfileOperatingSystem struct {
	// Major release of OS (aka the x version)
	Major int32 `json:"major,omitempty"`
	// Minor release of OS (aka the y version)
	Minor int32 `json:"minor,omitempty"`
	// Name of the distro/os
	Name string `json:"name,omitempty"`
}
