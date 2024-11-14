
data "dnacenter_sda_extranet_policies" "example" {
  provider             = dnacenter
  extranet_policy_name = "string"
  limit                = 1
  offset               = 1
}

output "dnacenter_sda_extranet_policies_example" {
  value = data.dnacenter_sda_extranet_policies.example.items
}
