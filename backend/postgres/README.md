# PostgreSQL


## Environment based configuration

Configuration for the PostgreSQL connection pool of the microservice.

* `POSTGRES_PORT` default: `5432`
    * Port to be used for listening used if address is not specified
* `POSTGRES_HOST` default: `localhost`
    * Host where the PostgreSQL can be found (dns or IP)
* `POSTGRES_PASSWORD` default: `pace1234!`
    * password to access the database
* `POSTGRES_USER` default: `postgres`
    * postgres user to access the database
* `POSTGRES_DB` default: `postgres`
    * database to access
* `POSTGRES_MAX_RETRIES` default: `5`
    * Maximum number of retries before giving up
* `POSTGRES_RETRY_STATEMENT_TIMEOUT` default: `false`
    * Whether to retry queries cancelled because of statement_timeout
* `POSTGRES_MIN_RETRY_BACKOFF` default: `250ms`
    *  Minimum backoff between each retry
* `POSTGRES_MAX_RETRY_BACKOFF` default: `4s`
    * Maximum backoff between each retry
* `POSTGRES_DIAL_TIMEOUT` default: `5s`
    * Dial timeout for establishing new connections
* `POSTGRES_READ_TIMEOUT` default: `30s`
    *  Timeout for socket reads. If reached, commands will fail with a timeout instead of blocking
* `POSTGRES_WRITE_TIMEOUT` default: `30s`
    * Timeout for socket writes. If reached, commands will fail with a timeout instead of blocking.
* `POSTGRES_POOL_SIZE` default: `100`
    * Maximum number of socket connections
* `POSTGRES_MIN_IDLE_CONNECTIONS` default: `10`
    * Minimum number of idle connections which is useful when establishing new connection is slow
* `POSTGRES_MAX_CONN_AGE` default: `30m`
    * Connection age at which client retires (closes) the connection
* `POSTGRES_POOL_TIMEOUT` default: `31s`
    * Time for which client waits for free connection if all connections are busy before returning an error
* `POSTGRES_IDLE_TIMEOUT` default: `5m`
    * Amount of time after which client closes idle connections
* `POSTGRES_IDLE_CHECK_FREQUENCY` default: `1m`
    * Frequency of idle checks made by idle connections reaper
