package jwt

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/patrickmn/go-cache"
	"github.com/square/go-jose"
)

type JwtValidator struct {
	cache *cache.Cache

	aud string
	iss string
}

func NewJwtValidator(audience, issuer string) *JwtValidator {
	return &JwtValidator{
		cache: cache.New(3*time.Minute, 30*time.Second),

		aud: audience,
		iss: issuer,
	}
}

func (validator *JwtValidator) PublicKey(token *jwt.Token) (*rsa.PublicKey, error) {
	// TODO: The following address should be public like https://www.googleapis.com/oauth2/v3/certs.
	// TODO: The following address should be configurable.
	resp, err := http.Get("http://oathkeeper-api.oathkeeper:4456/.well-known/jwks.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	jwks := &jose.JSONWebKeySet{}
	if err = json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, err
	}

	kid, ok := token.Header["kid"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid type [%T] for header 'kid', expected 'string'", token.Header["kid"])
	}

	keys := jwks.Key(kid)
	if len(keys) == 0 {
		return nil, fmt.Errorf("could not find the 'kid' [%s]", kid)
	}

	publicKey, ok := keys[0].Key.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("invalid type [%T] for JWKS key [%s], expected '*rsa.PublicKey'", keys[0].Key, kid)
	}

	return publicKey, nil
}

func (validator *JwtValidator) ValidationKeyGetter(token *jwt.Token) (interface{}, error) {
	claims := token.Claims.(jwt.MapClaims)
	if err := claims.Valid(); err != nil {
		return nil, err
	}

	// TODO: If additional scopes / claims should be checked implement like in
	// TODO: https://auth0.com/docs/quickstart/backend/golang/01-authorization

	// ! https://github.com/auth0/go-jwt-middleware/issues/72
	// ! https://github.com/form3tech-oss/jwt-go/issues/7
	// ! https://github.com/form3tech-oss/jwt-go/issues/5

	if audienceI, ok := claims["aud"].([]interface{}); ok {
		var audience []string
		for _, i := range audienceI {
			value, ok := i.(string)
			if !ok {
				return nil, fmt.Errorf("invalid audience type [%T], expected '[]interface{}' or '[]string'", claims["aud"])
			}

			audience = append(audience, value)
		}

		token.Claims.(jwt.MapClaims)["aud"] = audience
	}

	if !claims.VerifyAudience(validator.aud, true) {
		return nil, fmt.Errorf("invalid audience [%s], expected '%s'", claims["aud"], validator.aud)
	}

	if !claims.VerifyIssuer(validator.iss, true) {
		return nil, fmt.Errorf("invalid issuer [%s], expected '%s'", claims["iss"], validator.iss)
	}

	publicKey, ok := validator.cache.Get("jwks")
	if ok {
		return publicKey, nil
	}

	publicKey, err := validator.PublicKey(token)
	if err != nil {
		return nil, err
	}

	validator.cache.Set("jwks", publicKey, cache.DefaultExpiration)
	return publicKey, nil
}
