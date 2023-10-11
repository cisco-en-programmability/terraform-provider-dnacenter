
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.22-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_floor" "example" {
  provider = dnacenter
  parameters {
    site {
      floor {
        floor_number = 5
        height       = 5
        length       = 100
        name         = "floor-1"
        parent_name  = "Global"
        rf_model     = "Cubes And Walled Offices"
        width        = 100
      }
    }
    type = "floor"
    # site_id ="70c232d5-141e-4a03-918e-5a81acda6f38"
  }
}

output "dnacenter_floor_example" {
  value = dnacenter_floor.example
}
