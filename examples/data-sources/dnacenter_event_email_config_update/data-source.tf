
data "dnacevent_email_config_update" "example" {
  provider        = dnac
  email_config_id = "string"
  from_email      = "string"
  item {

    # status_uri = ------
  }
  primary_smt_p_config {

    host_name = "string"
    password  = "******"
    port      = "string"
    user_name = "string"
  }
  secondary_smt_p_config {

    host_name = "string"
    password  = "******"
    port      = "string"
    user_name = "string"
  }
  subject  = "string"
  to_email = "string"
}