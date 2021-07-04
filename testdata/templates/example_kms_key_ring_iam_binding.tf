/**
 * Copyright 2019 Google LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "~> {{.Provider.version}}"
    }
  }
}

provider "google" {
  {{if .Provider.credentials }}credentials = "{{.Provider.credentials}}"{{end}}
}

data "google_kms_key_ring" "kms-keyring-dev" {
  name = "kms-keyring-dev"
  location = "global"
}

resource "google_kms_key_ring_iam_binding" "key_ring" {
  key_ring_id = data.google_kms_key_ring.kms-keyring-dev.id
  role        = "roles/cloudkms.cryptoKeyEncrypter"
  members = [
    "allUsers", "allAuthenticatedUsers"
  ]
}