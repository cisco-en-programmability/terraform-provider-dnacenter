
resource "dnacenter_applications_v2" "example" {
  provider = dnacenter

  parameters {

    display_name = "string"
    id           = "string"
    indicative_network_identity {

      display_name = "string"
      id           = "string"
      ipv4_subnet  = ["string"]
      ipv6_subnet  = ["string"]
      lower_port   = 1.0
      ports        = "string"
      protocol     = "string"
      upper_port   = 1.0
    }
    instance_id      = 1
    instance_version = 1.0
    name             = "string"
    namespace        = "string"
    network_applications {

      app_protocol         = "string"
      application_sub_type = "string"
      application_type     = "string"
      category_id          = "string"
      display_name         = "string"
      dscp                 = "string"
      engine_id            = 1
      help_string          = "string"
      id                   = "string"
      ignore_conflict      = "false"
      long_description     = "string"
      name                 = "string"
      popularity           = 1.0
      rank                 = 1
      selector_id          = "string"
      server_name          = "string"
      traffic_class        = "string"
      type                 = "string"
      url                  = "string"
    }
    network_identity {

      display_name = "string"
      id           = "string"
      ipv4_subnet  = ["string"]
      ipv6_subnet  = ["string"]
      lower_port   = 1.0
      ports        = "string"
      protocol     = "string"
      upper_port   = 1.0
    }
    parent_scalable_group {

      id_ref = "string"
    }
    qualifier                      = "string"
    scalable_group_external_handle = "string"
    scalable_group_type            = "string"
    type                           = "string"
  }
}

output "dnacenter_applications_v2_example" {
  value = dnacenter_applications_v2.example
}