package config

// FlagConfig - файл конфигурации
const FlagConfig = "config-path"

// EnvConfig - файл конфигурации
const EnvConfig string = "CONFIG_PATH"

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

// login attempts sender amqp connector
//
//nolint:gosec // G101 : app params
const (
	// FlagLoginAttemptsSenderAMQPConnectorURL maps the baseline AMQP network address for the target cluster endpoint.
	FlagLoginAttemptsSenderAMQPConnectorURL string = "login-attempts-sender-amqp-connector-url"
	// FlagLoginAttemptsSenderAMQPConnectorUsername defines user parameters for authenticating structural broker sessions.
	FlagLoginAttemptsSenderAMQPConnectorUsername string = "login-attempts-sender-amqp-connector-username"
	// FlagLoginAttemptsSenderAMQPConnectorPassword establishes cryptographic secret tokens identifying connection channels.
	FlagLoginAttemptsSenderAMQPConnectorPassword string = "login-attempts-sender-amqp-connector-password"
	// FlagLoginAttemptsSenderAMQPConnectorConnectTimeout sets network allocation block limits establishing connections.
	FlagLoginAttemptsSenderAMQPConnectorConnectTimeout string = "login-attempts-sender-amqp-connector-connect-timeout"
	// FlagLoginAttemptsSenderAMQPConnectorWriteTimeout caps maximum block durations pushing byte messages onto sockets.
	FlagLoginAttemptsSenderAMQPConnectorWriteTimeout string = "login-attempts-sender-amqp-connector-write-timeout"
	// FlagLoginAttemptsSenderAMQPConnectorIdleTimeout specifies durability boundaries for resting unassigned broker channels.
	FlagLoginAttemptsSenderAMQPConnectorIdleTimeout string = "login-attempts-sender-amqp-connector-idle-timeout"
	// FlagLoginAttemptsSenderAMQPConnectorShutdownTimeout handles block intervals dedicated to closing server connection context maps gracefully.
	FlagLoginAttemptsSenderAMQPConnectorShutdownTimeout string = "login-attempts-sender-amqp-connector-shutdown-timeout"
)

// login attempts sender
const (
	// FlagLoginAttemptsSenderKind chooses specific broker types driving events tracking workflows (e.g. amqp, kafka).
	FlagLoginAttemptsSenderKind string = "login-attempts-sender-kind"
	// FlagLoginAttemptsSenderNotifyTimeout limits wait budgets assigned for async pubsub broadcasts sequences.
	FlagLoginAttemptsSenderNotifyTimeout string = "login-attempts-sender-notify-timeout"
)

// login attempts sender amqp
const (
	// FlagLoginAttemptsSenderAMQPTargetName defines the routing exchange or queue target parameter name inside AMQP broker layouts.
	FlagLoginAttemptsSenderAMQPTargetName string = "login-attempts-sender-amqp-target-name"
	// FlagLoginAttemptsSenderAMQPConnectTimeout limits socket allocation budgets for publishing streams.
	FlagLoginAttemptsSenderAMQPConnectTimeout string = "login-attempts-sender-amqp-connect-timeout"
	// FlagLoginAttemptsSenderAMQPShutdownTimeout establishes safe timeout boundaries during event streaming closures.
	FlagLoginAttemptsSenderAMQPShutdownTimeout string = "login-attempts-sender-amqp-shutdown-timeout"
	// FlagLoginAttemptsSenderAMQPPublishMaxTryAttempts caps total retries loops executed upon network publishing degradations.
	FlagLoginAttemptsSenderAMQPPublishMaxTryAttempts string = "login-attempts-sender-amqp-publish-max-try-attempts"
	// FlagLoginAttemptsSenderAMQPPublishBaseRetryDelay defines initial backoff durations before retrying failed message publish routines.
	FlagLoginAttemptsSenderAMQPPublishBaseRetryDelay string = "login-attempts-sender-amqp-publish-base-retry-delay"
	// FlagLoginAttemptsSenderAMQPPublishMaxRetryDelay sets a strict cap bounding maximum potential backoff delays.
	FlagLoginAttemptsSenderAMQPPublishMaxRetryDelay string = "login-attempts-sender-amqp-publish-max-retry-delay"
)

