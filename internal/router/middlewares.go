package router

import (
	"dex-trades-parser/internal/models"
	"dex-trades-parser/internal/storage"
	"dex-trades-parser/pkg/helpers"
	"dex-trades-parser/pkg/response"
	"encoding/base64"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func СheckAuthSign(st *storage.Storage) gin.HandlerFunc {
	return func(c *gin.Context) {
		authRawHeader := c.Request.Header.Get("X-ELEKTRA")
		if authRawHeader == "" {
			response.Error(c, http.StatusForbidden, response.E{
				Code:    response.Unauthorized,
				Message:"No Authorization header provided",
			})
			return
		}

		decodedHeaderString, err := base64.StdEncoding.DecodeString(authRawHeader)
		if err != nil {
			response.Error(c, http.StatusForbidden, response.E{
				Code:    response.Unauthorized,
				Message:"Wrong Authorization header value provided",
			})
			return
		}

		parsedHeader := strings.Split(string(decodedHeaderString), ":")
		if helpers.IsValidAddress(parsedHeader[0]) == false {
			response.Error(c, http.StatusBadRequest, response.E{
				Code:    response.InvalidJSONBody,
				Message: "invalid wallet address",
			})
			return
		}

		walletAdr := strings.ToLower(parsedHeader[0])
		signature := parsedHeader[1]

		var nonce models.Nonce
		if err := st.DB.First(&nonce,
			&models.Nonce{
				Wallet: walletAdr,
			}).Error; err != nil {
			response.Error(c, http.StatusForbidden, response.E{
				Code:    response.Unauthorized,
				Message: "Nonce not found",
			})
			return
		}

		message := []byte(fmt.Sprintf("I am signing my one-time nonce: %d", nonce.Nonce))

		recoveredAddress, err := verifySig(signature, message)
		if err != nil {
			response.Error(c, http.StatusForbidden, response.E{
				Code:    response.Unauthorized,
				Message:"Wrong Sign Recover",
			})
			return
		}

		lowerRecoveredAddress := strings.ToLower(recoveredAddress.String())
		if walletAdr != lowerRecoveredAddress {
			response.Error(c, http.StatusUnauthorized, response.E{
				Code:    response.Unauthorized,
				Message: "invalid auth signature",
			})
			return
		}

		nonce.Nonce = helpers.RandomNonce()
		st.DB.Save(&nonce)
	}
}

func verifySig(sigHex string, msg []byte) (common.Address, error) {
	sig := hexutil.MustDecode(sigHex)

	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}, fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
	}
	sig[64] -= 27

	pubKey, err := crypto.SigToPub(signHash(msg), sig)
	if err != nil {
		return common.Address{}, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	return recoveredAddr, nil
}

// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L404
// signHash is a helper function that calculates a hash for the given message that can be
// safely used to calculate a signature from.
//
// The hash is calculated as
//   keccak256("\x19Ethereum Signed Message:\n"${message length}${message}).
//
// This gives context to the signed message and prevents signing of transactions.
func signHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}
