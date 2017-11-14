package tls

import (
	"github.com/elastic/beats/packetbeat/config"
	"github.com/elastic/beats/packetbeat/protos"
)

type tlsConfig struct {
	config.ProtocolCommon `config:",inline"`
	SendCertificates      bool `config:"send_certificates"`
}

var (
	defaultConfig = tlsConfig{
		ProtocolCommon: config.ProtocolCommon{
			TransactionTimeout: protos.DefaultTransactionExpiration,
		},
		SendCertificates: true,
	}
)