// login attempts sender kafka
const (
	// FlagLoginAttemptsSenderKafkaBrokers identifies network cluster addresses servicing Kafka streaming nodes.
	FlagLoginAttemptsSenderKafkaBrokers string = "login-attempts-sender-kafka-brokers"
	// FlagLoginAttemptsSenderKafkaTargetName defines destination topic patterns routing login tracking payloads.
	FlagLoginAttemptsSenderKafkaTargetName string = "login-attempts-sender-kafka-target-name"
	// FlagLoginAttemptsSenderKafkaConnectTimeout limits socket initialization budgets under active Kafka dialers.
	FlagLoginAttemptsSenderKafkaConnectTimeout string = "login-attempts-sender-kafka-connect-timeout"
	// FlagLoginAttemptsSenderKafkaIdleTimeout bounds durability metrics for resting broker pipeline connections.
	FlagLoginAttemptsSenderKafkaIdleTimeout string = "login-attempts-sender-kafka-idle-timeout"
	// FlagLoginAttemptsSenderKafkaShutdownTimeout restricts termination block intervals allowed for internal producer cleanup.
	FlagLoginAttemptsSenderKafkaShutdownTimeout string = "login-attempts-sender-kafka-shutdown-timeout"
	// FlagLoginAttemptsSenderKafkaPublishMaxTryAttempts manages retry policy limitations under producer transport faults.
	FlagLoginAttemptsSenderKafkaPublishMaxTryAttempts string = "login-attempts-sender-kafka-publish-max-try-attempts"
	// FlagLoginAttemptsSenderKafkaPublishMaxRetryDelay enforces upper thresholds bounding maximum retry intervals under cluster degradation.
	FlagLoginAttemptsSenderKafkaPublishMaxRetryDelay string = "login-attempts-sender-kafka-publish-max-retry-delay"
	// FlagLoginAttemptsSenderKafkaPublishBaseRetryDelay sets initial fallback intervals for backoff routing rules.
	FlagLoginAttemptsSenderKafkaPublishBaseRetryDelay string = "login-attempts-sender-kafka-publish-base-retry-delay"
	// FlagLoginAttemptsSenderKafkaUsername handles SASL mechanism parameters identifying active streaming users.
	FlagLoginAttemptsSenderKafkaUsername string = "login-attempts-sender-kafka-username"
	// FlagLoginAttemptsSenderKafkaPassword secures SASL credential verification tokens.
	FlagLoginAttemptsSenderKafkaPassword string = "login-attempts-sender-kafka-password"
	// FlagLoginAttemptsSenderKafkaBatchSize handles quantitative constraints sizing accumulated batch records counts.
	FlagLoginAttemptsSenderKafkaBatchSize string = "login-attempts-sender-kafka-batch-size"
	// FlagLoginAttemptsSenderKafkaBatchBytes sets byte size allocation limits triggering automated batch flushes.
	FlagLoginAttemptsSenderKafkaBatchBytes string = "login-attempts-sender-kafka-batch-bytes"
	// FlagLoginAttemptsSenderKafkaBatchTimeout dictates temporal constraints forcing lazy batch queues onto networks.
	FlagLoginAttemptsSenderKafkaBatchTimeout string = "login-attempts-sender-kafka-batch-timeout"
	// FlagLoginAttemptsSenderKafkaWriteTimeout bounds explicit durations assigned for network block socket transmissions.
	FlagLoginAttemptsSenderKafkaWriteTimeout string = "login-attempts-sender-kafka-write-timeout"
	// FlagLoginAttemptsSenderKafkaRequiredAcks manages consistency metrics configuring required cluster brokers confirmations (e.g. 0, 1, all).
	FlagLoginAttemptsSenderKafkaRequiredAcks string = "login-attempts-sender-kafka-required-acks"
)

