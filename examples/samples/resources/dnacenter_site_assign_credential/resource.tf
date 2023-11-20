
terraform {
  required_providers {
    dnacenter = {
      version = "1.1.27-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}
resource "dnacenter_site" "example" {
  provider = dnacenter
  parameters {
    site {
      /*
        area {
            name= "test_tf_area4"
            parent_name= "string"
        }
    
    */
      building {
        name        = "MyTestFinalq12"
        address     = "255 China Basin Street, San Francisco, California 94158, United States 2"
        parent_name = "Global"
        latitude    = 37.77178651716143
        longitude   = -122.39062051589885
      }
    }
    type = "building"
    # site_id ="70c232d5-141e-4a03-918e-5a81acda6f38"
  }
}

resource "dnacenter_site_assign_credential" "example" {
  provider   = dnacenter
  depends_on = [dnacenter_site.example]

  parameters {
    site_id = dnacenter_site.example.item.0.id
    cli_id  = "52b611b4-5ee4-4581-8a01-a2062cd2d3c1"
  }
}
