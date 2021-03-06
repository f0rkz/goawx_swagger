/*
 * Ansible Automation Platform controller API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Data3 struct {
	// Specify the type of credential you want to create. Refer to the documentation for details on each type.
	CredentialType int32 `json:"credential_type"`
	// 
	Description string `json:"description,omitempty"`
	// Enter inputs using either JSON or YAML syntax. Refer to the documentation for example syntax.
	Inputs interface{} `json:"inputs,omitempty"`
	// 
	Name string `json:"name"`
	// 
	Organization int32 `json:"organization,omitempty"`
}
