package repository

import (
	"refresh-token/service/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRequest(t *testing.T) {
	assert := assert.New(t)
	body := `{"user_id": "abc@bc.com", "refresh_token": "eyJjdHkiOiJKV1QiLCJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiUlNBLU9BRVAifQ.Z7c_yVi_KfX03q_aNn-vpvLoM7SHD6dRFfnWWziNXPw6NOudY0jbJE_BL8LlUmAmKjvz0-bHTnCUYgmatvsPnObiH6eLzT_6mR-qp12QjmKH_2FuVxncGmIhgcTDAVB9Ib6yQSfSQv04o0XZt3aGXNR3CM5Bc32rRrggrkQIUdnPjSlIFkbldqZz8bXhS6GJDkLPU4iJBiLazSNIBt1dyWPl20ijmGpx56iMxLS4hf8OJrwIRLbKdT0aGUJSdpOC67VNp4deXiWvFMSepdNRidJjKwVzWDyqKw-ntH0e6PQFWTMezPE5tporP1rTnI9TjSDhvYdcB4Gr9hdv8EzVXg.7aBYaRg_GfUIPg3r.rBdcpl4Mrac230chpmmqd7LsTNbD42O93Ser6-SvGw1W9WarplTHBLimmFzwlfWoMYPrNN4rVS6c9Ro7NTdA4-yxN2zYPsR3eyLPqCiL0naLDn_doYjnXHV0qj1aAKf-Zts_KEpWB_IdOwilPoLBk3-xVquX-yGN4ktNu34xfe3Q1_x3MCKj7wR5LVGnGNc-XFxYqvO-q5ZBit5VcfRbdQvX5gW4rGZBzEXF8SUzYPFe_7WN54AYM5kwe-ocWJzndrwvSAco_QXatXDftiafa6iOsu5sVW9oreHgKeZKuNdTDamkM3fV3sxUT4E5FbrpZb8V3NQ8lZ6GNnDJiTv6YZVTDvq2OesEIb0b6PS1KrR3AWMc4XLbgCTrK21IWQ56wgVAfeoHrgpF0Bp17qRJz1Ugdnk7b-KliVebRzg9OwrZQEXv4qu-fnJXjvmu6qR7jIGsj_fSeHXeHEpuwVs0YaEqRxNd5CZxtitjVAEoxpSJZ4vQfIS-sgbZWQSLqY-gscFDSgY15_Y452xzm-OXv2rCwHpVDo7laEwSG-3kmh4Suz07MNeIxxn4jUMBNuoInO4rIEPAxmP4T62ST6XtgZcwfvR5z-7zM4iC02O3O4_uCXQ7OIU2Pe4WAIHe1OkS87Y_pfFtdLiE9K8wT7MGjo-CLXceucxWJjjKZYPgQ4bjW4zSo0YEP08RUvGOIjR26NWTvUebCC5KtEiAMPZfkX5HauAr-N1LW_M1D_UdGDM51b3plqje246T_YkKFHluY1WXK0oNDj1b0ANvGT6zhMO2ih7eXFeyE6np1muXuLLySeUOCiuiRFaoYEfaHAUH3d3KKAL7nUccn8eHT2TJOHwi1Ny6K0Iii4De3ULXktBIvLeteE8JjpCnsJiIo1OsDyEH29sFLcdH-RpIYCu6dlyh66L0upRAUqXY01g6XPDhgSDNYTaS-uzt4cagPRIwmeEU3DPIRoNKeMR-55rRfD80454aP0_CaSNwcn8WKngI50E8v5naPLCSbK14bBGr-l8lgSMtupcWDr5xY2055FVFPHXOdO1vlwgD-F31c7datp2mfHkjG5AiDSAUmh2di1dx_thv2KwDg1k_AaVMZEWtSBQJjYxxOlDtXTpGnMDD0r2rV-m7YeOrN6FPZhJBebOLtxSgoMmlzfrdAW3SFn30wb5DOZGSrwsdzDBAa8nSMIDlrIf--QJ009yvcpA8-vsrrNTZEYCuYkmoyRe8i2sCd0sMxzm98WqmGG1vgvLsOfU1nS98qBAltewgFdXUdvPhKo8KxB4auYa6WXNH-5grD3iY-HsdrDZuGmPA_Slc6mNTQTEjg4nU9w.w7M-O3LBw5urRhP4zcU1EA"}`

	requestParameter, err := ParseRequest(&body)

	assert.NoError(err)

	// assert for not nil (good when you expect something)
	if assert.NotNil(requestParameter) {
		assert.Equal(*requestParameter.UserId, "abc@bc.com")
		assert.Equal(*requestParameter.RefreshToken, "eyJjdHkiOiJKV1QiLCJlbmMiOiJBMjU2R0NNIiwiYWxnIjoiUlNBLU9BRVAifQ.Z7c_yVi_KfX03q_aNn-vpvLoM7SHD6dRFfnWWziNXPw6NOudY0jbJE_BL8LlUmAmKjvz0-bHTnCUYgmatvsPnObiH6eLzT_6mR-qp12QjmKH_2FuVxncGmIhgcTDAVB9Ib6yQSfSQv04o0XZt3aGXNR3CM5Bc32rRrggrkQIUdnPjSlIFkbldqZz8bXhS6GJDkLPU4iJBiLazSNIBt1dyWPl20ijmGpx56iMxLS4hf8OJrwIRLbKdT0aGUJSdpOC67VNp4deXiWvFMSepdNRidJjKwVzWDyqKw-ntH0e6PQFWTMezPE5tporP1rTnI9TjSDhvYdcB4Gr9hdv8EzVXg.7aBYaRg_GfUIPg3r.rBdcpl4Mrac230chpmmqd7LsTNbD42O93Ser6-SvGw1W9WarplTHBLimmFzwlfWoMYPrNN4rVS6c9Ro7NTdA4-yxN2zYPsR3eyLPqCiL0naLDn_doYjnXHV0qj1aAKf-Zts_KEpWB_IdOwilPoLBk3-xVquX-yGN4ktNu34xfe3Q1_x3MCKj7wR5LVGnGNc-XFxYqvO-q5ZBit5VcfRbdQvX5gW4rGZBzEXF8SUzYPFe_7WN54AYM5kwe-ocWJzndrwvSAco_QXatXDftiafa6iOsu5sVW9oreHgKeZKuNdTDamkM3fV3sxUT4E5FbrpZb8V3NQ8lZ6GNnDJiTv6YZVTDvq2OesEIb0b6PS1KrR3AWMc4XLbgCTrK21IWQ56wgVAfeoHrgpF0Bp17qRJz1Ugdnk7b-KliVebRzg9OwrZQEXv4qu-fnJXjvmu6qR7jIGsj_fSeHXeHEpuwVs0YaEqRxNd5CZxtitjVAEoxpSJZ4vQfIS-sgbZWQSLqY-gscFDSgY15_Y452xzm-OXv2rCwHpVDo7laEwSG-3kmh4Suz07MNeIxxn4jUMBNuoInO4rIEPAxmP4T62ST6XtgZcwfvR5z-7zM4iC02O3O4_uCXQ7OIU2Pe4WAIHe1OkS87Y_pfFtdLiE9K8wT7MGjo-CLXceucxWJjjKZYPgQ4bjW4zSo0YEP08RUvGOIjR26NWTvUebCC5KtEiAMPZfkX5HauAr-N1LW_M1D_UdGDM51b3plqje246T_YkKFHluY1WXK0oNDj1b0ANvGT6zhMO2ih7eXFeyE6np1muXuLLySeUOCiuiRFaoYEfaHAUH3d3KKAL7nUccn8eHT2TJOHwi1Ny6K0Iii4De3ULXktBIvLeteE8JjpCnsJiIo1OsDyEH29sFLcdH-RpIYCu6dlyh66L0upRAUqXY01g6XPDhgSDNYTaS-uzt4cagPRIwmeEU3DPIRoNKeMR-55rRfD80454aP0_CaSNwcn8WKngI50E8v5naPLCSbK14bBGr-l8lgSMtupcWDr5xY2055FVFPHXOdO1vlwgD-F31c7datp2mfHkjG5AiDSAUmh2di1dx_thv2KwDg1k_AaVMZEWtSBQJjYxxOlDtXTpGnMDD0r2rV-m7YeOrN6FPZhJBebOLtxSgoMmlzfrdAW3SFn30wb5DOZGSrwsdzDBAa8nSMIDlrIf--QJ009yvcpA8-vsrrNTZEYCuYkmoyRe8i2sCd0sMxzm98WqmGG1vgvLsOfU1nS98qBAltewgFdXUdvPhKo8KxB4auYa6WXNH-5grD3iY-HsdrDZuGmPA_Slc6mNTQTEjg4nU9w.w7M-O3LBw5urRhP4zcU1EA")
	}
}

func TestComputeHashMap(t *testing.T) {
	assert := assert.New(t)
	user_id := "abc@bc.com"
	config.ClientId = "clientID"
	config.ClientSecret = "clientSecret"

	hashedSecret := ComputeSecretHash(&user_id)

	assert.NotNil(hashedSecret)
	assert.Equal("aMIWMCPgXsRusULKA/OJ1AfztPK2JgBykLKnCZ1k/yQ=", hashedSecret)
}

func TestEmptyRefreshToken(t *testing.T) {
	assert := assert.New(t)
	body := `{"user_id": "abc@bc.com"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal(err.Error(), "the refresh token is empty")
	assert.Nil(requestParameter)
}

func TestEmptyUserId(t *testing.T) {
	assert := assert.New(t)
	body := `{"refresh_token": "123456787632145678919"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Equal(err.Error(), "the user id is empty")
	assert.Nil(requestParameter)
}

func TestInvalidRefreshToken(t *testing.T) {
	assert := assert.New(t)
	body := `{"user_id": "abc@bc.com", "refresh_token": "225"}`

	requestParameter, err := ParseRequest(&body)

	assert.Error(err)
	assert.Nil(requestParameter)
	assert.Equal(err.Error(), "the refresh token entered is invalid")
}
