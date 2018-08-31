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
    image: txn2/kwatch:1.0.0
    env:
    - name: TOPIC
      value: "some-topic"
```

[txn2/kwatch]:https://hub.docker.com/r/txn2/kwatch/