variable "aws_region" {
  default = "eu-central-1"
}

variable "admin_username" {
  default = "admin"
}

variable "admin_password" {
  default = "acctest"
}

#
# Note, you have to manually subscribe to this first in 
# the console I believe.
#
# https://aws.amazon.com/marketplace/pp/B00WH2LGM0
#
variable "asav_ami_id" {
  default = "ami-12f0d2f9"
}
