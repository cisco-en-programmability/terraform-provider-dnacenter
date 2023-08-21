
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.11-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_building" "example" {
  provider = dnacenter
  parameters {
    site {
      building {
        name        = "MyTestFinalq"
        address     = "255 China Basin Street, San Francisco, California 94158, United States 1"
        parent_name = "Global"
        latitude    = 37.77178651716143
        longitude   = -122.39062051589885
      }
    }
    type = "building"
    # site_id ="70c232d5-141e-4a03-918e-5a81acda6f38"
  }
}

output "dnacenter_building_example" {
  value = dnacenter_building.example
}
