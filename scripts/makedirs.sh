#!/usr/bin/env bash

cd ..

mkdir -p api/{proto/tiny-auth-service,rest} \
    cmd/tiny-auth-service \
    configs \
    deployments \
    docs \
    internal/{app,config,domain,facade/{dto,mapper},repository/postgres,transport/{errs,rest,grpc},usecase} \
    migrations/tiny-auth-service \
    pkg \
    scripts

cd scripts
