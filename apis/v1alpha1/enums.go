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

type ActionsSuppressedBy string

const (
	ActionsSuppressedBy_WaitPeriod ActionsSuppressedBy = "WaitPeriod"
	ActionsSuppressedBy_ExtensionPeriod ActionsSuppressedBy = "ExtensionPeriod"
	ActionsSuppressedBy_Alarm ActionsSuppressedBy = "Alarm"
)

type AlarmType string

const (
	AlarmType_CompositeAlarm AlarmType = "CompositeAlarm"
	AlarmType_MetricAlarm AlarmType = "MetricAlarm"
)

type AnomalyDetectorStateValue string

const (
	AnomalyDetectorStateValue_PENDING_TRAINING AnomalyDetectorStateValue = "PENDING_TRAINING"
	AnomalyDetectorStateValue_TRAINED_INSUFFICIENT_DATA AnomalyDetectorStateValue = "TRAINED_INSUFFICIENT_DATA"
	AnomalyDetectorStateValue_TRAINED AnomalyDetectorStateValue = "TRAINED"
)

type AnomalyDetectorType string

const (
	AnomalyDetectorType_SINGLE_METRIC AnomalyDetectorType = "SINGLE_METRIC"
	AnomalyDetectorType_METRIC_MATH AnomalyDetectorType = "METRIC_MATH"
)

type ComparisonOperator string

const (
	ComparisonOperator_GreaterThanOrEqualToThreshold ComparisonOperator = "GreaterThanOrEqualToThreshold"
	ComparisonOperator_GreaterThanThreshold ComparisonOperator = "GreaterThanThreshold"
	ComparisonOperator_LessThanThreshold ComparisonOperator = "LessThanThreshold"
	ComparisonOperator_LessThanOrEqualToThreshold ComparisonOperator = "LessThanOrEqualToThreshold"
	ComparisonOperator_LessThanLowerOrGreaterThanUpperThreshold ComparisonOperator = "LessThanLowerOrGreaterThanUpperThreshold"
	ComparisonOperator_LessThanLowerThreshold ComparisonOperator = "LessThanLowerThreshold"
	ComparisonOperator_GreaterThanUpperThreshold ComparisonOperator = "GreaterThanUpperThreshold"
)

type EvaluationState string

const (
	EvaluationState_PARTIAL_DATA EvaluationState = "PARTIAL_DATA"
)

type HistoryItemType string

const (
	HistoryItemType_ConfigurationUpdate HistoryItemType = "ConfigurationUpdate"
	HistoryItemType_StateUpdate HistoryItemType = "StateUpdate"
	HistoryItemType_Action HistoryItemType = "Action"
)

type MetricStreamOutputFormat string

const (
	MetricStreamOutputFormat_json MetricStreamOutputFormat = "json"
	MetricStreamOutputFormat_opentelemetry0_7 MetricStreamOutputFormat = "opentelemetry0.7"
)

type RecentlyActive string

const (
	RecentlyActive_PT3H RecentlyActive = "PT3H"
)

type ScanBy string

const (
	ScanBy_TimestampDescending ScanBy = "TimestampDescending"
	ScanBy_TimestampAscending ScanBy = "TimestampAscending"
)

type StandardUnit string

const (
	StandardUnit_Seconds StandardUnit = "Seconds"
	StandardUnit_Microseconds StandardUnit = "Microseconds"
	StandardUnit_Milliseconds StandardUnit = "Milliseconds"
	StandardUnit_Bytes StandardUnit = "Bytes"
	StandardUnit_Kilobytes StandardUnit = "Kilobytes"
	StandardUnit_Megabytes StandardUnit = "Megabytes"
	StandardUnit_Gigabytes StandardUnit = "Gigabytes"
	StandardUnit_Terabytes StandardUnit = "Terabytes"
	StandardUnit_Bits StandardUnit = "Bits"
	StandardUnit_Kilobits StandardUnit = "Kilobits"
	StandardUnit_Megabits StandardUnit = "Megabits"
	StandardUnit_Gigabits StandardUnit = "Gigabits"
	StandardUnit_Terabits StandardUnit = "Terabits"
	StandardUnit_Percent StandardUnit = "Percent"
	StandardUnit_Count StandardUnit = "Count"
	StandardUnit_Bytes_Second StandardUnit = "Bytes/Second"
	StandardUnit_Kilobytes_Second StandardUnit = "Kilobytes/Second"
	StandardUnit_Megabytes_Second StandardUnit = "Megabytes/Second"
	StandardUnit_Gigabytes_Second StandardUnit = "Gigabytes/Second"
	StandardUnit_Terabytes_Second StandardUnit = "Terabytes/Second"
	StandardUnit_Bits_Second StandardUnit = "Bits/Second"
	StandardUnit_Kilobits_Second StandardUnit = "Kilobits/Second"
	StandardUnit_Megabits_Second StandardUnit = "Megabits/Second"
	StandardUnit_Gigabits_Second StandardUnit = "Gigabits/Second"
	StandardUnit_Terabits_Second StandardUnit = "Terabits/Second"
	StandardUnit_Count_Second StandardUnit = "Count/Second"
	StandardUnit_None StandardUnit = "None"
)

type StateValue string

const (
	StateValue_OK StateValue = "OK"
	StateValue_ALARM StateValue = "ALARM"
	StateValue_INSUFFICIENT_DATA StateValue = "INSUFFICIENT_DATA"
)

type Statistic string

const (
	Statistic_SampleCount Statistic = "SampleCount"
	Statistic_Average Statistic = "Average"
	Statistic_Sum Statistic = "Sum"
	Statistic_Minimum Statistic = "Minimum"
	Statistic_Maximum Statistic = "Maximum"
)

type StatusCode string

const (
	StatusCode_Complete StatusCode = "Complete"
	StatusCode_InternalError StatusCode = "InternalError"
	StatusCode_PartialData StatusCode = "PartialData"
	StatusCode_Forbidden StatusCode = "Forbidden"
)