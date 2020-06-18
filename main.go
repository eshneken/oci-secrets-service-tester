package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/common/auth"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/secrets"
)

func main() {
	secretID := "ocid1.vaultsecret.oc1.iad.amaaaaaabxdvnfaazdhrjrhy4fzedkffoqokwp6zl44doitxs5ora5k6dsia"

	var provider common.ConfigurationProvider
	provider, err := auth.InstancePrincipalConfigurationProvider()
	if err != nil {
		provider = common.DefaultConfigProvider()
	}

	client, err := secrets.NewSecretsClientWithConfigurationProvider(provider)
	helpers.FatalIfError(err)

	request := secrets.GetSecretBundleRequest{SecretId: &secretID}
	response, err := client.GetSecretBundle(context.Background(), request)
	helpers.FatalIfError(err)

	encodedResponse := fmt.Sprintf("%s", response.SecretBundleContent)
	encodedResponse = strings.TrimRight(strings.TrimLeft(encodedResponse, "{ Content="), " }")
	decodedByteArray, _ := base64.StdEncoding.DecodeString(encodedResponse)
	fmt.Println(string(decodedByteArray))
}
