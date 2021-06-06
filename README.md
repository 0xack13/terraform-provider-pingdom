# terraform-provider-pingdom

```
go build && cp terraform-provider-pingdom ~/.terraform.d/plugins/darwin_amd64
rm -fr .terraform*; terraform init && terraform plan
terraform apply
```

Run unit test:
```
go test -timeout 30s -run ^TestProvider$ terraform-provider-pingdom/pingdom
```
