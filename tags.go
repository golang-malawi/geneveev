package main

import (
	"fmt"
	"regexp"
	"strings"
)

func parseMinTagValue(tagValue string) string {
	r, _ := regexp.Compile(`min\:\d+`)
	found := r.FindString(tagValue)
	if found == "" {
		return ""
	}

	minExpr := strings.ReplaceAll(found, ":", "(")
	minExpr = minExpr + ")"
	return minExpr
}

func parseMaxTagValue(tagValue string) string {
	r, _ := regexp.Compile(`max\:\d+`)
	found := r.FindString(tagValue)
	if found == "" {
		return ""
	}

	maxExpr := strings.ReplaceAll(found, ":", "(")
	maxExpr = maxExpr + ")"
	return maxExpr
}

func mapStringTag(tagValue string) string {
	yupExprs := make([]string, 0)
	if strings.Index(tagValue, "required") >= 0 {
		yupExprs = append(yupExprs, "required()")
	} else {
		yupExprs = append(yupExprs, "optional()")
	}

	if strings.Index(tagValue, "email") >= 0 {
		yupExprs = append(yupExprs, "email()")
	}

	if strings.Index(tagValue, "min") >= 0 {
		yupExprs = append(yupExprs, parseMinTagValue(tagValue))
	}

	if strings.Index(tagValue, "max") >= 0 {
		yupExprs = append(yupExprs, parseMaxTagValue(tagValue))
	}

	return fmt.Sprintf("yup.string().%s", strings.Join(yupExprs, "."))
}

func mapBoolTag(tagValue string) string {
	yupExprs := make([]string, 0)
	if strings.Index(tagValue, "required") >= 0 {
		yupExprs = append(yupExprs, "required()")
	} else {
		yupExprs = append(yupExprs, "optional()")
	}
	return fmt.Sprintf("yup.bool().%s", strings.Join(yupExprs, "."))
}

func mapNumberTag(tagValue string) string {
	yupExprs := make([]string, 0)
	if strings.Index(tagValue, "required") >= 0 {
		yupExprs = append(yupExprs, "required()")
	} else {
		yupExprs = append(yupExprs, "optional()")
	}

	if strings.Index(tagValue, "min") >= 0 {
		yupExprs = append(yupExprs, parseMinTagValue(tagValue))
	}

	if strings.Index(tagValue, "max") >= 0 {
		yupExprs = append(yupExprs, parseMaxTagValue(tagValue))
	}

	return fmt.Sprintf("yup.number().%s", strings.Join(yupExprs, "."))
}

func mapTimeStructTag(tagValue string) string {
	yupExprs := make([]string, 0)
	if strings.Index(tagValue, "required") >= 0 {
		yupExprs = append(yupExprs, "required()")
	} else {
		yupExprs = append(yupExprs, "optional()")
	}

	if strings.Index(tagValue, "min") >= 0 {
		yupExprs = append(yupExprs, parseMinTagValue(tagValue))
	}

	if strings.Index(tagValue, "max") >= 0 {
		yupExprs = append(yupExprs, parseMaxTagValue(tagValue))
	}

	return fmt.Sprintf("yup.date().%s", strings.Join(yupExprs, "."))
}
