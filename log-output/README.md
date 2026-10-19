# Log output app

## Create container

Test locally with the following steps:

```bash
docker build --tag oppearo/log-output .
docker volume create log-data
docker run --rm -d --name generator -v log-data:/shared log-output /bin/generator
docker run --rm -d --name server 5000:5000 -v log-data:/shared log-output /bin/server
curl http://localhost:5000/status
```

Push to Docker Hub:

```bash
docker push oppearo/log-output
```

See `manifests/deployment.yaml` for configuring the shared volume for two containers in a pod.

## Example outputs

In (localhost:8081/status)[http://localhost:8081/status]

`2026-10-19T15:48:04Z: 98bac9c6-01e8-47c4-b4c8-d10fd2125b09`

And output from curl and kubectl logs for the pod:

```bash
❯ curl localhost:8081/status
2026-10-19T15:48:04Z: 98bac9c6-01e8-47c4-b4c8-d10fd2125b09

❯ kubectl logs log-output-dep-8499f658b7-bbjpg
Defaulted container "log-output-generator" out of: log-output-generator, log-output-contentserver
2026-10-19T15:47:49Z: 98bac9c6-01e8-47c4-b4c8-d10fd2125b09
2026-10-19T15:47:54Z: 98bac9c6-01e8-47c4-b4c8-d10fd2125b09
2026-10-19T15:47:59Z: 98bac9c6-01e8-47c4-b4c8-d10fd2125b09
2026-10-19T15:48:04Z: 98bac9c6-01e8-47c4-b4c8-d10fd2125b09
```