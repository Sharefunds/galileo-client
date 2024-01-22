/*
 * Program API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type GetGroupHierarchyResponseData struct {
	// Identifier for the group
	GroupId string `json:"group_id"`
	// Name of the group
	GroupName string `json:"group_name"`
	// Object containing the next level down in the hierarchy
	Children []interface{} `json:"children"`
}