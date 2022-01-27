
resource "dnacenter_applications" "example" {
  provider = dnacenter
  parameters {

    application_set {

      id_ref = "string"
    }
    id = "string"
    indicative_network_identity {

      display_name = "string"
      id           = "string"
      lower_port   = 1
      ports        = "string"
      protocol     = "string"
      upper_port   = 1
    }
    name = "string"
    network_applications {

      app_protocol         = "string"
      application_sub_type = "string"
      application_type     = "string"
      category_id          = "string"
      display_name         = "string"
      dscp                 = "string"
      engine_id            = "string"
      help_string          = "string"
      id                   = "string"
      ignore_conflict      = "string"
      long_description     = "string"
      name                 = "string"
      popularity           = "string"
      rank                 = "string"
      server_name          = "string"
      traffic_class        = "string"
      url                  = "string"
    }
    network_identity {

      display_name = "string"
      id           = "string"
      lower_port   = "string"
      ports        = "string"
      protocol     = "string"
      upper_port   = "string"
    }
  }
}

output "dnacenter_applications_example" {
  value = dnacenter_applications.example
}