// data audit sender amqp connector
//
//nolint:gosec // G101 : app params
const (
	// FlagDataAuditSenderAMQPConnectorURL maps the baseline AMQP network address for the target cluster endpoint.
	FlagDataAuditSenderAMQPConnectorURL string = "data-audit-sender-amqp-connector-url"
	// FlagDataAuditSenderAMQPConnectorUsername defines user parameters for authenticating structural broker sessions.
	FlagDataAuditSenderAMQPConnectorUsername string = "data-audit-sender-amqp-connector-username"
	// FlagDataAuditSenderAMQPConnectorPassword establishes cryptographic secret tokens identifying connection channels.
	FlagDataAuditSenderAMQPConnectorPassword string = "data-audit-sender-amqp-connector-password"
	// FlagDataAuditSenderAMQPConnectorConnectTimeout sets network allocation block limits establishing connections.
	FlagDataAuditSenderAMQPConnectorConnectTimeout string = "data-audit-sender-amqp-connector-connect-timeout"
	// FlagDataAuditSenderAMQPConnectorWriteTimeout caps maximum block durations pushing byte messages onto sockets.
	FlagDataAuditSenderAMQPConnectorWriteTimeout string = "data-audit-sender-amqp-connector-write-timeout"
	// FlagDataAuditSenderAMQPConnectorIdleTimeout specifies durability boundaries for resting unassigned broker channels.
	FlagDataAuditSenderAMQPConnectorIdleTimeout string = "data-audit-sender-amqp-connector-idle-timeout"
	// FlagDataAuditSenderAMQPConnectorShutdownTimeout handles block intervals dedicated to closing server connection context maps gracefully.
	FlagDataAuditSenderAMQPConnectorShutdownTimeout string = "data-audit-sender-amqp-connector-shutdown-timeout"
)

// data audit sender
const (
	// FlagDataAuditSenderKind chooses specific broker types driving events tracking workflows (e.g. amqp, kafka).
	FlagDataAuditSenderKind string = "data-audit-sender-kind"
	// FlagDataAuditSenderNotifyTimeout limits wait budgets assigned for async pubsub broadcasts sequences.
	FlagDataAuditSenderNotifyTimeout string = "data-audit-sender-notify-timeout"
)

// data audit sender amqp
const (
	// FlagDataAuditSenderAMQPTargetName defines the routing exchange or queue target parameter name inside AMQP broker layouts.
	FlagDataAuditSenderAMQPTargetName string = "data-audit-sender-amqp-target-name"
	// FlagDataAuditSenderAMQPConnectTimeout limits socket allocation budgets for publishing streams.
	FlagDataAuditSenderAMQPConnectTimeout string = "data-audit-sender-amqp-connect-timeout"
	// FlagDataAuditSenderAMQPShutdownTimeout establishes safe timeout boundaries during event streaming closures.
	FlagDataAuditSenderAMQPShutdownTimeout string = "data-audit-sender-amqp-shutdown-timeout"
	// FlagDataAuditSenderAMQPPublishMaxTryAttempts caps total retries loops executed upon network publishing degradations.
	FlagDataAuditSenderAMQPPublishMaxTryAttempts string = "data-audit-sender-amqp-publish-max-try-attempts"
	// FlagDataAuditSenderAMQPPublishBaseRetryDelay defines initial backoff durations before retrying failed message publish routines.
	FlagDataAuditSenderAMQPPublishBaseRetryDelay string = "data-audit-sender-amqp-publish-base-retry-delay"
	// FlagDataAuditSenderAMQPPublishMaxRetryDelay sets a strict cap bounding maximum potential backoff delays.
	FlagDataAuditSenderAMQPPublishMaxRetryDelay string = "data-audit-sender-amqp-publish-max-retry-delay"
)

