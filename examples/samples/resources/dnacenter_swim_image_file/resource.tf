
terraform {
  required_providers {
    dnacenter = {
      version = "1.0.16-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_swim_image_file" "example" {
  provider = dnacenter
  parameters {
    file_path = "$PATH/terraform-provider-dnacenter/examples/samples/resources/dnacenter_swim_image_file/testIMG1.zip"
    file_name = "testIMG1.zip"
  }
}

output "dnacenter_swim_image_file_example" {
  value = dnacenter_swim_image_file.example
}
