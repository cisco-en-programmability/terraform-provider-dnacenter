terraform {
  required_providers {
    dnacenter = {
      version = "1.1.26-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_tag" "example" {
  provider = dnacenter
  parameters {

    description = "test33"
    dynamic_rules {

      member_type = "networkdevice"
      rules {

        # items     = "string"
        operation = "OR"
        items {
          name      = "hostname"
          operation = "ILIKE"
          value     = "%tf-rules%"
        }
        items {
          name      = "hostname"
          operation = "ILIKE"
          value     = "%-border-%"
        }
      }
      # rules {

      #   # items     = ["string"]
      #   name      = "string"
      #   operation = "string"
      #   value     = "string"
      #   values    = ["string"]
      # }
    }

    # instance_tenant_id = "62fe613eaa85d640eb70561e"
    //id                 = "0c3c64ea-69de-4686-8114-3f3dd7d97882" //Just necesesary for update...
    name = "api-test"
    //system_tag         = "false"
  }
}

output "dnacenter_tag_example" {
  value = dnacenter_tag.example
}
