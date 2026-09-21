# Log output app

Create with:

```bash
docker build --tag oppearo/log-output
```

Import into default cluster and create deployment:

```bash
k3d image import oppearo/log-output
kubectl create deployment hashgenerator-dep --image=oppearo/log-output
# I did locally so edited imagePullPolicy of deployment
kubectl edit deployment hashgenerator-dep

$ kubectl get pods                             NAME                                 READY   STATUS    RESTARTS   AGE
hashgenerator-dep-56b76cf5fd-dkb7b   1/1     Running   0          6s
```

And output in logs:

```bash
kubectl logs -f hashgenerator-dep-56b76cf5fd-dkb7b
2026-09-21T19:16:41Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:16:46Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:16:50Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:16:55Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:17:00Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:17:05Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:17:10Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:17:15Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:17:19Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:17:24Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:17:29Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:17:34Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:17:39Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:17:44Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:17:48Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:17:53Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:17:58Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
2026-09-21T19:18:03Z: d2d3da9f-09c4-497f-83e5-1ade8e100bb2
```