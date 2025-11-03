This branch is a copy of TRex with the following changes
- Use OpenAPI schema from https://gitlab.cee.redhat.com/amarin/hyperfleet-openapi-typespec#
  - Run `make generate` to generate openapi.* types
- Introduce some api.* types:
  - Cluster
  - ClusterStatus
  - Condition
- Introduce Cluster handler with POST and GET operations
- Make the project compile and pass tests
- Remove dinosaur references

The focus is to make the OpenAPI schema compatible with TRex, like paging or objectReference types
