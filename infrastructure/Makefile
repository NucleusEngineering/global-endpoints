init:
	gsutil mb gs://stamer-experiments-global-endpoints-tf-state/ || exit 0
	terraform init

apply:
	terraform apply

destroy:
	terraform destroy

.PHONY: init infrastructure