// data audit sender kafka
//
//nolint:gosec // G101 : app params
const (
	// FlagDataAuditSenderKafkaBrokers identifies network cluster addresses servicing Kafka streaming nodes.
	FlagDataAuditSenderKafkaBrokers string = "data-audit-sender-kafka-brokers"
	// FlagDataAuditSenderKafkaTargetName defines destination topic patterns routing login tracking payloads.
	FlagDataAuditSenderKafkaTargetName string = "data-audit-sender-kafka-target-name"
	// FlagDataAuditSenderKafkaConnectTimeout limits socket initialization budgets under active Kafka dialers.
	FlagDataAuditSenderKafkaConnectTimeout string = "data-audit-sender-kafka-connect-timeout"
	// FlagDataAuditSenderKafkaIdleTimeout bounds durability metrics for resting broker pipeline connections.
	FlagDataAuditSenderKafkaIdleTimeout string = "data-audit-sender-kafka-idle-timeout"
	// FlagDataAuditSenderKafkaShutdownTimeout restricts termination block intervals allowed for internal producer cleanup.
	FlagDataAuditSenderKafkaShutdownTimeout string = "data-audit-sender-kafka-shutdown-timeout"
	// FlagDataAuditSenderKafkaPublishMaxTryAttempts manages retry policy limitations under producer transport faults.
	FlagDataAuditSenderKafkaPublishMaxTryAttempts string = "data-audit-sender-kafka-publish-max-try-attempts"
	// FlagDataAuditSenderKafkaPublishMaxRetryDelay enforces upper thresholds bounding maximum retry intervals under cluster degradation.
	FlagDataAuditSenderKafkaPublishMaxRetryDelay string = "data-audit-sender-kafka-publish-max-retry-delay"
	// FlagDataAuditSenderKafkaPublishBaseRetryDelay sets initial fallback intervals for backoff routing rules.
	FlagDataAuditSenderKafkaPublishBaseRetryDelay string = "data-audit-sender-kafka-publish-base-retry-delay"
	// FlagDataAuditSenderKafkaUsername handles SASL mechanism parameters identifying active streaming users.
	FlagDataAuditSenderKafkaUsername string = "data-audit-sender-kafka-username"
	// FlagDataAuditSenderKafkaPassword secures SASL credential verification tokens.
	FlagDataAuditSenderKafkaPassword string = "data-audit-sender-kafka-password"
	// FlagDataAuditSenderKafkaBatchSize handles quantitative constraints sizing accumulated batch records counts.
	FlagDataAuditSenderKafkaBatchSize string = "data-audit-sender-kafka-batch-size"
	// FlagDataAuditSenderKafkaBatchBytes sets byte size allocation limits triggering automated batch flushes.
	FlagDataAuditSenderKafkaBatchBytes string = "data-audit-sender-kafka-batch-bytes"
	// FlagDataAuditSenderKafkaBatchTimeout dictates temporal constraints forcing lazy batch queues onto networks.
	FlagDataAuditSenderKafkaBatchTimeout string = "data-audit-sender-kafka-batch-timeout"
	// FlagDataAuditSenderKafkaWriteTimeout bounds explicit durations assigned for network block socket transmissions.
	FlagDataAuditSenderKafkaWriteTimeout string = "data-audit-sender-kafka-write-timeout"
	// FlagDataAuditSenderKafkaRequiredAcks manages consistency metrics configuring required cluster brokers confirmations (e.g. 0, 1, all).
	FlagDataAuditSenderKafkaRequiredAcks string = "data-audit-sender-kafka-required-acks"
)

const (
	FlagAuthAuditOutboxTCTailLimit        string = "auth-audit-outbox-tc-tail-limit"
	FlagAuthAuditOutboxTCTailInterval     string = "auth-audit-outbox-tc-tail-interval"
	FlagAuthAuditOutboxTCTailCut          string = "auth-audit-outbox-tc-tail-cut"
	FlagAuthAuditOutboxTCWorkerCount      string = "auth-audit-outbox-tc-worker-count"
	FlagAuthAuditOutboxTCDataCapacity     string = "auth-audit-outbox-tc-data-capacity"
	FlagAuthAuditOutboxTCCompleteProcess  string = "auth-audit-outbox-tc-complete-process"
	FlagAuthAuditOutboxTCStartInterval    string = "auth-audit-outbox-tc-start-interval"
	FlagAuthAuditOutboxTCScheduleInterval string = "auth-audit-outbox-tc-schedule-interval"
	FlagAuthAuditOutboxTCStopTimeout      string = "auth-audit-outbox-tc-stop-timeout"
)

