provider "aws" {}

data "aws_availability_zones" "available" {
  # c4.large instance type not available in these AZ
  blacklisted_zone_ids = ["usw2-az4"]
  state                = "available"
}
