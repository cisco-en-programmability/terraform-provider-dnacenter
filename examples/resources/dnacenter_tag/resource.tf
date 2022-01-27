
resource "dnacenter_tag" "example" {
  provider = dnacenter
  parameters {

    description = "string"
    dynamic_rules {

      member_type = "string"
      rules {

        items     = ["string"]
        name      = "string"
        operation = "string"
        value     = "string"
        values    = ["string"]
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