package config

// FlagConfig - файл конфигурации
const FlagConfig = "config-path"

// App config flags
//
//nolint:gosec // G101 : app params
const (
	// FlagAppNodeName sets the identifier name specifying the unique physical or logical server node.
	FlagAppNodeName string = "app-node-name"
	// FlagAppMaxListLimit defines the boundary constraints threshold controlling maximal slice array pagination allocation lookups.
	FlagAppMaxListLimit string = "app-max-list-limit"
	// FlagAppTokenIssuer defines the identity signature field populated inside marshaled JWT token properties.
	FlagAppTokenIssuer string = "app-token-issuer"
	// FlagAppCipherKey encapsulates the symmetric configuration key payload used to encrypt data block cells inside persistence layers.
	FlagAppCipherKey string = "app-cipher-key"
	// FlagAppDefShutdownTimeout defines the baseline context graceful timeout bounds applied before a hard OS signal exit.
	FlagAppDefShutdownTimeout string = "app-def-shutdown-timeout"
)

// Creds
//
//nolint:gosec // G101 : app params constants
const (
	// FlagCredsUsername details system service-to-machine background worker username parameters keys.
	FlagCredsUsername string = "svc-creds-username"
	// FlagCredsPassword establishes system background credentials verification passwords tokens.
	FlagCredsPassword string = "svc-creds-password"
	// FlagCredsScheduleInterval sets standard interval duration bounds guiding active service token rotation checks.
	FlagCredsScheduleInterval string = "svc-creds-schedule-interval"
	// FlagCredsErrorScheduleInterval triggers fallback sleep delays managed by background tickers upon operational errors loops.
	FlagCredsErrorScheduleInterval string = "svc-creds-error-schedule-interval"
)

// Data-audit-client
const (
	// FlagDataAuditClientBaseURL maps the upstream destination REST target endpoint servicing audit record streams.
	FlagDataAuditClientBaseURL string = "data-audit-client-base-url"
	// FlagDataAuditClientTimeout dictates execution budget limits assigned to individual remote logging REST calls.
	FlagDataAuditClientTimeout string = "data-audit-client-timeout"
	// FlagDataAuditClientWorkerCount manages internal concurrent processor pools allocating data-streaming workers.
	FlagDataAuditClientWorkerCount string = "data-audit-client-worker-count"
	// FlagDataAuditClientDataCapacity sets buffer constraint dimensions sizing target channels storage thresholds.
	FlagDataAuditClientDataCapacity string = "data-audit-client-data-capacity"
	// FlagDataAuditClientCompleteProcessing ensures current queue pipelines drain down fully upon active lifecycle shutdown events.
	FlagDataAuditClientCompleteProcessing string = "data-audit-client-complete-processing"
	// FlagDataAuditClientShutdownTimeout bounds block intervals dedicated to flushing remaining audit logs records to networks.
	FlagDataAuditClientShutdownTimeout string = "data-audit-client-shutdown-timeout"
)

// Auth config flags
//
//nolint:gosec // G101 : app params
const (
	// FlagAuthJWTSecret defines the raw validation credentials secret used to cryptographically evaluate generic symmetric signatures keys.
	FlagAuthJWTSecret string = "auth-jwt-secret"
	// FlagAuthJWTSigningMethod maps the explicit encryption variant used to calculate token integrity sums.
	FlagAuthJWTSigningMethod string = "auth-jwt-signing-method"
	// FlagAuthAccessTokenTTL dictates validity ranges bounding temporary short-lived authentication access assets.
	FlagAuthAccessTokenTTL string = "auth-access-token-ttl"
	// FlagAuthRefreshTokenTTL sets expiration limits governing persistent long-lived session renewal certificates properties.
	FlagAuthRefreshTokenTTL string = "auth-refresh-token-ttl"
	// FlagAuthRSAPrivateKeyPath sets filesystem locations referencing asymmetric private signing credentials assets.
	FlagAuthRSAPrivateKeyPath string = "auth-rsa-private-key-path"
	// FlagAuthMasterPasswordSalt establishes unique localized entropy multipliers added before hashing user profile credentials.
	FlagAuthMasterPasswordSalt string = "auth-master-password-salt"
)

