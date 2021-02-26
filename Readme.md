# ads-service-modules

# Local build with dobi

Make sure that you manually run `docker login` for user `ci4rail` on your host system. The `~/.docker/config.json` gets mounted for the build steps in order to push the docker images.

# Build pipeline

Setting the pipeline:
```
$ fly -t prod set-pipeline -p ads-service-modules -c pipeline.yaml -l ci/config.yaml  -l ci/credentials.yaml
```

# Build misc informatin

**Note: the host (either locally or concourse CI server) needs to have the following packages installed:**
- `qemu-user-static`
- `binfmt-support`
