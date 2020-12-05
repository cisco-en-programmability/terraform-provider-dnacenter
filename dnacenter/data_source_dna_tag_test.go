package dnacenter

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDNACenterTagDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDNACenterTagDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.dna_tag.list", "order", "des"),
					resource.TestCheckResourceAttr("data.dna_tag.list", "sort_by", "name"),
					testAccDNACenterTagDataSourceItemsRead,
				),
			},
		},
	})
}

func testAccDNACenterTagDataSourceItemsRead(state *terraform.State) error {
	tagQuery := state.RootModule().Resources["data.dna_tag.list"]
	if tagQuery == nil {
		return fmt.Errorf("unable to find data.dna_tag.list")
	}

	if tagQuery.Primary.ID == "" {
		return fmt.Errorf("No id set")
	}

	attr := tagQuery.Primary.Attributes["items.#"]
	numberOfTagQuery, err := strconv.Atoi(attr)
	if err != nil {
		return err
	}
	if numberOfTagQuery < 1 {
		return fmt.Errorf("unable to find any items")
	}
	return nil
}

const testAccDNACenterTagDataSourceConfig = `
data "dna_tag" "list" {
  provider = dnacenter
  sort_by = "name"
  order = "des"
}
`
