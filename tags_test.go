package main

import "testing"

func TestMapString(t *testing.T) {
	tags := []struct {
		Tag string
		Yup string
	}{
		{
			Tag: "validate:required,min:0,max:255",
			Yup: "yup.string().required().min().max()",
		},
		{
			Tag: "validate:required,min:0,max:255,email",
			Yup: "yup.string().required().email().min().max()",
		},
		{
			Tag: "validate:min:0,max:255",
			Yup: "yup.string().optional().min().max()",
		},
	}

	for _, testCase := range tags {
		if testCase.Yup != mapStringTag(testCase.Tag) {
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
			Tag: "validate:",
			Yup: "yup.bool().optional()",
		},
	}

	for _, testCase := range tags {
		if testCase.Yup != mapBoolTag(testCase.Tag) {
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
			Tag: "validate:required,min:0,max:255",
			Yup: "yup.number().required().min().max()",
		},
		{
			Tag: "validate:min:0,max:255",
			Yup: "yup.number().optional().min().max()",
		},
	}

	for _, testCase := range tags {
		if testCase.Yup != mapNumberTag(testCase.Tag) {
			t.Errorf("failed to map tag=%s to yup expression", testCase.Tag)
		}
	}
}
