package main

import "testing"

func TestMapString(t *testing.T) {
	tags := []struct {
		Tag string
		Yup string
	}{
		{
			Tag: "validate:required,min=0,max=255",
			Yup: "yup.string().required().min(0).max(255)",
		},
		{
			Tag: "validate:required,min=0,max=255,email",
			Yup: "yup.string().required().email().min(0).max(255)",
		},
		{
			Tag: "validate:min=0,max=255",
			Yup: "yup.string().optional().min(0).max(255)",
		},
	}

	for _, testCase := range tags {
		if testCase.Yup != Yup().mapStringTag(testCase.Tag) {
			t.Errorf("failed to map tag=%s to yup expression", testCase.Tag)
		}
	}
}

func TestMapBool(t *testing.T) {
	tags := []struct {
		Tag string
		Yup string
	}{
		{
			Tag: "validate:required",
			Yup: "yup.bool().required()",
		},
		{
			Tag: "validate:something,required",
			Yup: "yup.bool().required()",
		},
		{
			Tag: "validate:",
			Yup: "yup.bool().optional()",
		},
	}

	for _, testCase := range tags {
		if testCase.Yup != Yup().mapBoolTag(testCase.Tag) {
			t.Errorf("failed to map tag=%s to yup expression", testCase.Tag)
		}
	}
}

func TestMapNumber(t *testing.T) {
	tags := []struct {
		Tag string
		Yup string
	}{
		{
			Tag: "validate:required,min=0,max=255",
			Yup: "yup.number().required().min(0).max(255)",
		},
		{
			Tag: "validate:min=0,max=255",
			Yup: "yup.number().optional().min(0).max(255)",
		},
	}

	for _, testCase := range tags {
		if testCase.Yup != Yup().mapNumberTag(testCase.Tag) {
			t.Errorf("failed to map tag=%s to yup expression", testCase.Tag)
		}
	}
}
