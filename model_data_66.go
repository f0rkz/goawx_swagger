/*
 * Ansible Automation Platform controller API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Data66 struct {
	// 
	Email string `json:"email,omitempty"`
	// 
	FirstName string `json:"first_name,omitempty"`
	// Designates that this user has all permissions without explicitly assigning them.
	IsSuperuser bool `json:"is_superuser,omitempty"`
	// 
	IsSystemAuditor bool `json:"is_system_auditor,omitempty"`
	// 
	LastName string `json:"last_name,omitempty"`
	// Write-only field used to change the password.
	Password string `json:"password,omitempty"`
	// Required. 150 characters or fewer. Letters, digits and @/./+/-/_ only.
	Username string `json:"username"`
}
