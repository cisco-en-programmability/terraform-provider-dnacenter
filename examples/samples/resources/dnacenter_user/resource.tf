terraform {
  required_providers {
    dnacenter = {
      version = "1.1.28-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_user" "example" {
  provider = dnacenter
  parameters {

    email      = "john.doe@example.com"
    first_name = "John"
    last_name  = "Doe"
    password   = "Secure_password"
    role_list  = ["63060e76f5bc0031ecc68c14"]
    user_id    = "johndoe123"
    username   = "johndoe"
  }
}

output "dnacenter_user_example" {
  sensitive = true
  value     = dnacenter_user.example
}