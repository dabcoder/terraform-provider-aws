// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package flex_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
)

func TestBoolFromFramework(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input    types.Bool
		expected *bool
	}
	tests := map[string]testCase{
		"true": {
			input:    types.BoolValue(true),
			expected: aws.Bool(true),
		},
		"false": {
			input:    types.BoolValue(false),
			expected: aws.Bool(false),
		},
		"null": {
			input:    types.BoolNull(),
			expected: nil,
		},
		"unknown": {
			input:    types.BoolUnknown(),
			expected: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := flex.BoolFromFramework(context.Background(), test.input)

			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestBoolValueFromFramework(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input    types.Bool
		expected bool
	}
	tests := map[string]testCase{
		"true": {
			input:    types.BoolValue(true),
			expected: true,
		},
		"false": {
			input:    types.BoolValue(false),
			expected: false,
		},
		"null": {
			input:    types.BoolNull(),
			expected: false,
		},
		"unknown": {
			input:    types.BoolUnknown(),
			expected: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := flex.BoolValueFromFramework(context.Background(), test.input)

			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestBoolToFramework(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input    *bool
		expected types.Bool
	}
	tests := map[string]testCase{
		"true": {
			input:    aws.Bool(true),
			expected: types.BoolValue(true),
		},
		"false": {
			input:    aws.Bool(false),
			expected: types.BoolValue(false),
		},
		"nil": {
			input:    nil,
			expected: types.BoolNull(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := flex.BoolToFramework(context.Background(), test.input)

			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}

func TestBoolToFrameworkLegacy(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input    *bool
		expected types.Bool
	}
	tests := map[string]testCase{
		"true": {
			input:    aws.Bool(true),
			expected: types.BoolValue(true),
		},
		"false": {
			input:    aws.Bool(false),
			expected: types.BoolValue(false),
		},
		"nil": {
			input:    nil,
			expected: types.BoolValue(false),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := flex.BoolToFrameworkLegacy(context.Background(), test.input)

			if diff := cmp.Diff(got, test.expected); diff != "" {
				t.Errorf("unexpected diff (+wanted, -got): %s", diff)
			}
		})
	}
}
