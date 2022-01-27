
terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
resource "dnacenter_applications" "example" {
  provider = dnacenter
  parameters {

    application_set {
      id_ref = "f6f5a7e8-18ea-4e58-ac00-3f5c6f97186b"
    }
    id   = "00a1043c-429c-4211-bc7c-cd172c54603d"
    name = "qmqp"
    network_applications {
      app_protocol         = "tcp/udp"
      application_sub_type = "NONE"
      application_type     = "DEFAULT"
      category_id          = "694c6e60-2774-49ad-9248-027292f059ef"
      display_name         = "17012"
      engine_id            = "3"
      help_string          = "qmqp"
      id                   = "96b555f2-bf7b-41aa-a569-c26d3d3f68d9"
      ignore_conflict      = "string"
      long_description     = "Quick Mail Queuing Protocol (QMQP) is a network protocol designed to share e-mail queues between several hosts. It is designed and implemented in qmail."
      name                 = "qmqp"
      popularity           = 777
      rank                 = 65535
      traffic_class        = "BULK_DATA"
    }
    network_identity {

      display_name = "19743"
      id           = "dfb61714-eea8-4e58-9d2a-69652b3f6121"
      lower_port   = 0
      ports        = "628"
      protocol     = "UDP"
      upper_port   = 0
    }

    /*
      indicative_network_identity= [
        {
          id= "b80b15a4-51d3-472c-88ee-1dcc24159138"
          display_name= "19746"
          lower_port= 0
          ports= "628"
          protocol= "UDP"
          upper_port= 0
        },
        {
          id= "defda173-3c89-4869-b6f9-ff32ec2acece"
          display_name= "19745"
          lower_port= 0
          ports= "628"
          protocol= "TCP"
          upper_port= 0
        }
      ]
      name= "qmqp"
      network_applications= [
        {
          id= "96b555f2-bf7b-41aa-a569-c26d3d3f68d9"
          app_protocol= "tcp/udp"
          application_subType= "NONE"
          application_type= "DEFAULT"
          category_id= "694c6e60-2774-49ad-9248-027292f059ef"
          display_name= "17012"
          engine_id= "3"
          help_string= "qmqp"
          long_description= "Quick Mail Queuing Protocol (QMQP) is a network protocol designed to share e-mail queues between several hosts. It is designed and implemented in qmail."
          name= "qmqp"
          popularity= 3
          rank= 65535
          selecto_id= "628"
          traffic_class= "BULK_DATA"
        }
      ]
      network_identity= [
        {
          id= "dfb61714-eea8-4e58-9d2a-69652b3f6121"
          display_name= "19743"
          lower_port= 0
          ports= "628"
          protocol= "UDP"
          upper_port= 0
        },
        {
          id= "ab406368-88b7-4755-bda6-6d305d87ca84"
          display_name= "19744"
          lower_port= 0
          ports= "628"
          protocol= "TCP"
          upper_port= 0
        }
      ]
      application_set= {
        id_ref= "f6f5a7e8-18ea-4e58-ac00-3f5c6f97186b"
      }
      */
  }
}

output "dnacenter_applications_example" {
  value = dnacenter_applications.example
}