// DB config flags
const (
	// FlagDBDSN specifies the downstream target relational СУБД connection string sequence parameter payload.
	FlagDBDSN string = "db-dsn"
	// FlagDBDriver explicitly specifies vendor runtime storage engines types (e.g. postgres, pgx).
	FlagDBDriver string = "db-driver"
	// FlagDBMaxOpenConns caps maximum parallel TCP connections entries kept active inside database pools.
	FlagDBMaxOpenConns string = "db-max-open-conns"
	// FlagDBMaxIdleConns caps maximum idle connection resources maintained inside shared pools boundaries.
	FlagDBMaxIdleConns string = "db-max-idle-conns"
	// FlagDBMaxIdleLifetime bounds connection lifecycles durations before executing active resource recyclings.
	FlagDBMaxIdleLifetime string = "db-max-idle-lifetime"
	// FlagDBConnTimeout specifies allocation connection block limits before raising network error states.
	FlagDBConnTimeout string = "db-conn-timeout"
)

// gRPC config flags
const (
	// FlagGRPCAddress binds binary protobuf streaming servers to precise host port targets interfaces.
	FlagGRPCAddress string = "grpc-address"
	// FlagGRPCMaxConnIdle sets max passive connection durability ranges before executing termination pipelines.
	FlagGRPCMaxConnIdle string = "grpc-max-conn-idle"
	// FlagGRPCMaxConnAge dictates maximal durability boundaries allocated to active single connections before triggering recycles.
	FlagGRPCMaxConnAge string = "grpc-max-conn-age"
	// FlagGRPCMaxConnAgeGrace establishes graceful cooling periods allowing in-flight requests processing loops to finish prior connection kills.
	FlagGRPCMaxConnAgeGrace string = "grpc-max-conn-age-grace"
	// FlagGRPCTimeout limits streaming I/O execution budgets across individual remote procedure invocations.
	FlagGRPCTimeout string = "grpc-timeout"
	// FlagGRPCKeepAliveTime sets interval clocks pinging active channels to evaluate baseline network connectivity status metrics.
	FlagGRPCKeepAliveTime string = "grpc-keep-alive-time"
	// FlagGRPCKeepAliveTimeout bounds response delay budgets before marking silent transport pipes as fully degraded.
	FlagGRPCKeepAliveTimeout string = "grpc-keep-alive-timeout"
	// FlagGRPCShutdownTimeout sets maximum block intervals allowed for server termination procedures to successfully close listeners.
	FlagGRPCShutdownTimeout string = "grpc-shutdown-timeout"
)

// http config flags
//
//nolint:gosec // G101 : app paramss
const (
	// FlagHTTPAddress maps standard Web server multiplexers to concrete host port networking layout paths.
	FlagHTTPAddress string = "http-address"
	// FlagHTTPReadTimeout bounds time windows assigned for receiving raw headers and body streams from a client network socket.
	FlagHTTPReadTimeout string = "http-read-timeout"
	// FlagHTTPWriteTimeout bounds total time budgets dedicated to fully flushing server response packets back to transport buffers.
	FlagHTTPWriteTimeout string = "http-write-timeout"
	// FlagHTTPIdleTimeout manages keep-alive connection thresholds before forcing socket eviction procedures.
	FlagHTTPIdleTimeout string = "http-idle-timeout"
	// FlagHTTPShutdownTimeout shutdown server timeout
	FlagHTTPShutdownTimeout string = "http-shutdown-timeout"
	// FlagHTTPPrivateKeyPath maps filesystem trajectories loading SSL/TLS private server keys.
	FlagHTTPPrivateKeyPath string = "http-private-key-path"
	// FlagHTTPCertificatePath points to structural validation certificates blocks enabling encrypted HTTPS connections.
	FlagHTTPCertificatePath string = "http-certificate-path"
	// FlagHTTPSecure toggles operational rules forcing server instances to run encrypted transport delivery loops.
	FlagHTTPSecure string = "http-secure"
	// FlagHTTPMaxRequestBodySize manages automated defensive payload size constraints protecting internal memory maps from OOM exploits.
	FlagHTTPMaxRequestBodySize string = "http-max-request-body-size"
)

// log config flags
const (
	// FlagLogLevel specifies execution diagnostic filtering levels (e.g. debug, info, warn, error).
	FlagLogLevel string = "log-level"
	// FlagLogFormat switches structured layout serializers between plain text patterns and rigid production JSON fields.
	FlagLogFormat string = "log-format"
)

