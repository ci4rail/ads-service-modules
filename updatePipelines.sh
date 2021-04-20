#!/bin/bash
PROJECT=ads-service-modules
fly -t prod set-pipeline -p ${PROJECT} -c pipeline.yaml -l ci/config.yaml  -l ci/credentials.yaml
fly -t prod set-pipeline -p ${PROJECT}-pull-requests -c pipeline-pullrequests.yaml -l ci/config.yaml  -l ci/credentials-pullrequests.yaml
