package usables

//
// import (
// 	"context"
// 	"errors"
// 	"net/http"
// )
//
// type DatosContexto struct {
// 	idsesion int64
// }
//
// func ObtenerContexto(solicitud http.Request) (context.Context, error) {
// 	var cookies []*http.Cookie
// 	var idsesion string = ""
// 	var datos DatosContexto
// 	var contexto context.Context = context.Background()
//
// 	cookies = solicitud.Cookies()
// 	for contador := len(cookies); contador >= 0; contador-- {
// 		if cookies[contador].Name == "idsesion" {
// 			idsesion = cookies[contador].Value
// 			break
// 		}
// 	}
// 	if idsesion == "" {
// 		return contexto, errors.New("servidor/src/necesarios: no se encontró la cookie de sesión")
// 	}
//
// 	datos = DatosContexto{LeerToken(idsesion)}
// 	//contexto = *solicitud.WithContext(contexto)
// 	contexto = context.WithValue(contexto, 0, datos)
// 	return contexto, nil
// }
