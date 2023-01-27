package autenticacao

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {

	permissoes := jwt.MapClaims{

	}
	permissoes["autorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()//expiração data de agora + 6h
	permissoes["usuarioID"] = usuarioID
	//metodo de assinatura do token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	//assianrr o token
	return token.SignedString([]byte(config.SecretKey)) //secretkey - fazer assinatura e garantir autenticidade do token

}
