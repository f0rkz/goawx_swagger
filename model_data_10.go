/*
 * Ansible Automation Platform controller API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Data10 struct {
	// 
	Description string `json:"description,omitempty"`
	// 
	Inventory int32 `json:"inventory,omitempty"`
	// 
	Name string `json:"name,omitempty"`
	// Group variables in JSON or YAML format.
	Variables string `json:"variables,omitempty"`
}