module go.mondoo.com/cnquery/v9/providers/aws

replace go.mondoo.com/cnquery/v9 => ../..

go 1.21

toolchain go1.21.3

require (
	github.com/aws/aws-sdk-go v1.46.6
	github.com/aws/aws-sdk-go-v2 v1.21.2
	github.com/aws/aws-sdk-go-v2/config v1.19.1
	github.com/aws/aws-sdk-go-v2/credentials v1.13.43
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.13
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v1.21.2
	github.com/aws/aws-sdk-go-v2/service/acm v1.19.2
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.18.2
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v1.22.7
	github.com/aws/aws-sdk-go-v2/service/autoscaling v1.32.0
	github.com/aws/aws-sdk-go-v2/service/backup v1.25.2
	github.com/aws/aws-sdk-go-v2/service/cloudfront v1.28.7
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.29.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.28.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.24.2
	github.com/aws/aws-sdk-go-v2/service/codebuild v1.22.2
	github.com/aws/aws-sdk-go-v2/service/configservice v1.38.0
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v1.31.2
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.23.0
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.128.0
	github.com/aws/aws-sdk-go-v2/service/ecr v1.20.2
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.18.2
	github.com/aws/aws-sdk-go-v2/service/ecs v1.30.4
	github.com/aws/aws-sdk-go-v2/service/efs v1.21.9
	github.com/aws/aws-sdk-go-v2/service/eks v1.30.0
	github.com/aws/aws-sdk-go-v2/service/elasticache v1.30.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.18.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.22.0
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v1.21.0
	github.com/aws/aws-sdk-go-v2/service/emr v1.30.0
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.30.0
	github.com/aws/aws-sdk-go-v2/service/iam v1.24.0
	github.com/aws/aws-sdk-go-v2/service/kms v1.24.7
	github.com/aws/aws-sdk-go-v2/service/lambda v1.41.0
	github.com/aws/aws-sdk-go-v2/service/organizations v1.20.8
	github.com/aws/aws-sdk-go-v2/service/rds v1.58.0
	github.com/aws/aws-sdk-go-v2/service/redshift v1.33.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.40.2
	github.com/aws/aws-sdk-go-v2/service/s3control v1.34.0
	github.com/aws/aws-sdk-go-v2/service/sagemaker v1.113.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.21.6
	github.com/aws/aws-sdk-go-v2/service/securityhub v1.38.0
	github.com/aws/aws-sdk-go-v2/service/sns v1.23.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.40.0
	github.com/aws/aws-sdk-go-v2/service/sts v1.23.2
	github.com/aws/smithy-go v1.15.0
	github.com/cockroachdb/errors v1.11.1
	github.com/rs/zerolog v1.31.0
	github.com/spf13/afero v1.10.0
	github.com/stretchr/testify v1.8.4
	go.mondoo.com/cnquery/v9 v9.4.0
	k8s.io/client-go v0.28.3
)

require (
	cloud.google.com/go v0.110.9 // indirect
	cloud.google.com/go/compute v1.23.2 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v1.1.4 // indirect
	cloud.google.com/go/kms v1.15.4 // indirect
	cloud.google.com/go/secretmanager v1.11.3 // indirect
	cloud.google.com/go/storage v1.33.0 // indirect
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/99designs/keyring v1.2.2 // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/GoogleCloudPlatform/berglas v1.0.3 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.14 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.43 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.37 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.45 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.1.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect v1.17.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.38 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.37 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.37 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.15.6 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.15.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.17.3 // indirect
	github.com/awslabs/amazon-ecr-credential-helper/ecr-login v0.0.0-20231024185945-8841054dbdb8 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.14.3 // indirect
	github.com/danieljoos/wincred v1.2.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/distribution/reference v0.5.0 // indirect
	github.com/docker/cli v24.0.7+incompatible // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/docker v24.0.7+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.8.0 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/dvsekhvalnov/jose2go v1.5.0 // indirect
	github.com/fatih/color v1.15.0 // indirect
	github.com/getsentry/sentry-go v0.25.0 // indirect
	github.com/go-jose/go-jose/v3 v3.0.0 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-containerregistry v0.16.1 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.4 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.8 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.5 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/vault/api v1.10.0 // indirect
	github.com/hashicorp/yamux v0.1.1 // indirect
	github.com/hnakamur/go-scp v1.0.2 // indirect
	github.com/hokaccha/go-prettyjson v0.0.0-20211117102719-0474bc63780f // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/compress v1.17.2 // indirect
	github.com/kr/fs v0.1.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/sftp v1.13.6 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/segmentio/fasthash v1.0.3 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	github.com/sethvargo/go-password v0.2.0 // indirect
	github.com/sethvargo/go-retry v0.2.4 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/vbatts/tar-split v0.11.5 // indirect
	go.mondoo.com/ranger-rpc v0.5.2 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/mod v0.13.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/oauth2 v0.13.0 // indirect
	golang.org/x/sync v0.4.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/term v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.14.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/api v0.148.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20231016165738-49dd2c1f3d0b // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231016165738-49dd2c1f3d0b // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231016165738-49dd2c1f3d0b // indirect
	google.golang.org/grpc v1.59.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect
	moul.io/http2curl v1.0.0 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
