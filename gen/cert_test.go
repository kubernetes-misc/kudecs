package gen

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
	"time"
)

func TestGenerateCert(t *testing.T) {
	private, public := GenerateCert(GenerateRequest{
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(7 * 24 * time.Hour),
	})

	assert.NotEmpty(t, private)
	assert.NotEmpty(t, public)

	ioutil.WriteFile("private.key", private, 0644)
	ioutil.WriteFile("public.pem", public, 0644)

}