// telemetry
const (
	// FlagTelemetryEnabled toggles trace tracking pipelines propagation.
	FlagTelemetryEnabled string = "telemetry-enabled"
	// FlagTelemetryServiceName maps identifier names detailing microservices taxonomies inside OpenTelemetry record spans.
	FlagTelemetryServiceName string = "telemetry-service-name"
	// FlagTelemetryExporterEndpoint points to the upstream distributed tracing backend (e.g. Jaeger gRPC/HTTP OTLP receiver).
	FlagTelemetryExporterEndpoint string = "telemetry-exporter-endpoint"
	// FlagTelemetrySampleRate enforces decimal evaluation probabilities governing head-based tracing span collection filters.
	FlagTelemetrySampleRate string = "telemetry-sample-rate"
	// FlagTelemetryTimeout sets maximal transport budget allowances before dropping unexported trace batches packages.
	FlagTelemetryTimeout string = "telemetry-timeout"
)

// EnvConfig - файл конфигурации
const EnvConfig string = "CONFIG_PATH"

// amqp connector
//
//nolint:gosec // G101 : app params
const (
	// FlagAMQPConnectorURL maps the baseline AMQP network address for the target cluster endpoint.
	FlagAMQPConnectorURL string = "amqp-connector-url"
	// FlagAMQPConnectorUsername defines user parameters for authenticating structural broker sessions.
	FlagAMQPConnectorUsername string = "amqp-connector-username"
	// FlagAMQPConnectorPassword establishes cryptographic secret tokens identifying connection channels.
	FlagAMQPConnectorPassword string = "amqp-connector-password"
	// FlagAMQPConnectorConnectTimeout sets network allocation block limits establishing connections.
	FlagAMQPConnectorConnectTimeout string = "amqp-connector-connect-timeout"
	// FlagAMQPConnectorWriteTimeout caps maximum block durations pushing byte messages onto sockets.
	FlagAMQPConnectorWriteTimeout string = "amqp-connector-write-timeout"
	// FlagAMQPConnectorIdleTimeout specifies durability boundaries for resting unassigned broker channels.
	FlagAMQPConnectorIdleTimeout string = "amqp-connector-idle-timeout"
	// FlagAMQPConnectorShutdownTimeout handles block intervals dedicated to closing server connection context maps gracefully.
	FlagAMQPConnectorShutdownTimeout string = "amqp-connector-shutdown-timeout"
)

const (
	// FlagLoginAttemptsSenderKind chooses specific broker types driving events tracking workflows (e.g. amqp, kafka).
	FlagLoginAttemptsSenderKind string = "login-attempts-sender-kind"
	// FlagLoginAttemptsSenderNotifyTimeout limits wait budgets assigned for async pubsub broadcasts sequences.
	FlagLoginAttemptsSenderNotifyTimeout string = "login-attempts-sender-notify-timeout"
)

// amqp login attempts sender
const (
	// FlagLoginAttemptsSenderAMQPConfigTargetName defines the routing exchange or queue target parameter name inside AMQP broker layouts.
	FlagLoginAttemptsSenderAMQPConfigTargetName string = "login-attempts-sender-amqp-target-name"
	// FlagLoginAttemptsSenderAMQPConfigConnectTimeout limits socket allocation budgets for publishing streams.
	FlagLoginAttemptsSenderAMQPConfigConnectTimeout string = "login-attempts-sender-amqp-connect-timeout"
	// FlagLoginAttemptsSenderAMQPConfigShutdownTimeout establishes safe timeout boundaries during event streaming closures.
	FlagLoginAttemptsSenderAMQPConfigShutdownTimeout string = "login-attempts-sender-amqp-shutdown-timeout"
	// FlagLoginAttemptsSenderAMQPConfigPublishMaxTryAttempts caps total retries loops executed upon network publishing degradations.
	FlagLoginAttemptsSenderAMQPConfigPublishMaxTryAttempts string = "login-attempts-sender-amqp-publish-max-try-attempts"
	// FlagLoginAttemptsSenderAMQPConfigPublishBaseRetryDelay defines initial backoff durations before retrying failed message publish routines.
	FlagLoginAttemptsSenderAMQPConfigPublishBaseRetryDelay string = "login-attempts-sender-amqp-publish-base-retry-delay"
	// FlagLoginAttemptsSenderAMQPConfigPublishMaxRetryDelay sets a strict cap bounding maximum potential backoff delays.
	FlagLoginAttemptsSenderAMQPConfigPublishMaxRetryDelay string = "login-attempts-sender-amqp-publish-max-retry-delay"
)

