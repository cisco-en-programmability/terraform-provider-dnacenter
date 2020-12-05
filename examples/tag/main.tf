terraform {
  required_providers {
    dnacenter = {
      versions = ["0.2"]
      source   = "hashicorp.com/edu/dnacenter"
    }
  }
}

provider "dnacenter" {
}

# > Tag Query
data "dna_tag" "found" {
  provider = dnacenter
  sort_by  = "name"
  order    = "des"
}
output "tag_found_list" {
  value = data.dna_tag.found
}

# > Tag
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

# > Tag Member Types
data "dna_tag_member_type" "list" {
  provider = dnacenter
}
output "tag_member_list" {
  value = data.dna_tag_member_type.list
}


# > Tag Count
data "dna_tag_count" "amount" {
  provider = dnacenter
  name     = "Tag012"
  # system_tag = "False"
  # level = "0"
  # attribute_name = "1"
}
output "tag_amount" {
  value = data.dna_tag_count.amount
}

