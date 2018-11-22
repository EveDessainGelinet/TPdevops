provider "github" {
  token        = "f5dcbbcc3bc368baa4965e61836b244a0dda8ef3"
  organization = "HerokuPEE"
}

resource "github_repository" "RepoEmma" {
  private = false
  name        = "RepoEmma"
  description = "TP devops Heroku"
}

# Configure the Heroku provider
provider "heroku" {
  email   = "${var.heroku_email}"
  api_key = "${var.heroku_key}"
}

# Je déclare mes variables
variable "heroku_email" {
  type = "string"
}

variable "heroku_key" {
  type = "string"
}