// kafka login attempts sender
const (
	// FlagLoginAttemptsSenderKafkaConfigBrokers identifies network cluster addresses servicing Kafka streaming nodes.
	FlagLoginAttemptsSenderKafkaConfigBrokers string = "login-attempts-sender-kafka-brokers"
	// FlagLoginAttemptsSenderKafkaConfigTargetName defines destination topic patterns routing login tracking payloads.
	FlagLoginAttemptsSenderKafkaConfigTargetName string = "login-attempts-sender-kafka-target-name"
	// FlagLoginAttemptsSenderKafkaConfigConnectTimeout limits socket initialization budgets under active Kafka dialers.
	FlagLoginAttemptsSenderKafkaConfigConnectTimeout string = "login-attempts-sender-kafka-connect-timeout"
	// FlagLoginAttemptsSenderKafkaConfigIdleTimeout bounds durability metrics for resting broker pipeline connections.
	FlagLoginAttemptsSenderKafkaConfigIdleTimeout string = "login-attempts-sender-kafka-idle-timeout"
	// FlagLoginAttemptsSenderKafkaConfigShutdownTimeout restricts termination block intervals allowed for internal producer cleanup.
	FlagLoginAttemptsSenderKafkaConfigShutdownTimeout string = "login-attempts-sender-kafka-shutdown-timeout"
	// FlagLoginAttemptsSenderKafkaConfigPublishMaxTryAttempts manages retry policy limitations under producer transport faults.
	FlagLoginAttemptsSenderKafkaConfigPublishMaxTryAttempts string = "login-attempts-sender-kafka-publish-max-try-attempts"
	// FlagLoginAttemptsSenderKafkaConfigPublishMaxRetryDelay enforces upper thresholds bounding maximum retry intervals under cluster degradation.
	FlagLoginAttemptsSenderKafkaConfigPublishMaxRetryDelay string = "login-attempts-sender-kafka-publish-max-retry-delay"
	// FlagLoginAttemptsSenderKafkaConfigPublishBaseRetryDelay sets initial fallback intervals for backoff routing rules.
	FlagLoginAttemptsSenderKafkaConfigPublishBaseRetryDelay string = "login-attempts-sender-kafka-publish-base-retry-delay"
	// FlagLoginAttemptsSenderKafkaConfigUsername handles SASL mechanism parameters identifying active streaming users.
	FlagLoginAttemptsSenderKafkaConfigUsername string = "login-attempts-sender-kafka-username"
	// FlagLoginAttemptsSenderKafkaConfigPassword secures SASL credential verification tokens.
	FlagLoginAttemptsSenderKafkaConfigPassword string = "login-attempts-sender-kafka-password"
	// FlagLoginAttemptsSenderKafkaConfigBatchSize handles quantitative constraints sizing accumulated batch records counts.
	FlagLoginAttemptsSenderKafkaConfigBatchSize string = "login-attempts-sender-kafka-batch-size"
	// FlagLoginAttemptsSenderKafkaConfigBatchBytes sets byte size allocation limits triggering automated batch flushes.
	FlagLoginAttemptsSenderKafkaConfigBatchBytes string = "login-attempts-sender-kafka-batch-bytes"
	// FlagLoginAttemptsSenderKafkaConfigBatchTimeout dictates temporal constraints forcing lazy batch queues onto networks.
	FlagLoginAttemptsSenderKafkaConfigBatchTimeout string = "login-attempts-sender-kafka-batch-timeout"
	// FlagLoginAttemptsSenderKafkaConfigWriteTimeout bounds explicit durations assigned for network block socket transmissions.
	FlagLoginAttemptsSenderKafkaConfigWriteTimeout string = "login-attempts-sender-kafka-write-timeout"
	// FlagLoginAttemptsSenderKafkaConfigRequiredAcks manages consistency metrics configuring required cluster brokers confirmations (e.g. 0, 1, all).
	FlagLoginAttemptsSenderKafkaConfigRequiredAcks string = "login-attempts-sender-kafka-required-acks"
)

