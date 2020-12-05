package dnacenter

import "testing"

func TestValidateStringIsValueFunc(t *testing.T) {
	reqValues := []string{"order", "des", "asc"}
	validStrings := []string{"order", "des", "asc"}
	invalidStrings := []string{"Order", "desc", "Asc"}

	for i, v := range validStrings {
		_, errors := validateStringIsValueFunc(reqValues[i])(v, "name")
		if len(errors) != 0 {
			t.Fatalf("%q should be equal to %q: %q", v, reqValues[i], errors)
		}
	}

	for i, v := range invalidStrings {
		_, errors := validateStringIsValueFunc(reqValues[i])(v, "name")
		if len(errors) == 0 {
			t.Fatalf("%q should not be equal to %q", v, reqValues[i])
		}
	}
}

func TestValidateStringHasValueFunc(t *testing.T) {
	reqValues := []string{"order", "des", "asc"}
	validStrings := []string{"order", "des", "asc"}
	invalidStrings := []string{"Order", "desc", "Asc"}
	for _, v := range validStrings {
		_, errors := validateStringHasValueFunc(reqValues)(v, "name")
		if len(errors) != 0 {
			t.Fatalf("%q should have one of these values %q: %q", v, reqValues, errors)
		}
	}

	for _, v := range invalidStrings {
		_, errors := validateStringHasValueFunc(reqValues)(v, "name")
		if len(errors) == 0 {
			t.Fatalf("%q should not have one of these values %q", v, reqValues)
		}
	}

}

func TestValidateIntegerInRange(t *testing.T) {
	validIntegers := []int{-259, 0, 1, 5, 999}
	min := -259
	max := 999
	for _, v := range validIntegers {
		_, errors := validateIntegerInRange(min, max)(v, "name")
		if len(errors) != 0 {
			t.Fatalf("%q should be an integer in range (%d, %d): %q", v, min, max, errors)
		}
	}

	invalidIntegers := []int{-260, -99999, 1000, 25678}
	for _, v := range invalidIntegers {
		_, errors := validateIntegerInRange(min, max)(v, "name")
		if len(errors) == 0 {
			t.Fatalf("%q should be an integer outside range (%d, %d)", v, min, max)
		}
	}
}

func TestValidateIntegerGeqThan0(t *testing.T) {
	v := 1
	if _, error := validateIntegerGeqThan(0)(v, "name"); error != nil {
		t.Fatalf("%q should be an integer greater than 0", v)
	}

	v = -4
	if _, error := validateIntegerGeqThan(0)(v, "name"); error == nil {
		t.Fatalf("%q should be an invalid integer smaller than 0", v)
	}
}

func TestValidateStringMatchesPattern(t *testing.T) {
	pattern := `^(pause|continue-mate|break)$`
	v := "pause"
	if _, error := validateStringMatchesPattern(pattern)(v, "name"); error != nil {
		t.Fatalf("%q should match the pattern", v)
	}
	v = "doesnotmatch"
	if _, error := validateStringMatchesPattern(pattern)(v, "name"); error == nil {
		t.Fatalf("%q should not match the pattern", v)
	}
	v = "continue-mate"
	if _, error := validateStringMatchesPattern(pattern)(v, "name"); error != nil {
		t.Fatalf("%q should match the pattern", v)
	}
}
