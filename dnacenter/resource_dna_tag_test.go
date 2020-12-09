package dnacenter

import (
	"fmt"
	"strconv"
	"testing"

	dnac "github.com/cisco-en-programmability/dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var tagCreateValues map[string]string
var tagUpdateValues map[string]string

func tagResourceInit() {
	tagCreateValues = map[string]string{
		"system_tag":         "false",
		"description":        "New tag description for Terraform 013",
		"name":               "Tag013",
		"instance_tenant_id": "5cdc6c45a8405f00c80c6ba3",
	}
	tagUpdateValues = map[string]string{
		"system_tag":         "false",
		"description":        "New tag description for Tag013 from Terraform test",
		"name":               "Tag013",
		"instance_tenant_id": "5cdc6c45a8405f00c80c6ba3",
	}
}

func TestAccDNACenterTagBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDNACenterTagDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDNACenterTagConfigBasic(tagCreateValues),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDNACenterTagExists("dna_tag.new", tagCreateValues),
				),
			},
			{
				Config: testAccCheckDNACenterTagConfigBasic(tagUpdateValues),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDNACenterTagExists("dna_tag.new", tagUpdateValues),
				),
			},
		},
	})
}

func testAccCheckDNACenterTagDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*dnac.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dna_tag" {
			continue
		}

		tagID := rs.Primary.ID

		_, _, err := c.Tag.DeleteTag(tagID)
		if err != nil {
			return err
		}
	}

	return nil
}

func testAccCheckDNACenterTagConfigBasic(values map[string]string) string {
	return fmt.Sprintf(`
	resource "dna_tag" "new" {
		provider = dnacenter
		item {
			system_tag = %s
			description = "%s"
			name = "%s"
			instance_tenant_id = "%s"
		}
	}
	`, values["system_tag"], values["description"], values["name"], values["instance_tenant_id"])
}

func testAccCheckDNACenterTagExists(n string, values map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No TagID set")
		}

		nItem, err := strconv.Atoi(rs.Primary.Attributes["item.#"])
		if err != nil {
			return err
		}
		if nItem < 1 {
			return fmt.Errorf("unable to find any item")
		}

		for key, value := range values {
			attribute := rs.Primary.Attributes[fmt.Sprintf("item.0.%s", key)]
			if fmt.Sprintf("%s", attribute) != value {
				return fmt.Errorf("attribute %s has different value %v %v", key, attribute, value)
			}
		}

		return nil
	}
}
