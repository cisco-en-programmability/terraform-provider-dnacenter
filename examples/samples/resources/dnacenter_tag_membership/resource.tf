terraform {
  required_providers {
    dnacenter = {
      version = "1.1.32-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_tag_membership" "example" {
  provider = dnacenter
  parameters {
    #tag_id= "68f7ebd1-81fc-4099-944c-ac38ef6455a5"
    tag_id      = "f8f4a38f-e0e8-4fb2-91ca-d79ac4078f18"
    member_type = "template"
    member_id   = "8bf887f3-7cc2-4bd2-9407-df0fe6b7f63a"
    #member_id="fe2bd8b9-2cf0-4b73-b7dc-755ff0f26363"

  }
}

output "dnacenter_tag_membership_example" {
  value = dnacenter_tag_membership.example
}
