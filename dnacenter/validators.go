package dnacenter

import (
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func validateStringIsValueFunc(reqValue string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (we []string, errors []error) {
		value := v.(string)
		if value != reqValue {
			errors = append(errors, fmt.Errorf("%q is an invalid value for argument %s. Only supported attribute is %q", value, k, reqValue))
		}
		return
	}
}

func validateStringHasValueFunc(values []string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (we []string, errors []error) {
		value := v.(string)
		valid := false
		for _, role := range values {
			if value == role {
				valid = true
				break
			}
		}

		if !valid {
			errors = append(errors, fmt.Errorf("%q is an invalid value for argument %s. Available values are %s", value, k, listNicely(values)))
		}
		return
	}
}

func validateIntegerInRange(min, max int) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		if value < min {
			errors = append(errors, fmt.Errorf(
				"%q cannot be lower than %d: %d", k, min, value))
		}
		if value > max {
			errors = append(errors, fmt.Errorf(
				"%q cannot be higher than %d: %d", k, max, value))
		}
		return
	}
}

func validateIntegerGeqThan(threshold int) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		value := v.(int)
		if value < threshold {
			errors = append(errors, fmt.Errorf(
				"%q cannot be lower than %d", k, threshold))
		}
		return
	}
}

func validateStringMatchesPattern(pattern string) schema.SchemaValidateFunc {
	return func(v interface{}, k string) (ws []string, errors []error) {
		compiledRegex, err := regexp.Compile(pattern)
		if err != nil {
			errors = append(errors, fmt.Errorf(
				"%q regex does not compile", pattern))
			return
		}

		value := v.(string)
		if !compiledRegex.MatchString(value) {
			errors = append(errors, fmt.Errorf(
				"%q doesn't match the pattern (%q): %q",
				k, pattern, value))
		}

		return
	}
}
