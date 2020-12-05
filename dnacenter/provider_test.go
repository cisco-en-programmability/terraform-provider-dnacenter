package dnacenter

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"dnacenter": testAccProvider,
	}
	tagResourceInit()
	siteResourceInit()
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if err := os.Getenv("DNAC_BASE_URL"); err == "" {
		t.Fatal("DNAC_BASE_URL must be set for acceptance tests")
	}
	if err := os.Getenv("DNAC_USERNAME"); err == "" {
		t.Fatal("DNAC_USERNAME must be set for acceptance tests")
	}
	if err := os.Getenv("DNAC_PASSWORD"); err == "" {
		t.Fatal("DNAC_PASSWORD must be set for acceptance tests")
	}
	if err := os.Getenv("DNAC_DEBUG"); err == "" {
		t.Fatal("DNAC_DEBUG must be set for acceptance tests")
	}
	if err := os.Getenv("DNAC_SSL_VERIFY"); err == "" {
		t.Fatal("DNAC_SSL_VERIFY must be set for acceptance tests")
	}
}
