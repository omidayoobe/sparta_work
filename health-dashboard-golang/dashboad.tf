
provider "aws" {
  region = var.region
  access_key = var.AWS_ACCESS_KEY_ID   # aws access key
 secret_key = var.AWS_SECRET_ACCESS_KEY # aws secret key
}

module "health-dashboard-notifier" {
  source = "https://github.com/omidayoobe/sparta_work/tree/master/terraform-health-dashboard"

}