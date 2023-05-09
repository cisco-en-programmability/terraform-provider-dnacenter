
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.4-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_golden_image" "example" {
  provider = dnacenter
  parameters {
    image_id                 = "e7f80aaa-62d3-4390-a8ee-49bbfba036a3"
    site_id                  = "2397da83-4e12-4d04-9bd3-a57b2ad91652"
    device_role              = "ALL"
    device_family_identifier = "284389362"
  }
}

output "dnacenter_golden_image_example" {
  value = dnacenter_golden_image.example
}
