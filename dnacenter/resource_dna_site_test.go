package dnacenter

import (
	"fmt"
	"strconv"
	"testing"

	dnac "dnacenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var siteAreaCreateValues map[string]string
var siteBuildingCreateValues map[string]string
var siteFloorCreateValues map[string]string

func siteResourceInit() {
	siteAreaCreateValues = map[string]string{
		"type":        "area",
		"name":        "Peru",
		"parent_name": "Global",
	}
	siteBuildingCreateValues = map[string]string{
		"type":        "building",
		"name":        "Miraflores",
		"parent_name": "Global/Peru",
		"address":     "Miraflores, Lima, Lima Province, Peru",
		"latitude":    "-12.1209",
		"longitude":   "-77.0289",
	}
	siteFloorCreateValues = map[string]string{
		"type":        "floor",
		"name":        "Floor 1",
		"parent_name": "Global/Peru/Miraflores",
		"rf_model":    "Cubes And Walled Offices",
		"height":      "100.1",
		"length":      "100.2",
		"width":       "100.1",
	}
}

func TestAccDNACenterSiteBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDNACenterSiteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDNACenterSiteConfigBasic("area", "", siteAreaCreateValues) +
					testAccCheckDNACenterSiteConfigBasic("building", "area", siteBuildingCreateValues) +
					testAccCheckDNACenterSiteConfigBasic("floor", "building", siteFloorCreateValues),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDNACenterSiteExists("dna_site.area", siteAreaCreateValues),
					testAccCheckDNACenterSiteExists("dna_site.building", siteBuildingCreateValues),
					testAccCheckDNACenterSiteExists("dna_site.floor", siteFloorCreateValues),
				),
			},
		},
	})
}

func testAccCheckDNACenterSiteDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*dnac.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dna_site" {
			continue
		}

		siteID := rs.Primary.Attributes["item.0.id"]

		_, _, err := c.Sites.DeleteSite(siteID)
		if err != nil {
			return err
		}
	}

	return nil
}

func testAccCheckDNACenterSiteConfigBasic(key string, dependsOn string, values map[string]string) string {
	if values["type"] == "area" {
		return fmt.Sprintf(`
		resource "dna_site" "%s" {
			provider = dnacenter
			item {
				type = "%s"
				name = "%s"
				parent_name = "%s"
			}
		}

		`, key, values["type"], values["name"], values["parent_name"])
	}
	if values["type"] == "building" {
		return fmt.Sprintf(`
		resource "dna_site" "%s" {
			provider = dnacenter
  		depends_on = [ dna_site.%s ]
			item {
				type = "%s"
				name = "%s"
				parent_name = "%s"
				address = "%s"
    		latitude = %s
    		longitude = %s
			}
		}

		`, key, dependsOn, values["type"], values["name"], values["parent_name"],
			values["address"], values["latitude"], values["longitude"])
	}
	if values["type"] == "floor" {
		// Ignore rf_model because value changes once created, invalidates test
		return fmt.Sprintf(`
		resource "dna_site" "%s" {
			provider = dnacenter
			depends_on = [ dna_site.%s ]
			lifecycle {
				ignore_changes = [ item.0.rf_model ]
			}
			item {
				type = "%s"
				name = "%s"
				parent_name = "%s"
				rf_model = "%s"
				height = %s
				length = %s
				width = %s
			}
		}

		`, key, dependsOn, values["type"], values["name"], values["parent_name"],
			values["rf_model"], values["height"], values["length"], values["width"])
	}
	return ""
}

func testAccCheckDNACenterSiteExists(n string, values map[string]string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No SiteID set")
		}

		nItem, err := strconv.Atoi(rs.Primary.Attributes["item.#"])
		if err != nil {
			return err
		}
		if nItem < 1 {
			return fmt.Errorf("unable to find any item")
		}

		for key, value := range values {
			if key == "rf_model" { // Ignore rf_model because value changes once created
				continue
			}
			attribute := rs.Primary.Attributes[fmt.Sprintf("item.0.%s", key)]
			if fmt.Sprintf("%s", attribute) != value {
				return fmt.Errorf("attribute %s has different value %v %v", key, attribute, value)
			}
		}

		return nil
	}
}
