# Kafka Topic Watcher

Docker utility container to run as a Pod in Kubernetes. The [txn2/kwatch] Pod log display Kakfa messages for a given topic.

Example Pod:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: kafka-kwatch
  namespace: the-project
spec:
  containers:
  - name: kafka-kwatch
    image: txn2/kwatch:1.0.1
    env:
    - name: BROKER
      value: "kafka-headless:9092"
    - name: TOPIC
      value: "some-topic"
    # Setting the offset to -1 means to seek to the first offset.
    # Setting the offset to -2 means to seek to the last offset.
    - name: OFFSET
      value: "-2"
    - name: PARTITION
      value: "0"
```

Tail the kwatch logs:
```bash
kubectl logs -f kafka-kwatch -n the-project
```

[txn2/kwatch]:https://hub.docker.com/r/txn2/kwatch/