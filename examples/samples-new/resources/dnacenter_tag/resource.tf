terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
resource "dnacenter_tag" "example" {
  provider = dnacenter
  parameters {

    description = "test3"
    /*dynamic_rules {

      member_type = "string"
      rules {

        items     = ["string"]
        name      = "string"
        operation = "string"
        value     = "string"
        values    = ["string"]
      }
    }
    
    instance_tenant_id = "string"*/
    //id                 = "0c3c64ea-69de-4686-8114-3f3dd7d97882" //Just necesesary for update...
    name               = "Day0Configuration2"
    //system_tag         = "false"
  }
}

output "dnacenter_tag_example" {
  value = dnacenter_tag.example
}