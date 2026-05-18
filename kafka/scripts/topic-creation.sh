docker exec -it kafka /opt/kafka/bin/kafka-topics.sh \
  --create \
  --topic orders \
  --partitions 3 \
  --replication-factor 1 \
  --bootstrap-server localhost:9092