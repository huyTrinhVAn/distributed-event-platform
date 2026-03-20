Introduction


                               +----------------------+
                               |   producer-simulator |
                               |  (auth/order/pay...) |
                               +----------+-----------+
                                          |
                                          | publish raw logs/events
                                          v
                                +--------------------+
                                |   Kafka Topic      |
                                |    events.raw      |
                                +---------+----------+
                                          |
                                          | consume
                                          v
                              +-------------------------+
                              |   ingestion-service     |
                              | - validate schema       |
                              | - enrich metadata       |
                              | - assign event_id       |
                              | - normalize format      |
                              +-----------+-------------+
                                          |
                       +------------------+------------------+
                       |                                     |
                       | valid                               | invalid
                       v                                     v
             +----------------------+              +----------------------+
             | Kafka Topic          |              | Kafka Topic          |
             | events.validated     |              | events.invalid       |
             +----------+-----------+              +----------------------+
                        |
                        | consume
                        v
              +--------------------------+
              |   processor-service      |
              | - classify severity      |
              | - route by event_type    |
              | - detect critical events |
              | - send to downstream     |
              +-----+----------+---------+
                    |          |
          normal    |          | retry needed / downstream failed
                    |          |
                    v          v
      +------------------+   +------------------+
      | Kafka Topic      |   | Kafka Topic      |
      | events.processed |   | events.retry     |
      +--------+---------+   +--------+---------+
               |                      |
               |                      | consume
               |                      v
               |            +----------------------+
               |            |    retry-worker      |
               |            | - backoff policy     |
               |            | - retry_count += 1   |
               |            | - republish          |
               |            +-----+-----------+----+
               |                  |           |
               |                  | success   | max retry exceeded
               |                  v           v
               |      +------------------+   +------------------+
               |      | events.processed |   |   events.dlq     |
               |      +------------------+   +------------------+
               |
               | critical / alert-worthy
               v
      +----------------------+
      | Kafka Topic          |
      | alerts.generated     |
      +----------+-----------+
                 |
                 | consume
                 v
      +----------------------+
      |    alert-service     |
      | - log alert          |
      | - email/webhook mock |
      +----------------------+

     ---------------------------------------------------------------
     Observability / Ops
     ---------------------------------------------------------------
     Each Go service exposes:
       /healthz
       /readyz
       /metrics

     Prometheus <---------------- scrapes metrics from all services
     Grafana    <---------------- dashboards: throughput, errors, lag, DLQ

# /opt/kafka/bin/kafka-topics.sh --list --bootstrap-server localhost:9092

# /opt/kafka/bin/kafka-topics.sh --create --topic events.raw --bootstrap-server localhost:9092
# /opt/kafka/bin/kafka-topics.sh --list --bootstrap-server localhost:9092
# >   --bootstrap-server localhost:9092


# Flow hien tai:
# POST /publish
# -> producer-simulator tạo event
# -> publish vào events.raw
# -> ingestion-service consume events.raw
# -> unmarshal + enrich
# -> publish vào events.validated