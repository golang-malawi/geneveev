package main

import (
	"fmt"
	"strings"
)

type yup struct {
}

func Yup() *yup {
	return &yup{}
}

func (y *yup) jsImport() string {
	return "import * as yup from 'yup'"
}

func (y *yup) object() string {
	return "yup.object"
}

func (y *yup) mapStringTag(tagValue string) string {
	yupExprs := make([]string, 0)
	if strings.Index(tagValue, "required") != -1 {
		yupExprs = append(yupExprs, "required()")
	} else {
		yupExprs = append(yupExprs, "optional()")
	}

	if strings.Index(tagValue, "email") != -1 {
		yupExprs = append(yupExprs, "email()")
	}

	if strings.Index(tagValue, "min") != -1 {
		yupExprs = append(yupExprs, parseMinTagValue(tagValue))
	}

	if strings.Index(tagValue, "max") != -1 {
		yupExprs = append(yupExprs, parseMaxTagValue(tagValue))
	}

	return fmt.Sprintf("yup.string().%s", strings.Join(yupExprs, "."))
}

func (y *yup) mapBoolTag(tagValue string) string {
	yupExprs := make([]string, 0)
	if strings.Index(tagValue, "required") >= 0 {
		yupExprs = append(yupExprs, "required()")
	} else {
		yupExprs = append(yupExprs, "optional()")
	}
	return fmt.Sprintf("yup.bool().%s", strings.Join(yupExprs, "."))
}

func (y *yup) mapMixedFieldTag(tagValue string) string {
	yupExprs := make([]string, 0)
	if strings.Index(tagValue, "required") != -1 {
		yupExprs = append(yupExprs, "required()")
	} else {
		yupExprs = append(yupExprs, "optional()")
	}
	return fmt.Sprintf("yup.mixed().%s", strings.Join(yupExprs, "."))
}

func (y *yup) mapNumberTag(tagValue string) string {
	yupExprs := make([]string, 0)
	if strings.Index(tagValue, "required") != -1 {
		yupExprs = append(yupExprs, "required()")
	} else {
		yupExprs = append(yupExprs, "optional()")
	}

	if strings.Index(tagValue, "min") != -1 {
		yupExprs = append(yupExprs, parseMinTagValue(tagValue))
	}

	if strings.Index(tagValue, "max") != -1 {
		yupExprs = append(yupExprs, parseMaxTagValue(tagValue))
	}

	return fmt.Sprintf("yup.number().%s", strings.Join(yupExprs, "."))
}

func (y *yup) mapTimeStructTag(tagValue string) string {
	yupExprs := make([]string, 0)
	if strings.Index(tagValue, "required") != -1 {
		yupExprs = append(yupExprs, "required()")
	} else {
		yupExprs = append(yupExprs, "optional()")
	}

	if strings.Index(tagValue, "min") != -1 {
		yupExprs = append(yupExprs, parseMinTagValue(tagValue))
	}

	if strings.Index(tagValue, "max") != -1 {
		yupExprs = append(yupExprs, parseMaxTagValue(tagValue))
	}

	return fmt.Sprintf("yup.date().%s", strings.Join(yupExprs, "."))
}
