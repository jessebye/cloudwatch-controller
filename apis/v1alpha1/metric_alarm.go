// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MetricAlarmSpec defines the desired state of MetricAlarm.
//
// The details about a metric alarm.
type MetricAlarmSpec struct {

	// Indicates whether actions should be executed during any changes to the alarm
	// state. The default is TRUE.
	ActionsEnabled *bool `json:"actionsEnabled,omitempty"`
	// The actions to execute when this alarm transitions to the ALARM state from
	// any other state. Each action is specified as an Amazon Resource Name (ARN).
	// Valid values:
	//
	// EC2 actions:
	//
	//   - arn:aws:automate:region:ec2:stop
	//
	//   - arn:aws:automate:region:ec2:terminate
	//
	//   - arn:aws:automate:region:ec2:reboot
	//
	//   - arn:aws:automate:region:ec2:recover
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Recover/1.0
	//
	// Autoscaling action:
	//
	//   - arn:aws:autoscaling:region:account-id:scalingPolicy:policy-id:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name
	//
	// SNS notification action:
	//
	//   - arn:aws:sns:region:account-id:sns-topic-name:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name
	//
	// SSM integration actions:
	//
	//   - arn:aws:ssm:region:account-id:opsitem:severity#CATEGORY=category-name
	//
	//   - arn:aws:ssm-incidents::account-id:responseplan/response-plan-name
	AlarmActions []*string `json:"alarmActions,omitempty"`
	// The description for the alarm.
	AlarmDescription *string `json:"alarmDescription,omitempty"`
	// The arithmetic operation to use when comparing the specified statistic and
	// threshold. The specified statistic value is used as the first operand.
	//
	// The values LessThanLowerOrGreaterThanUpperThreshold, LessThanLowerThreshold,
	// and GreaterThanUpperThreshold are used only for alarms based on anomaly detection
	// models.
	// +kubebuilder:validation:Required
	ComparisonOperator *string `json:"comparisonOperator"`
	// The number of data points that must be breaching to trigger the alarm. This
	// is used only if you are setting an "M out of N" alarm. In that case, this
	// value is the M. For more information, see Evaluating an Alarm (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarm-evaluation)
	// in the Amazon CloudWatch User Guide.
	DatapointsToAlarm *int64 `json:"datapointsToAlarm,omitempty"`
	// The dimensions for the metric specified in MetricName.
	Dimensions []*Dimension `json:"dimensions,omitempty"`
	// Used only for alarms based on percentiles. If you specify ignore, the alarm
	// state does not change during periods with too few data points to be statistically
	// significant. If you specify evaluate or omit this parameter, the alarm is
	// always evaluated and possibly changes state no matter how many data points
	// are available. For more information, see Percentile-Based CloudWatch Alarms
	// and Low Data Samples (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#percentiles-with-low-samples).
	//
	// Valid Values: evaluate | ignore
	EvaluateLowSampleCountPercentile *string `json:"evaluateLowSampleCountPercentile,omitempty"`
	// The number of periods over which data is compared to the specified threshold.
	// If you are setting an alarm that requires that a number of consecutive data
	// points be breaching to trigger the alarm, this value specifies that number.
	// If you are setting an "M out of N" alarm, this value is the N.
	//
	// An alarm's total current evaluation period can be no longer than one day,
	// so this number multiplied by Period cannot be more than 86,400 seconds.
	// +kubebuilder:validation:Required
	EvaluationPeriods *int64 `json:"evaluationPeriods"`
	// The percentile statistic for the metric specified in MetricName. Specify
	// a value between p0.0 and p100. When you call PutMetricAlarm and specify a
	// MetricName, you must specify either Statistic or ExtendedStatistic, but not
	// both.
	ExtendedStatistic *string `json:"extendedStatistic,omitempty"`
	// The actions to execute when this alarm transitions to the INSUFFICIENT_DATA
	// state from any other state. Each action is specified as an Amazon Resource
	// Name (ARN). Valid values:
	//
	// EC2 actions:
	//
	//   - arn:aws:automate:region:ec2:stop
	//
	//   - arn:aws:automate:region:ec2:terminate
	//
	//   - arn:aws:automate:region:ec2:reboot
	//
	//   - arn:aws:automate:region:ec2:recover
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Recover/1.0
	//
	// Autoscaling action:
	//
	//   - arn:aws:autoscaling:region:account-id:scalingPolicy:policy-id:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name
	//
	// SNS notification action:
	//
	//   - arn:aws:sns:region:account-id:sns-topic-name:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name
	//
	// SSM integration actions:
	//
	//   - arn:aws:ssm:region:account-id:opsitem:severity#CATEGORY=category-name
	//
	//   - arn:aws:ssm-incidents::account-id:responseplan/response-plan-name
	InsufficientDataActions []*string `json:"insufficientDataActions,omitempty"`
	// The name for the metric associated with the alarm. For each PutMetricAlarm
	// operation, you must specify either MetricName or a Metrics array.
	//
	// If you are creating an alarm based on a math expression, you cannot specify
	// this parameter, or any of the Dimensions, Period, Namespace, Statistic, or
	// ExtendedStatistic parameters. Instead, you specify all this information in
	// the Metrics array.
	MetricName *string `json:"metricName,omitempty"`
	// An array of MetricDataQuery structures that enable you to create an alarm
	// based on the result of a metric math expression. For each PutMetricAlarm
	// operation, you must specify either MetricName or a Metrics array.
	//
	// Each item in the Metrics array either retrieves a metric or performs a math
	// expression.
	//
	// One item in the Metrics array is the expression that the alarm watches. You
	// designate this expression by setting ReturnData to true for this object in
	// the array. For more information, see MetricDataQuery (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDataQuery.html).
	//
	// If you use the Metrics parameter, you cannot include the MetricName, Dimensions,
	// Period, Namespace, Statistic, or ExtendedStatistic parameters of PutMetricAlarm
	// in the same operation. Instead, you retrieve the metrics you are using in
	// your math expression as part of the Metrics array.
	Metrics []*MetricDataQuery `json:"metrics,omitempty"`
	// The name for the alarm. This name must be unique within the Region.
	//
	// The name must contain only UTF-8 characters, and can't contain ASCII control
	// characters
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The namespace for the metric associated specified in MetricName.
	Namespace *string `json:"namespace,omitempty"`
	// The actions to execute when this alarm transitions to an OK state from any
	// other state. Each action is specified as an Amazon Resource Name (ARN). Valid
	// values:
	//
	// EC2 actions:
	//
	//   - arn:aws:automate:region:ec2:stop
	//
	//   - arn:aws:automate:region:ec2:terminate
	//
	//   - arn:aws:automate:region:ec2:reboot
	//
	//   - arn:aws:automate:region:ec2:recover
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Stop/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Terminate/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Reboot/1.0
	//
	//   - arn:aws:swf:region:account-id:action/actions/AWS_EC2.InstanceId.Recover/1.0
	//
	// Autoscaling action:
	//
	//   - arn:aws:autoscaling:region:account-id:scalingPolicy:policy-id:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name
	//
	// SNS notification action:
	//
	//   - arn:aws:sns:region:account-id:sns-topic-name:autoScalingGroupName/group-friendly-name:policyName/policy-friendly-name
	//
	// SSM integration actions:
	//
	//   - arn:aws:ssm:region:account-id:opsitem:severity#CATEGORY=category-name
	//
	//   - arn:aws:ssm-incidents::account-id:responseplan/response-plan-name
	OKActions []*string `json:"oKActions,omitempty"`
	// The length, in seconds, used each time the metric specified in MetricName
	// is evaluated. Valid values are 10, 30, and any multiple of 60.
	//
	// Period is required for alarms based on static thresholds. If you are creating
	// an alarm based on a metric math expression, you specify the period for each
	// metric within the objects in the Metrics array.
	//
	// Be sure to specify 10 or 30 only for metrics that are stored by a PutMetricData
	// call with a StorageResolution of 1. If you specify a period of 10 or 30 for
	// a metric that does not have sub-minute resolution, the alarm still attempts
	// to gather data at the period rate that you specify. In this case, it does
	// not receive data for the attempts that do not correspond to a one-minute
	// data resolution, and the alarm might often lapse into INSUFFICENT_DATA status.
	// Specifying 10 or 30 also sets this alarm as a high-resolution alarm, which
	// has a higher charge than other alarms. For more information about pricing,
	// see Amazon CloudWatch Pricing (https://aws.amazon.com/cloudwatch/pricing/).
	//
	// An alarm's total current evaluation period can be no longer than one day,
	// so Period multiplied by EvaluationPeriods cannot be more than 86,400 seconds.
	Period *int64 `json:"period,omitempty"`
	// The statistic for the metric specified in MetricName, other than percentile.
	// For percentile statistics, use ExtendedStatistic. When you call PutMetricAlarm
	// and specify a MetricName, you must specify either Statistic or ExtendedStatistic,
	// but not both.
	Statistic *string `json:"statistic,omitempty"`
	// A list of key-value pairs to associate with the alarm. You can associate
	// as many as 50 tags with an alarm.
	//
	// Tags can help you organize and categorize your resources. You can also use
	// them to scope user permissions by granting a user permission to access or
	// change only resources with certain tag values.
	//
	// If you are using this operation to update an existing alarm, any tags you
	// specify in this parameter are ignored. To change the tags of an existing
	// alarm, use TagResource (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_TagResource.html)
	// or UntagResource (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_UntagResource.html).
	Tags []*Tag `json:"tags,omitempty"`
	// The value against which the specified statistic is compared.
	//
	// This parameter is required for alarms based on static thresholds, but should
	// not be used for alarms based on anomaly detection models.
	Threshold *float64 `json:"threshold,omitempty"`
	// If this is an alarm based on an anomaly detection model, make this value
	// match the ID of the ANOMALY_DETECTION_BAND function.
	//
	// For an example of how to use this parameter, see the Anomaly Detection Model
	// Alarm example on this page.
	//
	// If your alarm uses this parameter, it cannot have Auto Scaling actions.
	ThresholdMetricID *string `json:"thresholdMetricID,omitempty"`
	// Sets how this alarm is to handle missing data points. If TreatMissingData
	// is omitted, the default behavior of missing is used. For more information,
	// see Configuring How CloudWatch Alarms Treats Missing Data (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/AlarmThatSendsEmail.html#alarms-and-missing-data).
	//
	// Valid Values: breaching | notBreaching | ignore | missing
	//
	// Alarms that evaluate metrics in the AWS/DynamoDB namespace always ignore
	// missing data even if you choose a different option for TreatMissingData.
	// When an AWS/DynamoDB metric has missing data, alarms that evaluate that metric
	// remain in their current state.
	TreatMissingData *string `json:"treatMissingData,omitempty"`
	// The unit of measure for the statistic. For example, the units for the Amazon
	// EC2 NetworkIn metric are Bytes because NetworkIn tracks the number of bytes
	// that an instance receives on all network interfaces. You can also specify
	// a unit when you create a custom metric. Units help provide conceptual meaning
	// to your data. Metric data points that specify a unit of measure, such as
	// Percent, are aggregated separately.
	//
	// If you don't specify Unit, CloudWatch retrieves all unit types that have
	// been published for the metric and attempts to evaluate the alarm. Usually,
	// metrics are published with only one unit, so the alarm works as intended.
	//
	// However, if the metric is published with multiple types of units and you
	// don't specify a unit, the alarm's behavior is not defined and it behaves
	// unpredictably.
	//
	// We recommend omitting Unit so that you don't inadvertently specify an incorrect
	// unit that is not published for this metric. Doing so causes the alarm to
	// be stuck in the INSUFFICIENT DATA state.
	Unit *string `json:"unit,omitempty"`
}

// MetricAlarmStatus defines the observed state of MetricAlarm
type MetricAlarmStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
}

// MetricAlarm is the Schema for the MetricAlarms API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type MetricAlarm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MetricAlarmSpec   `json:"spec,omitempty"`
	Status            MetricAlarmStatus `json:"status,omitempty"`
}

// MetricAlarmList contains a list of MetricAlarm
// +kubebuilder:object:root=true
type MetricAlarmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetricAlarm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MetricAlarm{}, &MetricAlarmList{})
}
