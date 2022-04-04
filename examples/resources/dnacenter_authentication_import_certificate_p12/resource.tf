
resource "dnacenter_authentication_import_certificate_p12" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    file_name     = "string"
    list_of_users = ["string"]
    p12_file_path = "string"
    p12_password  = "******"
    pk_password   = "******"
  }
}