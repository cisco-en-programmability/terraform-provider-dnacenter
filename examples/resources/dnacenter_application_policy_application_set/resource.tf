
resource "dnacenter_application_policy_application_set" "example" {
  provider = dnacenter
  parameters {

    default_business_relevance     = "string"
    id                             = "string"
    name                           = "string"
    namespace                      = "string"
    qualifier                      = "string"
    scalable_group_external_handle = "string"
    scalable_group_type            = "string"
    type                           = "string"
  }
}

output "dnacenter_application_policy_application_set_example" {
  value = dnacenter_application_policy_application_set.example
}