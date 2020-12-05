package dnacenter

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDNACenterApplicationCountDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDNACenterApplicationCountDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDNACenterApplicationCountDataSourceItemsRead,
				),
			},
		},
	})
}

func testAccDNACenterApplicationCountDataSourceItemsRead(state *terraform.State) error {
	tagQuery := state.RootModule().Resources["data.dna_applications_count.amount"]
	if tagQuery == nil {
		return fmt.Errorf("unable to find data.dna_applications_count.amount")
	}

	if tagQuery.Primary.ID == "" {
		return fmt.Errorf("No id set")
	}

	attr := tagQuery.Primary.Attributes["response"]
	numberOfApplicationCountQuery, err := strconv.Atoi(attr)
	if err != nil {
		return err
	}
	if numberOfApplicationCountQuery < 1 {
		return fmt.Errorf("unable to find any items")
	}
	return nil
}

const testAccDNACenterApplicationCountDataSourceConfig = `
data "dna_applications_count" "amount" {
  provider = dnacenter
}
`
