terraform {
  required_providers {
    dnacenter = {
      version = "1.0.17-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_tag_member" "example" {
  provider                = dnacenter
  id                      = "8f987b52-8698-45eb-9a80-bc5a02d4d972"
  level                   = "string"
  limit                   = 1
  member_association_type = "string"
  member_type             = "template"
  offset                  = 1
}

output "dnacenter_tag_member_example" {
  value = data.dnacenter_tag_member.example.items
}
