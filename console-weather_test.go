package weather

import (
	"testing"
)

// fmt.Println(weatherCelsius(25, "Bangkok few cloud"))
// fmt.Println(weatherCelsius(34, "Tak sunny"))
// fmt.Println(weatherCelsius(17, "Phuket rainy"))
// fmt.Println(weatherCelsius(9, "Chiang-mai cold"))

func Test_input_degree_9_and_Chiang_mai_cold_display_digital_9_and_text_Chiang_mai_cold(t *testing.T) {
	expectedDigital := " _ \n"
	expectedDigital += "|_|\n"
	expectedDigital += " _| c\n"
	expectedDigital += "Chiang-mai cold"

	if weatherCelsius(9, "Chiang-mai cold") != expectedDigital {
		t.Error("Actual is \n" + weatherCelsius(9, "Chiang-mai cold"))
	}
}

func Test_input_degree_17_and_Phuket_rainy_display_digital_17_and_text_Phuket_rainy(t *testing.T) {
	expectedDigital := "    _ \n"
	expectedDigital += "  |  |\n"
	expectedDigital += "  |  | c\n"
	expectedDigital += "Phuket rainy"

	if weatherCelsius(17, "Phuket rainy") != expectedDigital {
		t.Error("Actual is \n" + weatherCelsius(17, "Phuket rainy"))
	}
}

func Test_input_degree_34_and_Tak_sunny_display_digital_34_and_text_Tak_sunny(t *testing.T) {
	expectedDigital := " _    \n"
	expectedDigital += " _||_|\n"
	expectedDigital += " _|  | c\n"
	expectedDigital += "Tak sunny"

	if weatherCelsius(34, "Tak sunny") != expectedDigital {
		t.Error("Actual is \n" + weatherCelsius(34, "Tak sunny"))
	}
}

func Test_input_degree_25_and_Bangkok_few_cloud_display_digital_25_and_text_Bangkok_few_cloud(t *testing.T) {
	expectedDigital := " _  _ \n"
	expectedDigital += " _||_ \n"
	expectedDigital += "|_  _| c\n"
	expectedDigital += "Bangkok few cloud"

	if weatherCelsius(25, "Bangkok few cloud") != expectedDigital {
		t.Error("Actual is \n" + weatherCelsius(25, "Bangkok few cloud"))
	}
}

func Test_input_degree_25_display_digital_25(t *testing.T) {
	expectedDigital := " _  _ \n"
	expectedDigital += " _||_ \n"
	expectedDigital += "|_  _|"

	if digital(25) != expectedDigital {
		t.Error("Actual is \n" + digital(25))
	}
}

func Test_input_degree_34_display_digital_34(t *testing.T) {
	expectedDigital := " _    \n"
	expectedDigital += " _||_|\n"
	expectedDigital += " _|  |"

	if digital(34) != expectedDigital {
		t.Error("Actual is \n" + digital(34))
	}
}

func Test_input_degree_17_display_digital_17(t *testing.T) {
	expectedDigital := "    _ \n"
	expectedDigital += "  |  |\n"
	expectedDigital += "  |  |"

	if digital(17) != expectedDigital {
		t.Error("Actual is \n" + digital(17))
	}
}

func Test_input_degree_2_display_digital_2(t *testing.T) {
	expectedDigital := " _ \n"
	expectedDigital += " _|\n"
	expectedDigital += "|_ "

	if digital(2) != expectedDigital {
		t.Error("Actual is \n" + digital(2))
	}
}

func Test_input_degree_5_display_digital_5(t *testing.T) {
	expectedDigital := " _ \n"
	expectedDigital += "|_ \n"
	expectedDigital += " _|"

	if digital(5) != expectedDigital {
		t.Error("Actual is \n" + digital(5))
	}
}

func Test_input_degree_9_display_digital_9(t *testing.T) {
	expectedDigital := " _ \n"
	expectedDigital += "|_|\n"
	expectedDigital += " _|"

	if digital(9) != expectedDigital {
		t.Error("Actual is \n" + digital(9))
	}
}
