
resource "dnacenter_tag" "example" {
  provider = dnacenter

  parameters {

    description = "string"
    dynamic_rules {

      member_type = "string"
      rules {
        operation = "string"
        items {
          name      = "string"
          operation = "string"
          value     = "string"
        }
      }
    }
    id                 = "string"
    instance_tenant_id = "string"
    name               = "string"
    system_tag         = "false"
  }
}

output "dnacenter_tag_example" {
  value = dnacenter_tag.example
}