// package conceptos

// "errors"

// func divide(dividendo, divisor int) (int, error) {
// 	if divisor == 0 {
// 		return 0, fmt.Errorf("No se puede dividir por cero")
// 	}
// 	return dividendo / divisor, nil
// }

// func divide(dividendo, divisor int) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println(r)
// 		}
// 	}()
// 	validateZero(divisor)
// 	fmt.Println(dividendo / divisor)
// }

// func validateZero(divisor int) {
// 	if divisor == 0 {
// 		panic("No se puede dividir por cero")
// 	}
// }

func conceptos() {
	//! Manejo de errores
	// result, err := divide(10, 0)

	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }

	// fmt.Println("Resultado: ", result)

	//!Defer- pospone la ejecucion de esa funcion hasta antes de cerrar el flujo
	// file, err := os.Create("Hola.text")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// defer file.Close()

	// _, err = file.Write([]byte("Hola Alex, estoy tomando el curso go"))

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	//!panic y recover
	//Panic-generar una interrupcion inmediata
	//Panic y recover solo se recomienda usar para casos excepcionales o errores graves, no se recomienda
	// para control de errores normal, solo para situaciones inesperadas

	// divide(100, 10)
	// divide(200, 25)
	// divide(34, 0)
	// divide(100, 4)

	//!Registro de errores
	// log.SetPrefix("main(): ")
	// file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// log.SetOutput(file)
	// log.Print("!Oye, soy un log")
}
