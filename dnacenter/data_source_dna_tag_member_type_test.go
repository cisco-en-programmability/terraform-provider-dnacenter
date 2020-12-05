package dnacenter

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDNACenterTagMemberTypeDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccDNACenterTagMemberTypeDataSourceConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDNACenterTagMemberTypeDataSourceItemsRead,
				),
			},
		},
	})
}

func testAccDNACenterTagMemberTypeDataSourceItemsRead(state *terraform.State) error {
	tagQuery := state.RootModule().Resources["data.dna_tag_member_type.list"]
	if tagQuery == nil {
		return fmt.Errorf("unable to find data.dna_tag_member_type.list")
	}

	if tagQuery.Primary.ID == "" {
		return fmt.Errorf("No id set")
	}

	attr := tagQuery.Primary.Attributes["items.#"]
	numberOfTagMemberTypeQuery, err := strconv.Atoi(attr)
	if err != nil {
		return err
	}
	if numberOfTagMemberTypeQuery < 1 {
		return fmt.Errorf("unable to find any items")
	}
	return nil
}

const testAccDNACenterTagMemberTypeDataSourceConfig = `
data "dna_tag_member_type" "list" {
  provider = dnacenter
}
`
