module github.com/tenderly/nitro

go 1.20

replace github.com/VictoriaMetrics/fastcache => ./fastcache

require (
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible
	github.com/Shopify/toxiproxy v2.1.4+incompatible
	github.com/alicebob/miniredis/v2 v2.21.0
	github.com/allegro/bigcache v1.2.1-0.20190218064605-e24eb225f156
	github.com/andybalholm/brotli v1.0.4
	github.com/aws/aws-sdk-go-v2 v1.21.2
	github.com/aws/aws-sdk-go-v2/config v1.18.45
	github.com/aws/aws-sdk-go-v2/credentials v1.13.43
	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.11.10
	github.com/aws/aws-sdk-go-v2/service/s3 v1.26.9
	github.com/cavaliergopher/grab/v3 v3.0.1
	github.com/cockroachdb/pebble v0.0.0-20230928194634-aa077af62593
	github.com/codeclysm/extract/v3 v3.0.2
	github.com/dgraph-io/badger/v3 v3.2103.2
	github.com/enescakir/emoji v1.0.0
	github.com/fatih/structtag v1.2.0
	github.com/gdamore/tcell/v2 v2.6.0
	github.com/google/go-cmp v0.5.9
	github.com/hashicorp/golang-lru/v2 v2.0.2
	github.com/holiman/uint256 v1.2.4
	github.com/ipfs/go-cid v0.4.1
	github.com/ipfs/go-libipfs v0.6.2 // indirect
	github.com/ipfs/interface-go-ipfs-core v0.11.0
	github.com/knadh/koanf v1.4.0
	github.com/libp2p/go-libp2p v0.27.8 // indirect
	github.com/multiformats/go-multiaddr v0.9.0 // indirect
	github.com/multiformats/go-multihash v0.2.1
	github.com/r3labs/diff/v3 v3.0.1
	github.com/rivo/tview v0.0.0-20230814110005-ccc2c8119703
	github.com/spf13/pflag v1.0.5
	github.com/wealdtech/go-merkletree v1.0.0
	golang.org/x/crypto v0.17.0
	golang.org/x/sys v0.15.0
	golang.org/x/term v0.15.0
	golang.org/x/tools v0.15.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

require github.com/gofrs/flock v0.8.1

require (
	github.com/DataDog/zstd v1.5.2 // indirect
	github.com/alicebob/gopher-json v0.0.0-20200520072559-a9ecdc9d1d3a // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.1 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.13 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.43 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.37 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.45 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.37 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.15.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.23.2 // indirect
	github.com/aws/smithy-go v1.15.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.10.0 // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.2 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cockroachdb/errors v1.9.1 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.3 // indirect
	github.com/consensys/bavard v0.1.13 // indirect
	github.com/consensys/gnark-crypto v0.12.1
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/crackcomm/go-gitignore v0.0.0-20170627025303-887ab5e44cc3 // indirect
	github.com/crate-crypto/go-kzg-4844 v0.7.0
	github.com/deckarep/golang-set/v2 v2.1.0
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.1.0 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/dlclark/regexp2 v1.7.0 // indirect
	github.com/dop251/goja v0.0.0-20230806174421-c933cf95e127
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/ethereum/c-kzg-4844 v0.4.0
	github.com/fsnotify/fsnotify v1.6.0
	github.com/gammazero/deque v0.2.1
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/getsentry/sentry-go v0.18.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0
	github.com/golang/glog v1.1.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3
	github.com/google/flatbuffers v1.12.1 // indirect
	github.com/google/pprof v0.0.0-20230405160723-4a4c7d95572b // indirect
	github.com/graph-gophers/graphql-go v1.3.0
	github.com/h2non/filetype v1.0.6 // indirect
	github.com/holiman/billy v0.0.0-20230718173358-1c7e68d277a7
	github.com/huin/goupnp v1.3.0
	github.com/ipfs/bbloom v0.0.4 // indirect
	github.com/ipfs/go-block-format v0.1.1 // indirect
	github.com/ipfs/go-blockservice v0.5.1 // indirect
	github.com/ipfs/go-datastore v0.6.0 // indirect
	github.com/ipfs/go-ipfs-blockstore v1.2.0 // indirect
	github.com/ipfs/go-ipfs-ds-help v1.1.0 // indirect
	github.com/ipfs/go-ipfs-exchange-interface v0.2.0 // indirect
	github.com/ipfs/go-ipfs-util v0.0.2 // indirect
	github.com/ipfs/go-ipld-cbor v0.0.6 // indirect
	github.com/ipfs/go-ipld-format v0.4.0 // indirect
	github.com/ipfs/go-ipld-legacy v0.1.1 // indirect
	github.com/ipfs/go-log v1.0.5 // indirect
	github.com/ipfs/go-log/v2 v2.5.1 // indirect
	github.com/ipfs/go-merkledag v0.9.0 // indirect
	github.com/ipfs/go-metrics-interface v0.0.1 // indirect
	github.com/ipfs/go-path v0.3.1 // indirect
	github.com/ipfs/go-verifcid v0.0.2 // indirect
	github.com/ipld/go-codec-dagpb v1.5.0 // indirect
	github.com/ipld/go-ipld-prime v0.19.0 // indirect
	github.com/jbenet/goprocess v0.1.4 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/juju/errors v0.0.0-20181118221551-089d3ea4e4d5 // indirect
	github.com/klauspost/compress v1.16.4 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/minio/sha256-simd v1.0.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/mmcloughlin/addchain v0.4.0 // indirect
	github.com/mr-tron/base58 v1.2.0 // indirect
	github.com/multiformats/go-base32 v0.1.0 // indirect
	github.com/multiformats/go-base36 v0.2.0 // indirect
	github.com/multiformats/go-multibase v0.2.0 // indirect
	github.com/multiformats/go-multicodec v0.8.1 // indirect
	github.com/multiformats/go-multistream v0.4.1 // indirect
	github.com/multiformats/go-varint v0.0.7 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/polydawn/refmt v0.89.0 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.42.0 // indirect
	github.com/prometheus/procfs v0.9.0 // indirect
	github.com/rhnvrm/simples3 v0.6.1 // indirect
	github.com/rivo/uniseg v0.4.3 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/supranational/blst v0.3.11
	github.com/urfave/cli/v2 v2.25.7
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/whyrusleeping/cbor-gen v0.0.0-20230126041949-52956bd4c9aa // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	github.com/yuin/gopher-lua v0.0.0-20210529063254-f4c35e4016d9 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/otel v1.7.0 // indirect
	go.opentelemetry.io/otel/trace v1.7.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/exp v0.0.0-20231110203233-9a3e6036ecaa
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sync v0.5.0
	golang.org/x/text v0.14.0
	golang.org/x/time v0.3.0
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	lukechampine.com/blake3 v1.1.7 // indirect
	rsc.io/tmplfunc v0.0.3 // indirect
)

require (
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gobwas/httphead v0.1.0
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.1.0
	github.com/gobwas/ws-examples v0.0.0-20190625122829-a9e8908d9484
	github.com/mailru/easygo v0.0.0-20190618140210-3c14a0dc985f
)

require (
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/VictoriaMetrics/fastcache v1.12.1
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5
	github.com/gballet/go-libpcsclite v0.0.0-20190607065134-2772fd86a8ff
	github.com/go-ole/go-ole v1.2.5 // indirect
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-stack/stack v1.8.1
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb
	github.com/google/uuid v1.3.1
	github.com/gorilla/websocket v1.5.0
	github.com/hashicorp/go-bexpr v0.1.10
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d // indirect
	github.com/holiman/bloomfilter/v2 v2.0.3
	github.com/jackpal/go-nat-pmp v1.0.2
	github.com/mattn/go-colorable v0.1.13
	github.com/mattn/go-isatty v0.0.18
	github.com/mattn/go-runewidth v0.0.14 // indirect
	github.com/mitchellh/mapstructure v1.4.2
	github.com/mitchellh/pointerstructure v1.2.0 // indirect
	github.com/olekukonko/tablewriter v0.0.5
	github.com/rs/cors v1.7.0
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible
	github.com/status-im/keycard-go v0.2.0
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
	github.com/tklauser/go-sysconf v0.3.12 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/storage/azblob v1.2.0
	github.com/aws/aws-sdk-go-v2/service/route53 v1.30.2
	github.com/cespare/cp v0.1.0
	github.com/cloudflare/cloudflare-go v0.79.0
	github.com/docker/docker v1.6.2
	github.com/ethereum/go-ethereum v1.13.10
	github.com/fatih/color v1.13.0
	github.com/fjl/gencodec v0.0.0-20230517082657-f9840df7b83e
	github.com/gballet/go-verkle v0.1.1-0.20231031103413-a67434b50f46
	github.com/google/gofuzz v1.2.0
	github.com/influxdata/influxdb-client-go/v2 v2.4.0
	github.com/influxdata/influxdb1-client v0.0.0-20220302092344-a9ab5670611c
	github.com/jedisct1/go-minisign v0.0.0-20230811132847-661be99b8267
	github.com/julienschmidt/httprouter v1.3.0
	github.com/karalabe/usb v0.0.3-0.20230711191512-61db3e06439c
	github.com/kylelemons/godebug v1.1.0
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	github.com/peterh/liner v1.1.1-0.20190123174540-a2c9a5303de7
	github.com/protolambda/bls12-381-util v0.0.0-20220416220906-d8552aa452c7
	github.com/stretchr/testify v1.8.4
	go.uber.org/automaxprocs v1.5.2
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.7.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.3.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.17.3 // indirect
	github.com/cockroachdb/tokenbucket v0.0.0-20230807174530-cc333fc44b06 // indirect
	github.com/crate-crypto/go-ipa v0.0.0-20231025140028-3c0104f4b233 // indirect
	github.com/deepmap/oapi-codegen v1.8.2 // indirect
	github.com/garslo/gogen v0.0.0-20170306192744-1d203ffc1f61 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20190430165422-3e4dfb77656c // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.4 // indirect
	github.com/influxdata/line-protocol v0.0.0-20210311194329-9aa0e372d097 // indirect
	github.com/kilic/bls12-381 v0.1.0 // indirect
	github.com/naoina/go-stringutil v0.1.0 // indirect
	github.com/onsi/gomega v1.27.4 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