// app
//
//nolint:gosec // G101 : app conf keys
const (
	keyAppEnv                string = "app.env"
	keyAppNodeName           string = "app.node_name"
	keyAppMaxListLimit       string = "app.max_list_limit"
	keyAppTokenIssuer        string = "app.token_issuer"
	keyAppCipherKey          string = "app.cipher_key"
	keyAppDefShutdownTimeout string = "app.def_shutdown_timeout"
)

// service credentials
//
//nolint:gosec // G101 : app conf keys
const (
	keySvcCredsUsername              string = "svc_creds.username"
	keySvcCredsPassword              string = "svc_creds.password"
	keySvcCredsScheduleInterval      string = "svc_creds.schedule_interval"
	keySvcCredsErrorScheduleInterval string = "svc_creds.error_schedule_interval"
)

// data audit client
const (
	keyDataAuditClientBaseURL            string = "data_audit_client.base_url"
	keyDataAuditClientTimeout            string = "data_audit_client.timeout"
	keyDataAuditClientWorkerCount        string = "data_audit_client.worker_count"
	keyDataAuditClientDataCapacity       string = "data_audit_client.data_capacity"
	keyDataAuditClientCompleteProcessing string = "data_audit_client.complete_processing"
	keyDataAuditClientShutdownTimeout    string = "data_audit_client.shutdown_timeout"
)

// amqp connector
//
//nolint:gosec // G101 : app conf keys
const (
	// keyAMQPConnectorURL tracks the nested viper schema dot-key location for the baseline AMQP cluster network address string.
	keyAMQPConnectorURL string = "amqp_connector.url"
	// keyAMQPConnectorUsername maps nested path configuration nodes governing authorization user credentials blocks.
	keyAMQPConnectorUsername string = "amqp_connector.username"
	// keyAMQPConnectorPassword protects nested trajectories parsing access passwords.
	keyAMQPConnectorPassword string = "amqp_connector.password"
	// keyAMQPConnectorConnectTimeout points to internal property structures managing connection budget limits.
	keyAMQPConnectorConnectTimeout string = "amqp_connector.connect_timeout"
	// keyAMQPConnectorWriteTimeout targets configurations specifying maximal write block timeouts parameters.
	keyAMQPConnectorWriteTimeout string = "amqp_connector.write_timeout"
	// keyAMQPConnectorIdleTimeout dictates internal trajectories formatting unassigned channels durability lifetimes.
	keyAMQPConnectorIdleTimeout string = "amqp_connector.idle_timeout"
	// keyAMQPConnectorShutdownTimeout binds execution properties framing termination intervals budgets.
	keyAMQPConnectorShutdownTimeout string = "amqp_connector.shutdown_timeout"
)

// login attempts sender
const (
	// keyLoginAttemptsSenderKind chooses specific broker types driving events tracking workflows (e.g. amqp, kafka).
	keyLoginAttemptsSenderKind string = "login_attempts_sender.sender_kind"
	// keyLoginAttemptsSenderNotifyTimeout limits wait budgets assigned for async pubsub broadcasts sequences.
	keyLoginAttemptsSenderNotifyTimeout string = "login_attempts_sender.notify_timeout"
)

// amqp login attempts sender
const (
	// keyLoginAttemptsSenderAMQPConfigTargetName targets configuration nodes establishing destination queues or exchanges.
	keyLoginAttemptsSenderAMQPConfigTargetName string = "login_attempts_sender.amqp_config.target_name"
	// keyLoginAttemptsSenderAMQPConfigConnectTimeout dictates socket initialization budgets nested under parameters keys.
	keyLoginAttemptsSenderAMQPConfigConnectTimeout string = "login_attempts_sender.amqp_config.connect_timeout"
	// keyLoginAttemptsSenderAMQPConfigShutdownTimeout manages contextual timeout structures applied across streaming closures operations.
	keyLoginAttemptsSenderAMQPConfigShutdownTimeout string = "login_attempts_sender.amqp_config.shutdown_timeout"
	// keyLoginAttemptsSenderAMQPConfigPublishMaxTryAttempts restricts total retries counters upon network pipeline errors loops.
	keyLoginAttemptsSenderAMQPConfigPublishMaxTryAttempts string = "login_attempts_sender.amqp_config.publish_max_try_attempts"
	// keyLoginAttemptsSenderAMQPConfigPublishBaseRetryDelay maps property paths specifying backoff multiplier initial values duration.
	keyLoginAttemptsSenderAMQPConfigPublishBaseRetryDelay string = "login_attempts_sender.amqp_config.publish_base_retry_delay"
	// keyLoginAttemptsSenderAMQPConfigPublishMaxRetryDelay caps maximal delay limits managing cascading backoff parameters structures.
	keyLoginAttemptsSenderAMQPConfigPublishMaxRetryDelay string = "login_attempts_sender.amqp_config.publish_max_retry_delay"
)

