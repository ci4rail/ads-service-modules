# ads-service-modules

[![CI](https://concourse.ci4rail.com/api/v1/teams/edgefarm/pipelines/ads-service-modules/jobs/build-ads-service-modules/badge)](https://concourse.ci4rail.com/teams/edgefarm/pipelines/ads-service-modules) [![Go Report Card](https://goreportcard.com/badge/github.com/ci4rail/ads-service-modules)](https://goreportcard.com/badge/github.com/ci4rail/ads-service-modules)

# Local build with dobi

Make sure that you manually run `docker login` for user `ci4rail` on your host system. The `~/.docker/config.json` gets mounted for the build steps in order to push the docker images.

# Build pipeline

Make sure that you are logged in to team `edgefarm` in Concourse.
```bash
$ fly -t prod login -n edgefarm
```

## pipeline.yaml

The `pipeline.yaml` is the CI/CD pipeline that builds all the ads service module docker images for different architectures. The images are published on [docker hub](https://hub.docker.com/u/ci4rail) and [Ci4Rail harbor](https://harbor.ci4rail.com/harbor/projects/7/repositories).

### Usage

Copy `ci/credentials.template.yaml` to `ci/credentials.yaml` and copy  the content from `ads-service-modules credentials.yaml` found in bitwarden.
Apply the CI/CD pipeline to Concourse CI using
```bash
$ fly -t prod set-pipeline -p ads-service-modules -c pipeline.yaml -l ci/config.yaml  -l ci/credentials.yaml
```

## pipeline-pullrequests.yaml

The `pipeline-pullrequests.yaml` defines a pipeline that runs basic quality checks on pull requests. For this, consourse checks Github for new or changed pull requests. If a change is found, it downloads the branch and builds all location modules. Those are pushed to `ci4rail-dev` on [harbor](https://harbor.ci4rail.com/harbor/projects/14/repositories).

### Usage

Copy `ci/credentials-pullrequests.template.yaml` to `ci/credentials-pullrequests.yaml` and copy  the content from `ads-service-modules credentials-pullrequests.yaml` found in bitwarden.
Configure a Webhook on github using this URL and the same webhook_token:
`https://concourse.ci4rail.com/api/v1/teams/edgefarm/pipelines/ads-service-modules-pull-requests/resources/pull-request/check/webhook?webhook_token=<webhook_token>`

Apply the pipeline with the name `ads-service-modules-pull-requests`
```bash
$ fly -t prod set-pipeline -p ads-service-modules-pull-requests -c pipeline-pullrequests.yaml -l ci/credentials-pullrequests.yaml
```

# Build misc information

**Note: the host (either locally or concourse CI server) needs to have the following packages installed:**
- `qemu-user-static`
- `binfmt-support`
