Introduction


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