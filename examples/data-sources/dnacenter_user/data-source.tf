
data "dnacenter_user" "example" {
  provider      = dnacenter
  invoke_source = "string"
}

output "dnacenter_user_example" {
  value = data.dnacenter_user.example.item
}
