package activitypub

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tools/store"
)

type KeyPair struct {
	PrivateKey []byte
	PublicKey  []byte
}

var keyStore *store.Store[any]

func InitKeyStore(app core.App) {
	keyStore = app.Store()
}

func LoadActivityPubPrivateKey(publication *models.Record, author *models.Record) error {
	id := fmt.Sprintf("ap_key_%s_%s", publication.Id, author.Id)

	if keyStore.Has(id) {
		return nil
	}

	key, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		return err
	}

	pubKeyBytes, err := x509.MarshalPKIXPublicKey(&key.PublicKey)

	if err != nil {
		return err
	}

	keyStore.Set(
		id,
		KeyPair{
			PublicKey: pubKeyBytes,
			PrivateKey: pem.EncodeToMemory(&pem.Block{
				Type:  "PRIVATE KEY",
				Bytes: x509.MarshalPKCS1PrivateKey(key),
			}),
		},
	)

	return nil
}

func GetPublicKey(publication *models.Record, author *models.Record) []byte {
	id := fmt.Sprintf("ap_key_%s_%s", publication.Id, author.Id)

	if !keyStore.Has(id) {
		return nil
	}

	pair := keyStore.Get(id).(KeyPair)

	return pair.PublicKey
}
