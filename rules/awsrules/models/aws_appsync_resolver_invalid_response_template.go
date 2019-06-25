// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsAppsyncResolverInvalidResponseTemplateRule checks the pattern is valid
type AwsAppsyncResolverInvalidResponseTemplateRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsAppsyncResolverInvalidResponseTemplateRule returns new rule with default attributes
func NewAwsAppsyncResolverInvalidResponseTemplateRule() *AwsAppsyncResolverInvalidResponseTemplateRule {
	return &AwsAppsyncResolverInvalidResponseTemplateRule{
		resourceType:  "aws_appsync_resolver",
		attributeName: "response_template",
		max:           65536,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsAppsyncResolverInvalidResponseTemplateRule) Name() string {
	return "aws_appsync_resolver_invalid_response_template"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsAppsyncResolverInvalidResponseTemplateRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsAppsyncResolverInvalidResponseTemplateRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsAppsyncResolverInvalidResponseTemplateRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsAppsyncResolverInvalidResponseTemplateRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"response_template must be 65536 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"response_template must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}