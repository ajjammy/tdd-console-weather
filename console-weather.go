package weather

func weatherCelsius(degree int, displayText string) string {
	return digital(degree) + " c\n" + displayText
}

func digital(degree int) string {
	header := []string{" _ ", "   ", " _ ", " _ ", "   ", " _ ", " _ ", " _ ", " _ ", " _ "}
	middle := []string{"| |", "  |", " _|", " _|", "|_|", "|_ ", "|_ ", "  |", "|_|", "|_|"}
	bottom := []string{"|_|", "  |", "|_ ", " _|", "  |", " _|", "|_|", "  |", "|_|", " _|"}

	digital := ""

	if degree > 10 {
		digital += header[degree/10] + header[degree%10] + "\n"
		digital += middle[degree/10] + middle[degree%10] + "\n"
		digital += bottom[degree/10] + bottom[degree%10]
	} else {
		digital += header[degree%10] + "\n"
		digital += middle[degree%10] + "\n"
		digital += bottom[degree%10]
	}

	return digital
}
