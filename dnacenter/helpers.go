package dnacenter

import (
	"fmt"
	"strings"
)

// listNicely listNicely
/* Converts []string to string, by adding quotes and separate values by comma
@param values
*/
func listNicely(values []string) string {
	pvalues := fmt.Sprintf("%q", values)
	pvalues = pvalues[1 : len(pvalues)-1]
	return strings.Join(strings.Split(pvalues, " "), ", ")
}
