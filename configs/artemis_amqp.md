# apache artemis v2.54.0
console username: `artemis-activemq`

console password: `artemis-activemq`

console URL: http://localhost:8161

## URL
`amqp://localhost:5672`
## address
`tiny.auth`

## multicast (or topic)
`login.attempts`

## add user
```
/var/lib/artemis-test-cluster/bin/artemis user add \
  --user-command-user svc-auth \
  --user-command-password test \
  --role auth-producer
```
```
/var/lib/artemis-test-cluster/bin/artemis user add \
  --user-command-user svc-audit \
  --user-command-password test \
  --role auth-consumer
```

### producer: username:`svc-auth` password:`test`
### consumer: username:`svc-audit` password:`test`
### consumer: `svc-some_service` :-)
