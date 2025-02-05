module github.com/scaleway/scaleway-cli/v2

go 1.22

toolchain go1.22.2

require (
	github.com/alecthomas/assert v1.0.0
	github.com/aws/aws-sdk-go-v2 v1.36.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.75.2
	github.com/buildpacks/pack v0.34.2
	github.com/c-bata/go-prompt v0.2.6
	github.com/charmbracelet/bubbletea v1.3.0
	github.com/chzyer/readline v1.5.1
	github.com/containerd/console v1.0.4
	github.com/dnaeon/go-vcr v1.2.0
	github.com/docker/docker v27.5.1+incompatible
	github.com/dustin/go-humanize v1.0.1
	github.com/fatih/color v1.18.0
	github.com/getsentry/sentry-go v0.31.1
	github.com/ghodss/yaml v1.0.0
	github.com/gorilla/websocket v1.5.3
	github.com/hashicorp/go-version v1.7.0
	github.com/karrick/tparse/v2 v2.8.2
	github.com/mattn/go-colorable v0.1.14
	github.com/mattn/go-isatty v0.0.20
	github.com/moby/buildkit v0.13.2
	github.com/opencontainers/go-digest v1.0.0
	github.com/scaleway/scaleway-sdk-go v1.0.0-beta.32
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	github.com/spf13/cobra v1.8.1
	github.com/spf13/pflag v1.0.6
	github.com/stretchr/testify v1.10.0
	golang.org/x/crypto v0.32.0
	golang.org/x/term v0.28.0
	golang.org/x/text v0.21.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	dario.cat/mergo v1.0.0 // indirect
	github.com/AdaLogics/go-fuzz-headers v0.0.0-20230811130428-ced1acdcaa24 // indirect
	github.com/Azure/azure-sdk-for-go v68.0.0+incompatible // indirect
	github.com/Azure/go-ansiterm v0.0.0-20230124172434-306776ec8161 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest v0.11.29 // indirect
	github.com/Azure/go-autorest/autorest/adal v0.9.23 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.12 // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.6 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/Microsoft/hcsshim v0.12.2 // indirect
	github.com/ProtonMail/go-crypto v1.1.3 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/alecthomas/colour v0.1.0 // indirect
	github.com/alecthomas/repr v0.2.0 // indirect
	github.com/apex/log v1.9.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.8 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.27.7 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.17.7 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.15.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.31 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.31 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.31 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecr v1.24.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecrpublic v1.21.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.5.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.18.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.20.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.23.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.28.4 // indirect
	github.com/aws/smithy-go v1.22.2 // indirect
	github.com/awslabs/amazon-ecr-credential-helper/ecr-login v0.0.0-20231213181459-b0fcec718dc6 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/buildpacks/imgutil v0.0.0-20240605145725-186f89b2d168 // indirect
	github.com/buildpacks/lifecycle v0.19.6 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/charmbracelet/lipgloss v1.0.0 // indirect
	github.com/charmbracelet/x/ansi v0.8.0 // indirect
	github.com/charmbracelet/x/term v0.2.1 // indirect
	github.com/chrismellard/docker-credential-acr-env v0.0.0-20230304212654-82a0ddb27589 // indirect
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/containerd/containerd v1.7.16 // indirect
	github.com/containerd/continuity v0.4.3 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/containerd/stargz-snapshotter/estargz v0.15.1 // indirect
	github.com/containerd/ttrpc v1.2.3 // indirect
	github.com/containerd/typeurl/v2 v2.1.1 // indirect
	github.com/cyphar/filepath-securejoin v0.2.5 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/distribution/reference v0.6.0 // indirect
	github.com/docker/cli v26.1.3+incompatible // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/docker-credential-helpers v0.8.0 // indirect
	github.com/docker/go-connections v0.5.0 // indirect
	github.com/docker/go-metrics v0.0.1 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/erikgeiser/coninput v0.0.0-20211004153227-1c3628e74d0f // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/gdamore/tcell/v2 v2.7.4 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-git/go-billy/v5 v5.6.0 // indirect
	github.com/go-git/go-git/v5 v5.13.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/go-containerregistry v0.19.1 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1 // indirect
	github.com/heroku/color v0.0.6 // indirect
	github.com/in-toto/in-toto-golang v0.9.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mattn/go-tty v0.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/ioprogress v0.0.0-20180201004757-6a23b12fa88e // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/moby/docker-image-spec v1.3.1 // indirect
	github.com/moby/locker v1.0.1 // indirect
	github.com/moby/patternmatcher v0.6.0 // indirect
	github.com/moby/sys/sequential v0.5.0 // indirect
	github.com/moby/sys/signal v0.7.0 // indirect
	github.com/moby/sys/user v0.1.0 // indirect
	github.com/moby/sys/userns v0.1.0 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/opencontainers/selinux v1.11.0 // indirect
	github.com/pjbgf/sha1cd v0.3.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pkg/term v1.2.0-beta.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.19.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.48.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rivo/tview v0.0.0-20230406072732-e22ce9588bb4 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/sabhiram/go-gitignore v0.0.0-20210923224102-525f6e181f06 // indirect
	github.com/secure-systems-lab/go-securesystemslib v0.8.0 // indirect
	github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3 // indirect
	github.com/shibumi/go-pathspec v1.3.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/skeema/knownhosts v1.3.0 // indirect
	github.com/tonistiigi/fsutil v0.0.0-20240424095704-91a3fc46842c // indirect
	github.com/tonistiigi/units v0.0.0-20180711220420-6950e57a87ea // indirect
	github.com/tonistiigi/vt100 v0.0.0-20230623042737-f9a4f7ef6531 // indirect
	github.com/vbatts/tar-split v0.11.5 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20190905194746-02993c407bfb // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.2.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.50.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.50.0 // indirect
	go.opentelemetry.io/otel v1.25.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.25.0 // indirect
	go.opentelemetry.io/otel/metric v1.25.0 // indirect
	go.opentelemetry.io/otel/sdk v1.25.0 // indirect
	go.opentelemetry.io/otel/trace v1.25.0 // indirect
	go.opentelemetry.io/proto/otlp v1.2.0 // indirect
	golang.org/x/mod v0.17.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	google.golang.org/genproto v0.0.0-20240401170217-c3f982113cda // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240401170217-c3f982113cda // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240401170217-c3f982113cda // indirect
	google.golang.org/grpc v1.63.2 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
