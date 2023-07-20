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

package metric_alarm

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/cloudwatch"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/cloudwatch-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.CloudWatch{}
	_ = &svcapitypes.MetricAlarm{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadManyInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DescribeAlarmsOutput
	resp, err = rm.sdkapi.DescribeAlarmsWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeAlarms", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "UNKNOWN" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.CompositeAlarms {
		if elem.ActionsEnabled != nil {
			ko.Spec.ActionsEnabled = elem.ActionsEnabled
		} else {
			ko.Spec.ActionsEnabled = nil
		}
		if elem.AlarmActions != nil {
			f6 := []*string{}
			for _, f6iter := range elem.AlarmActions {
				var f6elem string
				f6elem = *f6iter
				f6 = append(f6, &f6elem)
			}
			ko.Spec.AlarmActions = f6
		} else {
			ko.Spec.AlarmActions = nil
		}
		if elem.AlarmDescription != nil {
			ko.Spec.AlarmDescription = elem.AlarmDescription
		} else {
			ko.Spec.AlarmDescription = nil
		}
		if elem.InsufficientDataActions != nil {
			f12 := []*string{}
			for _, f12iter := range elem.InsufficientDataActions {
				var f12elem string
				f12elem = *f12iter
				f12 = append(f12, &f12elem)
			}
			ko.Spec.InsufficientDataActions = f12
		} else {
			ko.Spec.InsufficientDataActions = nil
		}
		if elem.OKActions != nil {
			f13 := []*string{}
			for _, f13iter := range elem.OKActions {
				var f13elem string
				f13elem = *f13iter
				f13 = append(f13, &f13elem)
			}
			ko.Spec.OKActions = f13
		} else {
			ko.Spec.OKActions = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadManyInput returns true if there are any fields
// for the ReadMany Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadManyInput(
	r *resource,
) bool {
	return false
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeAlarmsInput, error) {
	res := &svcsdk.DescribeAlarmsInput{}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.PutMetricAlarmOutput
	_ = resp
	resp, err = rm.sdkapi.PutMetricAlarmWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "PutMetricAlarm", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.PutMetricAlarmInput, error) {
	res := &svcsdk.PutMetricAlarmInput{}

	if r.ko.Spec.ActionsEnabled != nil {
		res.SetActionsEnabled(*r.ko.Spec.ActionsEnabled)
	}
	if r.ko.Spec.AlarmActions != nil {
		f1 := []*string{}
		for _, f1iter := range r.ko.Spec.AlarmActions {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		res.SetAlarmActions(f1)
	}
	if r.ko.Spec.AlarmDescription != nil {
		res.SetAlarmDescription(*r.ko.Spec.AlarmDescription)
	}
	if r.ko.Spec.Name != nil {
		res.SetAlarmName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.ComparisonOperator != nil {
		res.SetComparisonOperator(*r.ko.Spec.ComparisonOperator)
	}
	if r.ko.Spec.DatapointsToAlarm != nil {
		res.SetDatapointsToAlarm(*r.ko.Spec.DatapointsToAlarm)
	}
	if r.ko.Spec.Dimensions != nil {
		f6 := []*svcsdk.Dimension{}
		for _, f6iter := range r.ko.Spec.Dimensions {
			f6elem := &svcsdk.Dimension{}
			if f6iter.Name != nil {
				f6elem.SetName(*f6iter.Name)
			}
			if f6iter.Value != nil {
				f6elem.SetValue(*f6iter.Value)
			}
			f6 = append(f6, f6elem)
		}
		res.SetDimensions(f6)
	}
	if r.ko.Spec.EvaluateLowSampleCountPercentile != nil {
		res.SetEvaluateLowSampleCountPercentile(*r.ko.Spec.EvaluateLowSampleCountPercentile)
	}
	if r.ko.Spec.EvaluationPeriods != nil {
		res.SetEvaluationPeriods(*r.ko.Spec.EvaluationPeriods)
	}
	if r.ko.Spec.ExtendedStatistic != nil {
		res.SetExtendedStatistic(*r.ko.Spec.ExtendedStatistic)
	}
	if r.ko.Spec.InsufficientDataActions != nil {
		f10 := []*string{}
		for _, f10iter := range r.ko.Spec.InsufficientDataActions {
			var f10elem string
			f10elem = *f10iter
			f10 = append(f10, &f10elem)
		}
		res.SetInsufficientDataActions(f10)
	}
	if r.ko.Spec.MetricName != nil {
		res.SetMetricName(*r.ko.Spec.MetricName)
	}
	if r.ko.Spec.Metrics != nil {
		f12 := []*svcsdk.MetricDataQuery{}
		for _, f12iter := range r.ko.Spec.Metrics {
			f12elem := &svcsdk.MetricDataQuery{}
			if f12iter.AccountID != nil {
				f12elem.SetAccountId(*f12iter.AccountID)
			}
			if f12iter.Expression != nil {
				f12elem.SetExpression(*f12iter.Expression)
			}
			if f12iter.ID != nil {
				f12elem.SetId(*f12iter.ID)
			}
			if f12iter.Label != nil {
				f12elem.SetLabel(*f12iter.Label)
			}
			if f12iter.MetricStat != nil {
				f12elemf4 := &svcsdk.MetricStat{}
				if f12iter.MetricStat.Metric != nil {
					f12elemf4f0 := &svcsdk.Metric{}
					if f12iter.MetricStat.Metric.Dimensions != nil {
						f12elemf4f0f0 := []*svcsdk.Dimension{}
						for _, f12elemf4f0f0iter := range f12iter.MetricStat.Metric.Dimensions {
							f12elemf4f0f0elem := &svcsdk.Dimension{}
							if f12elemf4f0f0iter.Name != nil {
								f12elemf4f0f0elem.SetName(*f12elemf4f0f0iter.Name)
							}
							if f12elemf4f0f0iter.Value != nil {
								f12elemf4f0f0elem.SetValue(*f12elemf4f0f0iter.Value)
							}
							f12elemf4f0f0 = append(f12elemf4f0f0, f12elemf4f0f0elem)
						}
						f12elemf4f0.SetDimensions(f12elemf4f0f0)
					}
					if f12iter.MetricStat.Metric.MetricName != nil {
						f12elemf4f0.SetMetricName(*f12iter.MetricStat.Metric.MetricName)
					}
					if f12iter.MetricStat.Metric.Namespace != nil {
						f12elemf4f0.SetNamespace(*f12iter.MetricStat.Metric.Namespace)
					}
					f12elemf4.SetMetric(f12elemf4f0)
				}
				if f12iter.MetricStat.Period != nil {
					f12elemf4.SetPeriod(*f12iter.MetricStat.Period)
				}
				if f12iter.MetricStat.Stat != nil {
					f12elemf4.SetStat(*f12iter.MetricStat.Stat)
				}
				if f12iter.MetricStat.Unit != nil {
					f12elemf4.SetUnit(*f12iter.MetricStat.Unit)
				}
				f12elem.SetMetricStat(f12elemf4)
			}
			if f12iter.Period != nil {
				f12elem.SetPeriod(*f12iter.Period)
			}
			if f12iter.ReturnData != nil {
				f12elem.SetReturnData(*f12iter.ReturnData)
			}
			f12 = append(f12, f12elem)
		}
		res.SetMetrics(f12)
	}
	if r.ko.Spec.Namespace != nil {
		res.SetNamespace(*r.ko.Spec.Namespace)
	}
	if r.ko.Spec.OKActions != nil {
		f14 := []*string{}
		for _, f14iter := range r.ko.Spec.OKActions {
			var f14elem string
			f14elem = *f14iter
			f14 = append(f14, &f14elem)
		}
		res.SetOKActions(f14)
	}
	if r.ko.Spec.Period != nil {
		res.SetPeriod(*r.ko.Spec.Period)
	}
	if r.ko.Spec.Statistic != nil {
		res.SetStatistic(*r.ko.Spec.Statistic)
	}
	if r.ko.Spec.Tags != nil {
		f17 := []*svcsdk.Tag{}
		for _, f17iter := range r.ko.Spec.Tags {
			f17elem := &svcsdk.Tag{}
			if f17iter.Key != nil {
				f17elem.SetKey(*f17iter.Key)
			}
			if f17iter.Value != nil {
				f17elem.SetValue(*f17iter.Value)
			}
			f17 = append(f17, f17elem)
		}
		res.SetTags(f17)
	}
	if r.ko.Spec.Threshold != nil {
		res.SetThreshold(*r.ko.Spec.Threshold)
	}
	if r.ko.Spec.ThresholdMetricID != nil {
		res.SetThresholdMetricId(*r.ko.Spec.ThresholdMetricID)
	}
	if r.ko.Spec.TreatMissingData != nil {
		res.SetTreatMissingData(*r.ko.Spec.TreatMissingData)
	}
	if r.ko.Spec.Unit != nil {
		res.SetUnit(*r.ko.Spec.Unit)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newUpdateRequestPayload(ctx, desired, delta)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.PutMetricAlarmOutput
	_ = resp
	resp, err = rm.sdkapi.PutMetricAlarmWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "PutMetricAlarm", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
	delta *ackcompare.Delta,
) (*svcsdk.PutMetricAlarmInput, error) {
	res := &svcsdk.PutMetricAlarmInput{}

	if r.ko.Spec.ActionsEnabled != nil {
		res.SetActionsEnabled(*r.ko.Spec.ActionsEnabled)
	}
	if r.ko.Spec.AlarmActions != nil {
		f1 := []*string{}
		for _, f1iter := range r.ko.Spec.AlarmActions {
			var f1elem string
			f1elem = *f1iter
			f1 = append(f1, &f1elem)
		}
		res.SetAlarmActions(f1)
	}
	if r.ko.Spec.AlarmDescription != nil {
		res.SetAlarmDescription(*r.ko.Spec.AlarmDescription)
	}
	if r.ko.Spec.Name != nil {
		res.SetAlarmName(*r.ko.Spec.Name)
	}
	if r.ko.Spec.ComparisonOperator != nil {
		res.SetComparisonOperator(*r.ko.Spec.ComparisonOperator)
	}
	if r.ko.Spec.DatapointsToAlarm != nil {
		res.SetDatapointsToAlarm(*r.ko.Spec.DatapointsToAlarm)
	}
	if r.ko.Spec.Dimensions != nil {
		f6 := []*svcsdk.Dimension{}
		for _, f6iter := range r.ko.Spec.Dimensions {
			f6elem := &svcsdk.Dimension{}
			if f6iter.Name != nil {
				f6elem.SetName(*f6iter.Name)
			}
			if f6iter.Value != nil {
				f6elem.SetValue(*f6iter.Value)
			}
			f6 = append(f6, f6elem)
		}
		res.SetDimensions(f6)
	}
	if r.ko.Spec.EvaluateLowSampleCountPercentile != nil {
		res.SetEvaluateLowSampleCountPercentile(*r.ko.Spec.EvaluateLowSampleCountPercentile)
	}
	if r.ko.Spec.EvaluationPeriods != nil {
		res.SetEvaluationPeriods(*r.ko.Spec.EvaluationPeriods)
	}
	if r.ko.Spec.ExtendedStatistic != nil {
		res.SetExtendedStatistic(*r.ko.Spec.ExtendedStatistic)
	}
	if r.ko.Spec.InsufficientDataActions != nil {
		f10 := []*string{}
		for _, f10iter := range r.ko.Spec.InsufficientDataActions {
			var f10elem string
			f10elem = *f10iter
			f10 = append(f10, &f10elem)
		}
		res.SetInsufficientDataActions(f10)
	}
	if r.ko.Spec.MetricName != nil {
		res.SetMetricName(*r.ko.Spec.MetricName)
	}
	if r.ko.Spec.Metrics != nil {
		f12 := []*svcsdk.MetricDataQuery{}
		for _, f12iter := range r.ko.Spec.Metrics {
			f12elem := &svcsdk.MetricDataQuery{}
			if f12iter.AccountID != nil {
				f12elem.SetAccountId(*f12iter.AccountID)
			}
			if f12iter.Expression != nil {
				f12elem.SetExpression(*f12iter.Expression)
			}
			if f12iter.ID != nil {
				f12elem.SetId(*f12iter.ID)
			}
			if f12iter.Label != nil {
				f12elem.SetLabel(*f12iter.Label)
			}
			if f12iter.MetricStat != nil {
				f12elemf4 := &svcsdk.MetricStat{}
				if f12iter.MetricStat.Metric != nil {
					f12elemf4f0 := &svcsdk.Metric{}
					if f12iter.MetricStat.Metric.Dimensions != nil {
						f12elemf4f0f0 := []*svcsdk.Dimension{}
						for _, f12elemf4f0f0iter := range f12iter.MetricStat.Metric.Dimensions {
							f12elemf4f0f0elem := &svcsdk.Dimension{}
							if f12elemf4f0f0iter.Name != nil {
								f12elemf4f0f0elem.SetName(*f12elemf4f0f0iter.Name)
							}
							if f12elemf4f0f0iter.Value != nil {
								f12elemf4f0f0elem.SetValue(*f12elemf4f0f0iter.Value)
							}
							f12elemf4f0f0 = append(f12elemf4f0f0, f12elemf4f0f0elem)
						}
						f12elemf4f0.SetDimensions(f12elemf4f0f0)
					}
					if f12iter.MetricStat.Metric.MetricName != nil {
						f12elemf4f0.SetMetricName(*f12iter.MetricStat.Metric.MetricName)
					}
					if f12iter.MetricStat.Metric.Namespace != nil {
						f12elemf4f0.SetNamespace(*f12iter.MetricStat.Metric.Namespace)
					}
					f12elemf4.SetMetric(f12elemf4f0)
				}
				if f12iter.MetricStat.Period != nil {
					f12elemf4.SetPeriod(*f12iter.MetricStat.Period)
				}
				if f12iter.MetricStat.Stat != nil {
					f12elemf4.SetStat(*f12iter.MetricStat.Stat)
				}
				if f12iter.MetricStat.Unit != nil {
					f12elemf4.SetUnit(*f12iter.MetricStat.Unit)
				}
				f12elem.SetMetricStat(f12elemf4)
			}
			if f12iter.Period != nil {
				f12elem.SetPeriod(*f12iter.Period)
			}
			if f12iter.ReturnData != nil {
				f12elem.SetReturnData(*f12iter.ReturnData)
			}
			f12 = append(f12, f12elem)
		}
		res.SetMetrics(f12)
	}
	if r.ko.Spec.Namespace != nil {
		res.SetNamespace(*r.ko.Spec.Namespace)
	}
	if r.ko.Spec.OKActions != nil {
		f14 := []*string{}
		for _, f14iter := range r.ko.Spec.OKActions {
			var f14elem string
			f14elem = *f14iter
			f14 = append(f14, &f14elem)
		}
		res.SetOKActions(f14)
	}
	if r.ko.Spec.Period != nil {
		res.SetPeriod(*r.ko.Spec.Period)
	}
	if r.ko.Spec.Statistic != nil {
		res.SetStatistic(*r.ko.Spec.Statistic)
	}
	if r.ko.Spec.Tags != nil {
		f17 := []*svcsdk.Tag{}
		for _, f17iter := range r.ko.Spec.Tags {
			f17elem := &svcsdk.Tag{}
			if f17iter.Key != nil {
				f17elem.SetKey(*f17iter.Key)
			}
			if f17iter.Value != nil {
				f17elem.SetValue(*f17iter.Value)
			}
			f17 = append(f17, f17elem)
		}
		res.SetTags(f17)
	}
	if r.ko.Spec.Threshold != nil {
		res.SetThreshold(*r.ko.Spec.Threshold)
	}
	if r.ko.Spec.ThresholdMetricID != nil {
		res.SetThresholdMetricId(*r.ko.Spec.ThresholdMetricID)
	}
	if r.ko.Spec.TreatMissingData != nil {
		res.SetTreatMissingData(*r.ko.Spec.TreatMissingData)
	}
	if r.ko.Spec.Unit != nil {
		res.SetUnit(*r.ko.Spec.Unit)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteAlarmsOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteAlarmsWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteAlarms", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteAlarmsInput, error) {
	res := &svcsdk.DeleteAlarmsInput{}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.MetricAlarm,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
