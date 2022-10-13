/*
 * Generated file models/prism/v4/config/config_model.go.
 *
 * Product version: 16.7.0.970-SNAPSHOT
 *
 * Part of the Service Manager for MSP
 *
 * (c) 2022 Nutanix Inc.  All rights reserved
 *
 */

/*
  Module prism.v4.config of Service Manager for MSP
*/
package config

/**
Reference to a task tracking the async operation.
*/
type TaskReference struct {
  
  ObjectType_ *string `json:"$objectType,omitempty"`
  
  Reserved_ map[string]interface{} `json:"$reserved,omitempty"`
  
  UnknownFields_ map[string]interface{} `json:"$unknownFields,omitempty"`
  /**
  Globally unique identifier of a task.
  */
  ExtId *string `json:"extId,omitempty"`
}


func NewTaskReference() *TaskReference {
  p := new(TaskReference)
  p.ObjectType_ = new(string)
  *p.ObjectType_ = "prism.v4.config.TaskReference"
  p.Reserved_ = map[string]interface{}{"$fqObjectType": "prism.v4.r0.a1.config.TaskReference"}
  p.UnknownFields_ = map[string]interface{}{}



  return p
}



