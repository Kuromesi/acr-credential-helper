# acr-credential-provider
A credential provider for Alibaba Cloud Container Registry (ACR)

This repository is developed for OpenKruise exec plugin based on https://github.com/mozillazg/docker-credential-acr-helper/tree/master.

## Configuration

### ACR Credentials

By default, the helper searches for ACR credentials in the following order:

1. It fetches the credentials via [RAM Roles for Service Accounts (RRSA) OIDC Token](https://www.alibabacloud.com/help/en/container-service-for-kubernetes/latest/use-rrsa-to-enforce-access-control)
   when the `ALIBABA_CLOUD_ROLE_ARN`, `ALIBABA_CLOUD_OIDC_PROVIDER_ARN`, and
   `ALIBABA_CLOUD_OIDC_TOKEN_FILE` environment variables are defined and are not empty.
2. Use access key id and access key secret that are specified by the `ALIBABA_CLOUD_ACCESS_KEY_ID` and
   `ALIBABA_CLOUD_ACCESS_KEY_SECRET` environment variables.
3. A profile file whose path is specified by the `ALIBABA_CLOUD_CREDENTIALS_FILE` environment variable.
4. A profile file in a default location:
   * On Windows, this is `C:\Users\USER_NAME\.alibabacloud\credentials`.
   * On other systems, it is `~/.alibabacloud/credentials`.
5. It fetches the credentials of the RAM Role associated with the VM from the metadata server when
   the environment variable `ALIBABA_CLOUD_ECS_METADATA` is defined and not empty.

For more information about configuring credentials, see [Provider](https://github.com/aliyun/credentials-go#provider)
in the @aliyun/credentials-go.

## Usage
Run `go build cmd/acr-credential-provider/main.go`, then copy the binary file to `pluginBinDir` defined in Kruise and rename the file as defined in `pluginConfigFile`.