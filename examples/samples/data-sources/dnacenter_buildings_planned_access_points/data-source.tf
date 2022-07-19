terraform {
  required_providers {
    dnacenter = {
      version = "1.0.7-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_buildings_planned_access_points" "example" {
  provider    = dnacenter
  building_id = "2397da83-4e12-4d04-9bd3-a57b2ad91652"
  //limit = 1
  //offset = 1
  //radios = "false"
}

output "dnacenter_buildings_planned_access_points_example" {
  value = data.dnacenter_buildings_planned_access_points.example.items
}
