/*
 * Ansible Automation Platform controller API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Data7 struct {
	// 
	Credential int32 `json:"credential,omitempty"`
	// 
	Description string `json:"description,omitempty"`
	// The full image location, including the container registry, image name, and version tag.
	Image string `json:"image,omitempty"`
	// 
	Name string `json:"name,omitempty"`
	// The organization used to determine access to this execution environment.
	Organization int32 `json:"organization,omitempty"`
	// Pull image before running?
	Pull string `json:"pull,omitempty"`
}