# Demo of how the reset would look like
docker exec -it kafka /opt/kafka/bin/kafka-consumer-groups.sh \
  --bootstrap-server localhost:9092 \
  --group $1 \
  --reset-offsets --to-earliest \
  --topic $2 \
  --dry-run

docker exec -it kafka /opt/kafka/bin/kafka-consumer-groups.sh \
  --bootstrap-server localhost:9092 \
  --group $1 \
  --reset-offsets --to-earliest \
  --topic $2 \
  --execute