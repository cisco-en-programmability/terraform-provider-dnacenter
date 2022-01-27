terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_site_design_floormap" "example" {
  provider = dnacenter
  floor_id = "0f6661c2-ba34-4f4d-ae60-459cf293f689"
}

output "dnacenter_site_design_floormap_example" {
  value = data.dnacenter_site_design_floormap.example.item
}