const (
	FlagDataAuditOutboxTCTailLimit        string = "data-audit-outbox-tc-tail-limit"
	FlagDataAuditOutboxTCTailInterval     string = "data-audit-outbox-tc-tail-interval"
	FlagDataAuditOutboxTCTailCut          string = "data-audit-outbox-tc-tail-cut"
	FlagDataAuditOutboxTCWorkerCount      string = "data-audit-outbox-tc-worker-count"
	FlagDataAuditOutboxTCDataCapacity     string = "data-audit-outbox-tc-data-capacity"
	FlagDataAuditOutboxTCCompleteProcess  string = "data-audit-outbox-tc-complete-process"
	FlagDataAuditOutboxTCStartInterval    string = "data-audit-outbox-tc-start-interval"
	FlagDataAuditOutboxTCScheduleInterval string = "data-audit-outbox-tc-schedule-interval"
	FlagDataAuditOutboxTCStopTimeout      string = "data-audit-outbox-tc-stop-timeout"
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

// login attempts sender amqp connector
//
//nolint:gosec // G101 : app conf keys
const (
	// keyLoginAttemptsSenderAMQPConnectorURL tracks the nested viper schema dot-key location for the baseline AMQP cluster network address string.
	keyLoginAttemptsSenderAMQPConnectorURL string = "login_attempts_connector.url"
	// keyLoginAttemptsSenderAMQPConnectorUsername maps nested path configuration nodes governing authorization user credentials blocks.
	keyLoginAttemptsSenderAMQPConnectorUsername string = "login_attempts_connector.username"
	// keyLoginAttemptsSenderAMQPConnectorPassword protects nested trajectories parsing access passwords.
	keyLoginAttemptsSenderAMQPConnectorPassword string = "login_attempts_connector.password"
	// keyLoginAttemptsSenderAMQPConnectorConnectTimeout points to internal property structures managing connection budget limits.
	keyLoginAttemptsSenderAMQPConnectorConnectTimeout string = "login_attempts_connector.connect_timeout"
	// keyLoginAttemptsSenderAMQPConnectorWriteTimeout targets configurations specifying maximal write block timeouts parameters.
	keyLoginAttemptsSenderAMQPConnectorWriteTimeout string = "login_attempts_connector.write_timeout"
	// keyLoginAttemptsSenderAMQPConnectorIdleTimeout dictates internal trajectories formatting unassigned channels durability lifetimes.
	keyLoginAttemptsSenderAMQPConnectorIdleTimeout string = "login_attempts_connector.idle_timeout"
	// keyLoginAttemptsSenderAMQPConnectorShutdownTimeout binds execution properties framing termination intervals budgets.
	keyLoginAttemptsSenderAMQPConnectorShutdownTimeout string = "login_attempts_connector.shutdown_timeout"
)

// login attempts sender
const (
	// keyLoginAttemptsSenderKind chooses specific broker types driving events tracking workflows (e.g. amqp, kafka).
	keyLoginAttemptsSenderKind string = "login_attempts_sender.sender_kind"
	// keyLoginAttemptsSenderNotifyTimeout limits wait budgets assigned for async pubsub broadcasts sequences.
	keyLoginAttemptsSenderNotifyTimeout string = "login_attempts_sender.notify_timeout"
)

// login attempts sender amqp
const (
	// keyLoginAttemptsSenderAMQPTargetName targets configuration nodes establishing destination queues or exchanges.
	keyLoginAttemptsSenderAMQPTargetName string = "login_attempts_sender.amqp_config.target_name"
	// keyLoginAttemptsSenderAMQPConnectTimeout dictates socket initialization budgets nested under parameters keys.
	keyLoginAttemptsSenderAMQPConnectTimeout string = "login_attempts_sender.amqp_config.connect_timeout"
	// keyLoginAttemptsSenderAMQPShutdownTimeout manages contextual timeout structures applied across streaming closures operations.
	keyLoginAttemptsSenderAMQPShutdownTimeout string = "login_attempts_sender.amqp_config.shutdown_timeout"
	// keyLoginAttemptsSenderAMQPPublishMaxTryAttempts restricts total retries counters upon network pipeline errors loops.
	keyLoginAttemptsSenderAMQPPublishMaxTryAttempts string = "login_attempts_sender.amqp_config.publish_max_try_attempts"
	// keyLoginAttemptsSenderAMQPPublishBaseRetryDelay maps property paths specifying backoff multiplier initial values duration.
	keyLoginAttemptsSenderAMQPPublishBaseRetryDelay string = "login_attempts_sender.amqp_config.publish_base_retry_delay"
	// keyLoginAttemptsSenderAMQPPublishMaxRetryDelay caps maximal delay limits managing cascading backoff parameters structures.
	keyLoginAttemptsSenderAMQPPublishMaxRetryDelay string = "login_attempts_sender.amqp_config.publish_max_retry_delay"
)

// login attemtpts sedner kafka
const (
	// keyLoginAttemptsSenderKafkaBrokers identifies compound array configurations trajectories pointing to target Kafka network nodes addresses.
	keyLoginAttemptsSenderKafkaBrokers string = "login_attempts_sender.kafka_config.brokers"
	// keyLoginAttemptsSenderKafkaTargetName defines nested topic destination properties.
	keyLoginAttemptsSenderKafkaTargetName string = "login_attempts_sender.kafka_config.target_name"
	// keyLoginAttemptsSenderKafkaConnectTimeout manages dialer socket initialization duration parameters keys.
	keyLoginAttemptsSenderKafkaConnectTimeout string = "login_attempts_sender.kafka_config.connect_timeout"
	// keyLoginAttemptsSenderKafkaIdleTimeout points to structures specifying passive channels connection lifetime guidelines.
	keyLoginAttemptsSenderKafkaIdleTimeout string = "login_attempts_sender.kafka_config.idle_timeout"
	// keyLoginAttemptsSenderKafkaShutdownTimeout constraints cleanup execution durations allowed during stream engine graceful shutdowns.
	keyLoginAttemptsSenderKafkaShutdownTimeout string = "login_attempts_sender.kafka_config.shutdown_timeout"
	// keyLoginAttemptsSenderKafkaPublishMaxTryAttempts targets retry configuration metrics driven under producer faults triggers.
	keyLoginAttemptsSenderKafkaPublishMaxTryAttempts string = "login_attempts_sender.kafka_config.publish_max_try_attempts"
	// keyLoginAttemptsSenderKafkaPublishMaxRetryDelay specifies ceiling boundaries for backoff delay parameters routines.
	keyLoginAttemptsSenderKafkaPublishMaxRetryDelay string = "login_attempts_sender.kafka_config.publish_max_retry_delay"
	// keyLoginAttemptsSenderKafkaPublishBaseRetryDelay sets baseline starting backoff delays duration indices.
	keyLoginAttemptsSenderKafkaPublishBaseRetryDelay string = "login_attempts_sender.kafka_config.publish_base_retry_delay"
	// keyLoginAttemptsSenderKafkaUsername resolves SASL auth parameters mapping active users names.
	keyLoginAttemptsSenderKafkaUsername string = "login_attempts_sender.kafka_config.username"
	// keyLoginAttemptsSenderKafkaPassword secures nested paths handling SASL verification secret tokens.
	keyLoginAttemptsSenderKafkaPassword string = "login_attempts_sender.kafka_config.password"
	// keyLoginAttemptsSenderKafkaBatchSize binds internal arrays constraints pacing maximum batch records numbers.
	keyLoginAttemptsSenderKafkaBatchSize string = "login_attempts_sender.kafka_config.batch_size"
	// keyLoginAttemptsSenderKafkaBatchBytes paths configurations checking buffer size limits in bytes before forcing automated network flushes.
	keyLoginAttemptsSenderKafkaBatchBytes string = "login_attempts_sender.kafka_config.batch_bytes"
	// keyLoginAttemptsSenderKafkaBatchTimeout manages lazy queue timer triggers forcing buffer flushes to downstream networks interfaces.
	keyLoginAttemptsSenderKafkaBatchTimeout string = "login_attempts_sender.kafka_config.batch_timeout"
	// keyLoginAttemptsSenderKafkaWriteTimeout dictates socket block durability limits managing individual record packages writes operations.
	keyLoginAttemptsSenderKafkaWriteTimeout string = "login_attempts_sender.kafka_config.write_timeout"
	// keyLoginAttemptsSenderKafkaRequiredAcks structures consensus verification metrics configuring brokers replication confirmations criteria numbers.
	keyLoginAttemptsSenderKafkaRequiredAcks string = "login_attempts_sender.kafka_config.required_acks"
)

// data audit sender amqp connector
//
//nolint:gosec // G101 : app conf keys
const (
	// keyDataAuditSenderAMQPConnectorURL tracks the nested viper schema dot-key location for the baseline AMQP cluster network address string.
	keyDataAuditSenderAMQPConnectorURL string = "data_audit_connector.url"
	// keyDataAuditSenderAMQPConnectorUsername maps nested path configuration nodes governing authorization user credentials blocks.
	keyDataAuditSenderAMQPConnectorUsername string = "data_audit_connector.username"
	// keyDataAuditSenderAMQPConnectorPassword protects nested trajectories parsing access passwords.
	keyDataAuditSenderAMQPConnectorPassword string = "data_audit_connector.password"
	// keyDataAuditSenderAMQPConnectorConnectTimeout points to internal property structures managing connection budget limits.
	keyDataAuditSenderAMQPConnectorConnectTimeout string = "data_audit_connector.connect_timeout"
	// keyDataAuditSenderAMQPConnectorWriteTimeout targets configurations specifying maximal write block timeouts parameters.
	keyDataAuditSenderAMQPConnectorWriteTimeout string = "data_audit_connector.write_timeout"
	// keyDataAuditSenderAMQPConnectorIdleTimeout dictates internal trajectories formatting unassigned channels durability lifetimes.
	keyDataAuditSenderAMQPConnectorIdleTimeout string = "data_audit_connector.idle_timeout"
	// keyDataAuditSenderAMQPConnectorShutdownTimeout binds execution properties framing termination intervals budgets.
	keyDataAuditSenderAMQPConnectorShutdownTimeout string = "data_audit_connector.shutdown_timeout"
)

// data audit sender
const (
	// keyLoginAttemptsSenderKind chooses specific broker types driving events tracking workflows (e.g. amqp, kafka).
	keyDataAuditSenderKind string = "login_attempts_sender.sender_kind"
	// keyLoginAttemptsSenderNotifyTimeout limits wait budgets assigned for async pubsub broadcasts sequences.
	keyDataAuditSenderNotifyTimeout string = "login_attempts_sender.notify_timeout"
)

// data audit sender amqp
const (
	// keyDataAuditSenderAMQPTargetName targets configuration nodes establishing destination queues or exchanges.
	keyDataAuditSenderAMQPTargetName string = "data_audit_sender.amqp_config.target_name"
	// keyDataAuditSenderAMQPConnectTimeout dictates socket initialization budgets nested under parameters keys.
	keyDataAuditSenderAMQPConnectTimeout string = "data_audit_sender.amqp_config.connect_timeout"
	// keyDataAuditSenderAMQPShutdownTimeout manages contextual timeout structures applied across streaming closures operations.
	keyDataAuditSenderAMQPShutdownTimeout string = "data_audit_sender.amqp_config.shutdown_timeout"
	// keyDataAuditSenderAMQPPublishMaxTryAttempts restricts total retries counters upon network pipeline errors loops.
	keyDataAuditSenderAMQPPublishMaxTryAttempts string = "data_audit_sender.amqp_config.publish_max_try_attempts"
	// keyDataAuditSenderAMQPPublishBaseRetryDelay maps property paths specifying backoff multiplier initial values duration.
	keyDataAuditSenderAMQPPublishBaseRetryDelay string = "data_audit_sender.amqp_config.publish_base_retry_delay"
	// keyDataAuditSenderAMQPPublishMaxRetryDelay caps maximal delay limits managing cascading backoff parameters structures.
	keyDataAuditSenderAMQPPublishMaxRetryDelay string = "data_audit_sender.amqp_config.publish_max_retry_delay"
)

// data audit sender kafka
//
//nolint:gosec // G101 : app params
const (
	// keyDataAuditSenderKafkaBrokers identifies compound array configurations trajectories pointing to target Kafka network nodes addresses.
	keyDataAuditSenderKafkaBrokers string = "data_audit_sender.kafka_config.brokers"
	// keyDataAuditSenderKafkaTargetName defines nested topic destination properties.
	keyDataAuditSenderKafkaTargetName string = "data_audit_sender.kafka_config.target_name"
	// keyDataAuditSenderKafkaConnectTimeout manages dialer socket initialization duration parameters keys.
	keyDataAuditSenderKafkaConnectTimeout string = "data_audit_sender.kafka_config.connect_timeout"
	// keyDataAuditSenderKafkaIdleTimeout points to structures specifying passive channels connection lifetime guidelines.
	keyDataAuditSenderKafkaIdleTimeout string = "data_audit_sender.kafka_config.idle_timeout"
	// keyDataAuditSenderKafkaShutdownTimeout constraints cleanup execution durations allowed during stream engine graceful shutdowns.
	keyDataAuditSenderKafkaShutdownTimeout string = "data_audit_sender.kafka_config.shutdown_timeout"
	// keyDataAuditSenderKafkaPublishMaxTryAttempts targets retry configuration metrics driven under producer faults triggers.
	keyDataAuditSenderKafkaPublishMaxTryAttempts string = "data_audit_sender.kafka_config.publish_max_try_attempts"
	// keyDataAuditSenderKafkaPublishMaxRetryDelay specifies ceiling boundaries for backoff delay parameters routines.
	keyDataAuditSenderKafkaPublishMaxRetryDelay string = "data_audit_sender.kafka_config.publish_max_retry_delay"
	// keyDataAuditSenderKafkaPublishBaseRetryDelay sets baseline starting backoff delays duration indices.
	keyDataAuditSenderKafkaPublishBaseRetryDelay string = "data_audit_sender.kafka_config.publish_base_retry_delay"
	// keyDataAuditSenderKafkaUsername resolves SASL auth parameters mapping active users names.
	keyDataAuditSenderKafkaUsername string = "data_audit_sender.kafka_config.username"
	// keyDataAuditSenderKafkaPassword secures nested paths handling SASL verification secret tokens.
	keyDataAuditSenderKafkaPassword string = "data_audit_sender.kafka_config.password"
	// keyDataAuditSenderKafkaBatchSize binds internal arrays constraints pacing maximum batch records numbers.
	keyDataAuditSenderKafkaBatchSize string = "data_audit_sender.kafka_config.batch_size"
	// keyDataAuditSenderKafkaBatchBytes paths configurations checking buffer size limits in bytes before forcing automated network flushes.
	keyDataAuditSenderKafkaBatchBytes string = "data_audit_sender.kafka_config.batch_bytes"
	// keyDataAuditSenderKafkaBatchTimeout manages lazy queue timer triggers forcing buffer flushes to downstream networks interfaces.
	keyDataAuditSenderKafkaBatchTimeout string = "data_audit_sender.kafka_config.batch_timeout"
	// keyDataAuditSenderKafkaWriteTimeout dictates socket block durability limits managing individual record packages writes operations.
	keyDataAuditSenderKafkaWriteTimeout string = "data_audit_sender.kafka_config.write_timeout"
	// keyDataAuditSenderKafkaRequiredAcks structures consensus verification metrics configuring brokers replication confirmations criteria numbers.
	keyDataAuditSenderKafkaRequiredAcks string = "data_audit_sender.kafka_config.required_acks"
)

const (
	keyAuthAuditOutboxTCTailLimit          string = "auth_audit_outbox_tail_cutter.tail_limit"
	keyAuthAuditOutboxTCTailInterval       string = "auth_audit_outbox_tail_cutter.tail_interval"
	keyAuthAuditOutboxTCTailCut            string = "auth_audit_outbox_tail_cutter.tail_cut"
	keyAuthAuditOutboxTCWorkerCount        string = "auth_audit_outbox_tail_cutter.dispatcher.worker_count"
	keyAuthAuditOutboxTCDataCapacity       string = "auth_audit_outbox_tail_cutter.dispatcher.data_capacity"
	keyAuthAuditOutboxTCCompleteProcessing string = "auth_audit_outbox_tail_cutter.dispatcher.complete_process"
	keyAuthAuditOutboxTCStartInterval      string = "auth_audit_outbox_tail_cutter.dispatcher.start_interval"
	keyAuthAuditOutboxTCScheduleInterval   string = "auth_audit_outbox_tail_cutter.dispatcher.schedule_interval"
	keyAuthAuditOutboxTCStopTimeout        string = "auth_audit_outbox_tail_cutter.dispatcher.stop_timeout"
)

const (
	keyDataAuditOutboxTCTailLimit          string = "data_audit_outbox_tail_cutter.tail_limit"
	keyDataAuditOutboxTCTailInterval       string = "data_audit_outbox_tail_cutter.tail_interval"
	keyDataAuditOutboxTCTailCut            string = "data_audit_outbox_tail_cutter.tail_cut"
	keyDataAuditOutboxTCWorkerCount        string = "data_audit_outbox_tail_cutter.dispatcher.worker_count"
	keyDataAuditOutboxTCDataCapacity       string = "data_audit_outbox_tail_cutter.dispatcher.data_capacity"
	keyDataAuditOutboxTCCompleteProcessing string = "data_audit_outbox_tail_cutter.dispatcher.complete_process"
	keyDataAuditOutboxTCStartInterval      string = "data_audit_outbox_tail_cutter.dispatcher.start_interval"
	keyDataAuditOutboxTCScheduleInterval   string = "data_audit_outbox_tail_cutter.dispatcher.schedule_interval"
	keyDataAuditOutboxTCStopTimeout        string = "data_audit_outbox_tail_cutter.dispatcher.stop_timeout"
)
