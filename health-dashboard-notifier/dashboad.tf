
provider "aws" {
  assume_role {
    role_arn = "health-dashboard-warning: arn:aws:iam::687447591496:role/omidsRoleForMonitoring"
  }
  region     = var.region
  access_key = var.AWS_ACCESS_KEY_ID     # aws access key
  secret_key = var.AWS_SECRET_ACCESS_KEY # aws secret key
}

module "health-dashboard-notifier" {
  source         = "../terraform-health-dashboard"
  env            = var.env
  lambdas_folder = ".../health-dashboard-notifier"
}