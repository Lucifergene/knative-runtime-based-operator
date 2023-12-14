/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v2

// ObjectMetricSourceApplyConfiguration represents an declarative configuration of the ObjectMetricSource type for use
// with apply.
type ObjectMetricSourceApplyConfiguration struct {
	DescribedObject *CrossVersionObjectReferenceApplyConfiguration `json:"describedObject,omitempty"`
	Target          *MetricTargetApplyConfiguration                `json:"target,omitempty"`
	Metric          *MetricIdentifierApplyConfiguration            `json:"metric,omitempty"`
}

// ObjectMetricSourceApplyConfiguration constructs an declarative configuration of the ObjectMetricSource type for use with
// apply.
func ObjectMetricSource() *ObjectMetricSourceApplyConfiguration {
	return &ObjectMetricSourceApplyConfiguration{}
}

// WithDescribedObject sets the DescribedObject field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DescribedObject field is set to the value of the last call.
func (b *ObjectMetricSourceApplyConfiguration) WithDescribedObject(value *CrossVersionObjectReferenceApplyConfiguration) *ObjectMetricSourceApplyConfiguration {
	b.DescribedObject = value
	return b
}

// WithTarget sets the Target field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Target field is set to the value of the last call.
func (b *ObjectMetricSourceApplyConfiguration) WithTarget(value *MetricTargetApplyConfiguration) *ObjectMetricSourceApplyConfiguration {
	b.Target = value
	return b
}

// WithMetric sets the Metric field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Metric field is set to the value of the last call.
func (b *ObjectMetricSourceApplyConfiguration) WithMetric(value *MetricIdentifierApplyConfiguration) *ObjectMetricSourceApplyConfiguration {
	b.Metric = value
	return b
}
