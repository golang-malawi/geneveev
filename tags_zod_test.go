package main

import "testing"

func TestZodMapString(t *testing.T) {
	tags := []struct {
		Tag string
		Zod string
	}{
		{
			Tag: "validate:required,min=0,max=255",
			Zod: "z.string().required().min(0).max(255)",
		},
		{
			Tag: "validate:required,min=0,max=255,email",
			Zod: "z.string().required().email().min(0).max(255)",
		},
		{
			Tag: "validate:min=0,max=255",
			Zod: "z.string().optional().min(0).max(255)",
		},
	}

	for _, testCase := range tags {
		if testCase.Zod != Zod().mapStringTag(testCase.Tag) {
			t.Errorf("failed to map tag=%s to Zod expression", testCase.Tag)
		}
	}
}

func TestZodMapBool(t *testing.T) {
	tags := []struct {
		Tag string
		Zod string
	}{
		{
			Tag: "validate:required",
			Zod: "z.boolean().required()",
		},
		{
			Tag: "validate:something,required",
			Zod: "z.boolean().required()",
		},
		{
			Tag: "validate:",
			Zod: "z.boolean().optional()",
		},
	}

	for _, testCase := range tags {
		if testCase.Zod != Zod().mapBoolTag(testCase.Tag) {
			t.Errorf("failed to map tag=%s to Zod expression", testCase.Tag)
		}
	}
}

func TestZodMapNumber(t *testing.T) {
	tags := []struct {
		Tag string
		Zod string
	}{
		{
			Tag: "validate:required,min=0,max=255",
			Zod: "z.number().required().min(0).max(255)",
		},
		{
			Tag: "validate:min=0,max=255",
			Zod: "z.number().optional().min(0).max(255)",
		},
	}

	for _, testCase := range tags {
		if testCase.Zod != Zod().mapNumberTag(testCase.Tag) {
			t.Errorf("failed to map tag=%s to Zod expression", testCase.Tag)
		}
	}
}
