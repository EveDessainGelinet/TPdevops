# Create Heroku apps for staging and production
resource "heroku_app" "staging" {
  name   = "app-pee-s"
  region = "eu"
}

resource "heroku_app" "production" {
  name   = "app-pee-p"
  region = "eu"
}

# Create a Heroku pipeline
resource "heroku_pipeline" "app-pee" {
  name = "app-pee"
}

# Couple apps to different pipeline stages
resource "heroku_pipeline_coupling" "staging" {
  app      = "${heroku_app.staging.name}"
  pipeline = "${heroku_pipeline.app-pee.id}"
  stage    = "staging"
}

resource "heroku_pipeline_coupling" "production" {
  app      = "${heroku_app.production.name}"
  pipeline = "${heroku_pipeline.app-pee.id}"
  stage    = "production"
}


