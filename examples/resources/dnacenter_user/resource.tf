
resource "dnacenter_user" "example" {
  provider = dnacenter
  parameters {

    email      = "string"
    first_name = "string"
    last_name  = "string"
    password   = "******"
    role_list  = ["string"]
    user_id    = "string"
    username   = "string"
  }
}

output "dnacenter_user_example" {
  value = dnacenter_user.example
}