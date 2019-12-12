/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// HostOut A database entry representing a single host with its Inventory metadata.
type HostOut struct {
	// A Red Hat Account number that owns the host.
	Account string `json:"account"`
	// The ansible host name for remediations
	AnsibleHost *string `json:"ansible_host,omitempty"`
	// A UUID of the host machine BIOS.  This field is considered to be a canonical fact.
	BiosUuid *string `json:"bios_uuid,omitempty"`
	// A timestamp when the entry was created.
	Created string `json:"created,omitempty"`
	// Timestamp from which the host is considered deleted.
	CulledTimestamp *string `json:"culled_timestamp,omitempty"`
	// A host’s human-readable display name, e.g. in a form of a domain name.
	DisplayName *string `json:"display_name,omitempty"`
	// Host’s reference in the external source e.g. AWS EC2, Azure, OpenStack, etc. This field is considered to be a canonical fact.
	ExternalId *string `json:"external_id,omitempty"`
	// A set of facts belonging to the host.
	Facts []FactSet `json:"facts,omitempty"`
	// A host’s Fully Qualified Domain Name.  This field is considered to be a canonical fact.
	Fqdn *string `json:"fqdn,omitempty"`
	// A durable and reliable platform-wide host identifier. Applications should use this identifier to reference hosts.
	Id string `json:"id,omitempty"`
	// An ID defined in /etc/insights-client/machine-id. This field is considered a canonical fact.
	InsightsId *string `json:"insights_id,omitempty"`
	// Host’s network IP addresses.  This field is considered to be a canonical fact.
	IpAddresses *[]string `json:"ip_addresses,omitempty"`
	// Host’s network interfaces MAC addresses.  This field is considered to be a canonical fact.
	MacAddresses *[]string `json:"mac_addresses,omitempty"`
	// Reporting source of the host. Used when updating the stale_timestamp.
	Reporter *string `json:"reporter,omitempty"`
	// A Machine ID of a RHEL host.  This field is considered to be a canonical fact.
	RhelMachineId *string `json:"rhel_machine_id,omitempty"`
	// A Red Hat Satellite ID of a RHEL host.  This field is considered to be a canonical fact.
	SatelliteId *string `json:"satellite_id,omitempty"`
	// Timestamp from which the host is considered stale.
	StaleTimestamp *string `json:"stale_timestamp,omitempty"`
	// Timestamp from which the host is considered too stale to be listed without an explicit toggle.
	StaleWarningTimestamp *string `json:"stale_warning_timestamp,omitempty"`
	// A Red Hat Subcription Manager ID of a RHEL host.  This field is considered to be a canonical fact.
	SubscriptionManagerId *string `json:"subscription_manager_id,omitempty"`
	// An array of the tags on the host
	Tags []StructuredTag `json:"tags,omitempty"`
	// A timestamp when the entry was last updated.
	Updated string `json:"updated,omitempty"`
}
