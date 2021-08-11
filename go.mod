module github.com/GoogleCloudPlatform/terraform-validator

require (
	github.com/GoogleCloudPlatform/terraform-google-conversion v0.0.0-20210810210027-33c2d711b84d
	github.com/forseti-security/config-validator v0.0.0-20210621194145-08e4202b50d8
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/terraform-json v0.12.0
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.5.0
	github.com/hashicorp/terraform-provider-google/v3 v3.70.1-0.20210603175730-be2009058913
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.1.3
	github.com/stretchr/testify v1.7.0
	google.golang.org/api v0.46.0
	google.golang.org/genproto v0.0.0-20210503173045-b96a97608f20
)

replace github.com/GoogleCloudPlatform/terraform-google-conversion v0.0.0-20210810210027-33c2d711b84d => /Users/palani/ITP/Palani/Jobs/HCL-GCP-RLI/GOOG/code/terraform-google-conversion

go 1.14
