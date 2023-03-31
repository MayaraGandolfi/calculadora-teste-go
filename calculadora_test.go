package main

import "testing"

func ShouldSumCorrect(t *testing.T){
	//arrange
	teste := soma(3,2,1)

	//act
	resultado := 5

	//assert
	if teste == resultado{
		t.Error("Valor esperado:",resultado,"valor retornado:",teste)
	}
}


func ShouldSubtractCorrect(t *testing.T){
	//arrange
	teste := subtrai(30,10)

	//act
	resultado := 20

	//assert
	if teste != resultado{
		t.Error("Valor esperado:",resultado,"valor retornado:",teste)
	}
}

func ShouldMultiplyCorrect(t *testing.T){
	//arrange
	teste := multiplica(4,4,4)

	//act
	resultado := 64

	//assert
	if teste != resultado{
		t.Error("Valor esperado:",resultado,"valor retornado:",teste)
	}
}

func ShouldDivideCorrect(t *testing.T){
	//arrange
	teste := divide(100,4)

	//act
	resultado := 20

	//assert
	if teste != resultado{
		t.Error("Valor esperado:",resultado,"valor retornado:",teste)
	}
}