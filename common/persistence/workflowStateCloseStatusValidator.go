// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package persistence

import (
	"fmt"

	workflow "github.com/uber/cadence/.gen/go/shared"
)

// ValidateCreateWorkflowStateCloseStatus validate workflow state and close status
func ValidateCreateWorkflowStateCloseStatus(state int, closeStatus int) error {
	// validate workflow state & close status
	if state == WorkflowStateCompleted || closeStatus != WorkflowCloseStatusNone {
		return &workflow.InternalServiceError{
			Message: fmt.Sprintf("Create workflow with invalid state: %v or close status: %v",
				state, closeStatus),
		}
	}
	return nil
}

// ValidateUpdateWorkflowStateCloseStatus validate workflow state and close status
func ValidateUpdateWorkflowStateCloseStatus(state int, closeStatus int) error {
	// validate workflow state & close status
	if closeStatus == WorkflowCloseStatusNone {
		if state == WorkflowStateCompleted {
			return &workflow.InternalServiceError{
				Message: fmt.Sprintf("Update workflow with invalid state: %v or close status: %v",
					state, closeStatus),
			}
		}
	} else {
		// WorkflowCloseStatusCompleted
		// WorkflowCloseStatusFailed
		// WorkflowCloseStatusCanceled
		// WorkflowCloseStatusTerminated
		// WorkflowCloseStatusContinuedAsNew
		// WorkflowCloseStatusTimedOut
		if state != WorkflowStateCompleted {
			return &workflow.InternalServiceError{
				Message: fmt.Sprintf("Update workflow with invalid state: %v or close status: %v",
					state, closeStatus),
			}
		}
	}
	return nil
}
