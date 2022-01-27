terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

resource "dna_site" "area" {
  provider = dnacenter
  item {
    type        = "area"
    name        = "Peru"
    parent_name = "Global"
  }
}
output "site_area" {
  value = dna_site.area
}

resource "dna_site" "building_1" {
  provider   = dnacenter
  depends_on = [dna_site.area]
  item {
    type        = "building"
    name        = "Ayacucho"
    parent_name = "Global/Peru"
    address     = "Ayacucho, Ayacucho, Peru"
    latitude    = -13.1604
    longitude   = -74.2257
  }
}
output "site_building_1" {
  value = dna_site.building_1
}

resource "dna_site" "building_2" {
  provider   = dnacenter
  depends_on = [dna_site.area]
  item {
    type        = "building"
    name        = "Miraflores"
    parent_name = "Global/Peru"
    address     = "Miraflores, Lima, Lima Province, Peru"
    latitude    = -12.1209
    longitude   = -77.0289
  }
}
output "site_building_2" {
  value = dna_site.building_2
}

resource "dna_site" "floor_1" {
  provider   = dnacenter
  depends_on = [dna_site.building_2]
  item {
    type        = "floor"
    name        = "Floor 1"
    parent_name = "Global/Peru/Miraflores"
    rf_model    = "Cubes And Walled Offices"
    height      = 100.1
    length      = 100.2
    width       = 100.1
  }
}

output "site_floor_1" {
  value = dna_site.floor_1
}
