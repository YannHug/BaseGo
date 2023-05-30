package main

import (
	"fmt"
	"gotest/boucle"
	"gotest/condition"
	"gotest/conversion"
	"gotest/erreur"
	"gotest/fonction"
	"gotest/forEach"
	"gotest/lesStruct"
	"gotest/pointeurReceiver"
	"gotest/pointeurs"
	"gotest/slice"
	"gotest/tableau"
	"gotest/variable"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("Hello World !"))

	variable.Variable()
	condition.Condition()
	conversion.Conversion()
	fonction.Main()
	tableau.Main()
	slice.Main()
	boucle.Main()
	forEach.Main()
	erreur.Main()
	lesStruct.Main()
	pointeurs.Main()
	pointeurReceiver.Main()

}
