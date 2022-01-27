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

resource "dna_tag" "tf_tag" {
  provider = dnacenter
  item {
    system_tag  = false
    description = "Terraform Tag"
    name        = "Terraform"
  }
}

data "dna_tag" "list" {
  provider = dnacenter
  sort_by  = "name"
  order    = "des"
}
output "tag_list_first_element" {
  value = data.dna_tag.list.items.0
}

resource "dna_tag" "data" {
  provider = dnacenter
  item {
    system_tag         = false
    description        = "New tag description for Terraform 012"
    name               = "Tag012"
    instance_tenant_id = "15cdc6c45a8405f00c80c6ba3"
    dynamic_rules {
      member_type = "networkdevice"
      rules {
        operation = "ILIKE"
        name      = "family"
        value     = "%Switches and Hubs%"
      }
    }
  }
}
output "tag_data" {
  value = dna_tag.data
}

data "dna_tag_member_type" "list" {
  provider = dnacenter
}
output "tag_member_list" {
  value = data.dna_tag_member_type.list
}


data "dna_tag_count" "amount" {
  provider   = dnacenter
  depends_on = [dna_tag.data]
  # name     = "Tag012"
  # system_tag = "False"
  # level = "0"
  # attribute_name = "1"
}
output "tag_amount" {
  value = data.dna_tag_count.amount
}


data "dna_tag_member" "response" {
  provider    = dnacenter
  depends_on  = [dna_tag.data]
  tag_id      = dna_tag.data.id
  member_type = "networkdevice"
  # member_type = data.dna_tag_member_type.list.items.0
}

output "tag_member_response" {
  value = data.dna_tag_member.response
}

data "dna_tag_member_count" "amount" {
  provider    = dnacenter
  depends_on  = [dna_tag.data]
  id          = dna_tag.data.id
  member_type = "networkdevice"
}
output "tag_member_count_amount" {
  value = data.dna_tag_member_count.amount.response
}
