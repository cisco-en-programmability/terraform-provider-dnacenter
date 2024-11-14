
resource "dnacenter_event_email_config_create" "example" {
  provider = dnacenter
  parameters {

    email_config_id = "string"
    from_email      = "string"
    primary_smt_p_config {

      host_name = "string"
      password  = "******"
      port      = "string"
      smtp_type = "string"
      user_name = "string"
    }
    secondary_smt_p_config {

      host_name = "string"
      password  = "******"
      port      = "string"
      smtp_type = "string"
      user_name = "string"
    }
    subject  = "string"
    to_email = "string"
  }
}

output "dnacenter_event_email_config_create_example" {
  value = dnacenter_event_email_config_create.example
}