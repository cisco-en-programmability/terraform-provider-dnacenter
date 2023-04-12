
resource "dnacenter_sp_profile_delete_v2" "example" {
  provider = dnacenter
  parameters {

    sp_profile_name = "string"
  }
}

output "dnacenter_sp_profile_delete_v2_example" {
  value = dnacenter_sp_profile_delete_v2.example
}