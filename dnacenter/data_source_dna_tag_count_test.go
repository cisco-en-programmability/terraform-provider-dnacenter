package dnacenter

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDNACenterTagCountDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDNACenterTagCountDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDNACenterTagCountDataSourceItemsRead,
				),
			},
		},
	})
}

func testAccDNACenterTagCountDataSourceItemsRead(state *terraform.State) error {
	tagQuery := state.RootModule().Resources["data.dna_tag_count.amount"]
	if tagQuery == nil {
		return fmt.Errorf("unable to find data.dna_tag_count.amount")
	}

	if tagQuery.Primary.ID == "" {
		return fmt.Errorf("No id set")
	}

	attr := tagQuery.Primary.Attributes["response"]
	numberOfTagCountQuery, err := strconv.Atoi(attr)
	if err != nil {
		return err
	}
	if numberOfTagCountQuery < 1 {
		return fmt.Errorf("unable to find any items")
	}
	return nil
}

const testAccDNACenterTagCountDataSourceConfig = `
data "dna_tag_count" "amount" {
  provider = dnacenter
}
`
