
resource "dnacenter_endpoint_analytics_profiling_rules" "example" {
  provider = dnacenter
  parameters {

    cluster_id = "string"
    condition_groups {

      condition {

        attribute            = "string"
        attribute_dictionary = "string"
        operator             = "string"
        value                = "string"
      }
      condition_group = ["string"]
      operator        = "string"
      type            = "string"
    }
    is_deleted       = "false"
    last_modified_by = "string"
    last_modified_on = 1
    plugin_id        = "string"
    rejected         = "false"
    result {

      device_type           = ["string"]
      hardware_manufacturer = ["string"]
      hardware_model        = ["string"]
      operating_system      = ["string"]
    }
    rule_id         = "string"
    rule_name       = "string"
    rule_priority   = 1
    rule_type       = "string"
    rule_version    = 1
    source_priority = 1
    used_attributes = ["string"]
  }
}

output "dnacenter_endpoint_analytics_profiling_rules_example" {
  value = dnacenter_endpoint_analytics_profiling_rules.example
}