// kafka login attempts
const (
	// keyLoginAttemptsSenderKafkaConfigBrokers identifies compound array configurations trajectories pointing to target Kafka network nodes addresses.
	keyLoginAttemptsSenderKafkaConfigBrokers string = "login_attempts_sender.kafka_config.brokers"
	// keyLoginAttemptsSenderKafkaConfigTargetName defines nested topic destination properties.
	keyLoginAttemptsSenderKafkaConfigTargetName string = "login_attempts_sender.kafka_config.target_name"
	// keyLoginAttemptsSenderKafkaConfigConnectTimeout manages dialer socket initialization duration parameters keys.
	keyLoginAttemptsSenderKafkaConfigConnectTimeout string = "login_attempts_sender.kafka_config.connect_timeout"
	// keyLoginAttemptsSenderKafkaConfigIdleTimeout points to structures specifying passive channels connection lifetime guidelines.
	keyLoginAttemptsSenderKafkaConfigIdleTimeout string = "login_attempts_sender.kafka_config.idle_timeout"
	// keyLoginAttemptsSenderKafkaConfigShutdownTimeout constraints cleanup execution durations allowed during stream engine graceful shutdowns.
	keyLoginAttemptsSenderKafkaConfigShutdownTimeout string = "login_attempts_sender.kafka_config.shutdown_timeout"
	// keyLoginAttemptsSenderKafkaConfigPublishMaxTryAttempts targets retry configuration metrics driven under producer faults triggers.
	keyLoginAttemptsSenderKafkaConfigPublishMaxTryAttempts string = "login_attempts_sender.kafka_config.publish_max_try_attempts"
	// keyLoginAttemptsSenderKafkaConfigPublishMaxRetryDelay specifies ceiling boundaries for backoff delay parameters routines.
	keyLoginAttemptsSenderKafkaConfigPublishMaxRetryDelay string = "login_attempts_sender.kafka_config.publish_max_retry_delay"
	// keyLoginAttemptsSenderKafkaConfigPublishBaseRetryDelay sets baseline starting backoff delays duration indices.
	keyLoginAttemptsSenderKafkaConfigPublishBaseRetryDelay string = "login_attempts_sender.kafka_config.publish_base_retry_delay"
	// keyLoginAttemptsSenderKafkaConfigUsername resolves SASL auth parameters mapping active users names.
	keyLoginAttemptsSenderKafkaConfigUsername string = "login_attempts_sender.kafka_config.username"
	// keyLoginAttemptsSenderKafkaConfigPassword secures nested paths handling SASL verification secret tokens.
	keyLoginAttemptsSenderKafkaConfigPassword string = "login_attempts_sender.kafka_config.password"
	// keyLoginAttemptsSenderKafkaConfigBatchSize binds internal arrays constraints pacing maximum batch records numbers.
	keyLoginAttemptsSenderKafkaConfigBatchSize string = "login_attempts_sender.kafka_config.batch_size"
	// keyLoginAttemptsSenderKafkaConfigBatchBytes paths configurations checking buffer size limits in bytes before forcing automated network flushes.
	keyLoginAttemptsSenderKafkaConfigBatchBytes string = "login_attempts_sender.kafka_config.batch_bytes"
	// keyLoginAttemptsSenderKafkaConfigBatchTimeout manages lazy queue timer triggers forcing buffer flushes to downstream networks interfaces.
	keyLoginAttemptsSenderKafkaConfigBatchTimeout string = "login_attempts_sender.kafka_config.batch_timeout"
	// keyLoginAttemptsSenderKafkaConfigWriteTimeout dictates socket block durability limits managing individual record packages writes operations.
	keyLoginAttemptsSenderKafkaConfigWriteTimeout string = "login_attempts_sender.kafka_config.write_timeout"
	// keyLoginAttemptsSenderKafkaConfigRequiredAcks structures consensus verification metrics configuring brokers replication confirmations criteria numbers.
	keyLoginAttemptsSenderKafkaConfigRequiredAcks string = "login_attempts_sender.kafka_config.required_acks"
)
