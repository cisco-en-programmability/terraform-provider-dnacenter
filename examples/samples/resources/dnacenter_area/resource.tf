
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.20-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_area" "example" {
  provider = dnacenter
  parameters {
    site {

      area {
        name        = "test_tf_area13"
        parent_name = "Global"
      }

    }
    type = "area"
    # site_id ="70c232d5-141e-4a03-918e-5a81acda6f38"
  }
}

output "dnacenter_area_example" {
  value = dnacenter_area.example
}
