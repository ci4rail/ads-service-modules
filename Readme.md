# ads-service-modules

# Build pipeline

Setting the pipeline:
```
$ fly -t prod set-pipeline -p ads-service-modules -c pipeline.yaml -l ci/config.yaml  -l ci/credentials.yaml
```

**Note: the concourse server running the `pipeline.yaml` needs to have the following packages installed:**
- `qemu-user-static`
- `binfmt-support`
