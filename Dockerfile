FROM ubuntu:20.04

WORKDIR /bittrace

COPY ./output/openapi /bittrace/

VOLUME ["/root/.bittrace"]

# for openapi-cli
EXPOSE 6060/tcp

ENTRYPOINT ["/bittrace/openapi"]
