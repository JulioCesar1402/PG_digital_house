/*
! Need to be refac:

var 1nome string 				   -- errada
var sobrenome string               -- certa
var int idade                      -- errada
1sobrenome := 6                    -- errada
var licenca_para_dirigir = true    -- errada
var estatura da pessoa int         -- errada
quantidadeDeFilhos := 2            -- certa, mas depende

*/

package evaluation

import "fmt"

var firstName string
var lastName string
var age uint8

/*
firstLastName := "Cesar", porem só funciona dentro de uma func
caso esteja fora:
lastName = "Cesar"
*/
var licencaParaDirigir = true
var estaturaDaPessoa float32

/*
quantidadeDeFilhos := 2, porem só funciona dentro de uma func
caso esteja fora:
var quantidadeDeFilhos = 2
*/

func Exame() {
	fmt.Println(
		firstName,
		lastName,
		age,
		licencaParaDirigir,
		estaturaDaPessoa,
	)
}
