package main

import (
	"fmt"
	"strings"
)

type zod struct{}

func Zod() *zod {
	return &zod{}
}

func (z *zod) jsImport() string {
	return "import { z } from 'zod'"
}

func (z *zod) object() string {
	return "z.object"
}

func (z *zod) mapStringTag(tagValue string) string {
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

	return fmt.Sprintf("z.string().%s", strings.Join(yupExprs, "."))
}

func (z *zod) mapBoolTag(tagValue string) string {
	yupExprs := make([]string, 0)
	if strings.Index(tagValue, "required") >= 0 {
		yupExprs = append(yupExprs, "required()")
	} else {
		yupExprs = append(yupExprs, "optional()")
	}
	return fmt.Sprintf("z.boolean().%s", strings.Join(yupExprs, "."))
}

func (z *zod) mapMixedFieldTag(tagValue string) string {
	yupExprs := make([]string, 0)
	if strings.Index(tagValue, "required") != -1 {
		yupExprs = append(yupExprs, "required()")
	} else {
		yupExprs = append(yupExprs, "optional()")
	}
	return fmt.Sprintf("z.mixed().%s", strings.Join(yupExprs, "."))
}

func (z *zod) mapNumberTag(tagValue string) string {
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

	return fmt.Sprintf("z.number().%s", strings.Join(yupExprs, "."))
}

func (z *zod) mapTimeStructTag(tagValue string) string {
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

	return fmt.Sprintf("z.date().%s", strings.Join(yupExprs, "."))
}
