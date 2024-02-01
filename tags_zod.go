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
	zodExprs := make([]string, 0)
	if strings.Index(tagValue, "required") != -1 {
		zodExprs = append(zodExprs, "required()")
	} else {
		zodExprs = append(zodExprs, "optional()")
	}

	if strings.Index(tagValue, "email") != -1 {
		zodExprs = append(zodExprs, "email()")
	}

	if strings.Index(tagValue, "min") != -1 {
		zodExprs = append(zodExprs, parseMinTagValue(tagValue))
	}

	if strings.Index(tagValue, "max") != -1 {
		zodExprs = append(zodExprs, parseMaxTagValue(tagValue))
	}

	return fmt.Sprintf("z.string().%s", strings.Join(zodExprs, "."))
}

func (z *zod) mapBoolTag(tagValue string) string {
	zodExprs := make([]string, 0)
	if strings.Index(tagValue, "required") >= 0 {
		zodExprs = append(zodExprs, "required()")
	} else {
		zodExprs = append(zodExprs, "optional()")
	}
	return fmt.Sprintf("z.boolean().%s", strings.Join(zodExprs, "."))
}

func (z *zod) mapMixedFieldTag(tagValue string) string {
	zodExprs := make([]string, 0)
	if strings.Index(tagValue, "required") != -1 {
		zodExprs = append(zodExprs, "required()")
	} else {
		zodExprs = append(zodExprs, "optional()")
	}
	return fmt.Sprintf("z.mixed().%s", strings.Join(zodExprs, "."))
}

func (z *zod) mapNumberTag(tagValue string) string {
	zodExprs := make([]string, 0)
	if strings.Index(tagValue, "required") != -1 {
		zodExprs = append(zodExprs, "required()")
	} else {
		zodExprs = append(zodExprs, "optional()")
	}

	if strings.Index(tagValue, "min") != -1 {
		zodExprs = append(zodExprs, parseMinTagValue(tagValue))
	}

	if strings.Index(tagValue, "max") != -1 {
		zodExprs = append(zodExprs, parseMaxTagValue(tagValue))
	}

	return fmt.Sprintf("z.number().%s", strings.Join(zodExprs, "."))
}

func (z *zod) mapTimeStructTag(tagValue string) string {
	zodExprs := make([]string, 0)
	if strings.Index(tagValue, "required") != -1 {
		zodExprs = append(zodExprs, "required()")
	} else {
		zodExprs = append(zodExprs, "optional()")
	}

	if strings.Index(tagValue, "min") != -1 {
		zodExprs = append(zodExprs, parseMinTagValue(tagValue))
	}

	if strings.Index(tagValue, "max") != -1 {
		zodExprs = append(zodExprs, parseMaxTagValue(tagValue))
	}

	return fmt.Sprintf("z.date().%s", strings.Join(zodExprs, "."))
}
