# Demo of how the reset would look like
docker exec -it kafka /opt/kafka/bin/kafka-consumer-groups.sh \
  --bootstrap-server localhost:9092 \
  --group order-processors \
  --reset-offsets --to-earliest \
  --topic orders \
  --dry-run

docker exec -it kafka /opt/kafka/bin/kafka-consumer-groups.sh \
  --bootstrap-server localhost:9092 \
  --group order-processors \
  --reset-offsets --to-earliest \
  --topic orders \
  --execute