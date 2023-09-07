# Globally Distributed Endpoints with Cloud Run and GCLB

This example shows how to deploy funcitonally-equivalent Cloud Run service in each GCP region around the globe and front them with global HTTPS forwarding rules so that global traffic destined for single static IPv4 Anycast IP is forwarded to the closest available, regional Cloud Run endpoint.

Use Terraform.
