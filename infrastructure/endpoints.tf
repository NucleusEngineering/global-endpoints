# Copyright 2023 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

data "google_cloud_run_locations" "default" {}

resource "google_cloud_run_service" "default" {
  for_each = toset(data.google_cloud_run_locations.default.locations)

  name     = "${local.prefix}-${each.value}"
  location = each.value

  template {
    spec {
      containers {
        image = "gcr.io/${local.project}/global-endpoint"
      }
    }
  }
}

resource "google_cloud_run_service_iam_member" "default" {
  for_each = toset(data.google_cloud_run_locations.default.locations)

  location = google_cloud_run_service.default[each.key].location
  project  = google_cloud_run_service.default[each.key].project
  service  = google_cloud_run_service.default[each.key].name
  role     = "roles/run.invoker"
  member   = "allUsers"
}

resource "google_compute_region_network_endpoint_group" "default" {
  for_each = toset(data.google_cloud_run_locations.default.locations)

  name                  = "${local.prefix}-neg-${each.key}"
  network_endpoint_type = "SERVERLESS"
  region                = google_cloud_run_service.default[each.key].location

  cloud_run {
    service = google_cloud_run_service.default[each.key].name
  }
}

module "lb-http" {
  source  = "GoogleCloudPlatform/lb-http/google//modules/serverless_negs"
  version = "~> 4.5"

  name    = "${local.prefix}-lb"
  project = local.project

  ssl                             = true
  managed_ssl_certificate_domains = ["global.stamer.nucleus-engineering.cloud."]
  https_redirect                  = true
  backends = {
    default = {
      description            = "Global Endpoints"
      enable_cdn             = false
      custom_request_headers = null

      log_config = {
        enable      = false
        sample_rate = null
      }

      groups = [
        for neg in google_compute_region_network_endpoint_group.default :
        {
          group = neg.id
        }
      ]

      iap_config = {
        enable               = false
        oauth2_client_id     = null
        oauth2_client_secret = null
      }
      security_policy = null
    }
  }
}

output "url" {
  value = "http://${module.lb-http.external_ip}"
}
