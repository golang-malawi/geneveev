package main

import (
	"regexp"
	"strings"
)

func parseMinTagValue(tagValue string) string {
	r, _ := regexp.Compile(`min=\d+`)
	found := r.FindString(tagValue)
	if found == "" {
		return ""
	}

	minExpr := strings.ReplaceAll(found, "=", "(")
	minExpr = minExpr + ")"
	return minExpr
}

func parseMaxTagValue(tagValue string) string {
	r, _ := regexp.Compile(`max=\d+`)
	found := r.FindString(tagValue)
	if found == "" {
		return ""
	}

	maxExpr := strings.ReplaceAll(found, "=", "(")
	maxExpr = maxExpr + ")"
	return maxExpr
}
