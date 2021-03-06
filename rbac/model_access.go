/*
 * Role Based Access Control
 *
 * The API for Role Based Access Control.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rbac
// Access struct for Access
type Access struct {
	Permission string `json:"permission"`
	ResourceDefinitions []ResourceDefinition `json:"resourceDefinitions"`
}
