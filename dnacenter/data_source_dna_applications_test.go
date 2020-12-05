package dnacenter

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDNACenterApplicationsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDNACenterApplicationsDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDNACenterApplicationsDataSourceItemsRead,
				),
			},
		},
	})
}

func testAccDNACenterApplicationsDataSourceItemsRead(state *terraform.State) error {
	tagQuery := state.RootModule().Resources["data.dna_applications.list"]
	if tagQuery == nil {
		return fmt.Errorf("unable to find data.dna_applications.list")
	}

	if tagQuery.Primary.ID == "" {
		return fmt.Errorf("No id set")
	}

	attr := tagQuery.Primary.Attributes["items.#"]
	numberOfApplicationsQuery, err := strconv.Atoi(attr)
	if err != nil {
		return err
	}
	if numberOfApplicationsQuery < 1 {
		return fmt.Errorf("unable to find any items")
	}
	return nil
}

const testAccDNACenterApplicationsDataSourceConfig = `
data "dna_applications" "list" {
  provider = dnacenter
  offset = 0
  limit = 4
}
`
