package substitution

import (
	"os"
)

const (
	secretPathEnv     = "CRUMBLECOG_PATH"
	defaultSecretPath = "crumblecog"
)

type SecretProvider int

const (
	Vault SecretProvider = iota
	Bitwarden
	Unknown
)

func secretProvider() SecretProvider {
	if _, ok := os.LookupEnv(`BW_SESSION`); ok {
		return Bitwarden
	}
	if _, ok := os.LookupEnv(`VAULT_ADDR`); ok {
		return Vault
	}
	return Unknown
}

func configSecretPath() string {
	if val, ok := os.LookupEnv(secretPathEnv); ok {
		return val
	}
	prefix := ``
	if secretProvider() == Vault {
		prefix = `secret/data/`
	}
	return prefix + defaultSecretPath
}
