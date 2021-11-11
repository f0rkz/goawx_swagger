/*
 * Ansible Automation Platform controller API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Data1 struct {
	// 
	Description string `json:"description,omitempty"`
	// 
	InputFieldName string `json:"input_field_name"`
	// 
	Metadata interface{} `json:"metadata,omitempty"`
	// 
	SourceCredential int32 `json:"source_credential"`
	// 
	TargetCredential int32 `json:"target_credential"`
}
