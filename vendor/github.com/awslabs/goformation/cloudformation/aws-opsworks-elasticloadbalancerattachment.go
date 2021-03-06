package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSOpsWorksElasticLoadBalancerAttachment AWS CloudFormation Resource (AWS::OpsWorks::ElasticLoadBalancerAttachment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html
type AWSOpsWorksElasticLoadBalancerAttachment struct {

	// ElasticLoadBalancerName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html#cfn-opsworks-elbattachment-elbname
	ElasticLoadBalancerName string `json:"ElasticLoadBalancerName,omitempty"`

	// LayerId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-elbattachment.html#cfn-opsworks-elbattachment-layerid
	LayerId string `json:"LayerId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksElasticLoadBalancerAttachment) AWSCloudFormationType() string {
	return "AWS::OpsWorks::ElasticLoadBalancerAttachment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksElasticLoadBalancerAttachment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSOpsWorksElasticLoadBalancerAttachment) MarshalJSON() ([]byte, error) {
	type Properties AWSOpsWorksElasticLoadBalancerAttachment
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSOpsWorksElasticLoadBalancerAttachment) UnmarshalJSON(b []byte) error {
	type Properties AWSOpsWorksElasticLoadBalancerAttachment
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSOpsWorksElasticLoadBalancerAttachment(*res.Properties)
	}

	return nil
}

// GetAllAWSOpsWorksElasticLoadBalancerAttachmentResources retrieves all AWSOpsWorksElasticLoadBalancerAttachment items from an AWS CloudFormation template
func (t *Template) GetAllAWSOpsWorksElasticLoadBalancerAttachmentResources() map[string]AWSOpsWorksElasticLoadBalancerAttachment {
	results := map[string]AWSOpsWorksElasticLoadBalancerAttachment{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSOpsWorksElasticLoadBalancerAttachment:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::OpsWorks::ElasticLoadBalancerAttachment" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSOpsWorksElasticLoadBalancerAttachment
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSOpsWorksElasticLoadBalancerAttachmentWithName retrieves all AWSOpsWorksElasticLoadBalancerAttachment items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksElasticLoadBalancerAttachmentWithName(name string) (AWSOpsWorksElasticLoadBalancerAttachment, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSOpsWorksElasticLoadBalancerAttachment:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::OpsWorks::ElasticLoadBalancerAttachment" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSOpsWorksElasticLoadBalancerAttachment
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSOpsWorksElasticLoadBalancerAttachment{}, errors.New("resource not found")
}
