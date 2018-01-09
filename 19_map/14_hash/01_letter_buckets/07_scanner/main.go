package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// sluzi za input
	const tekstic = "It is not the critic who counts; not the man who points out how the strong man stumbles, or where the doer of deeds could have done them better. The credit belongs to the man who is actually in the arena, whose face is marred by dust and sweat and blood; who strives valiantly; who errs, who comes short again and again, because there is no effort without error and shortcoming; but who does actually strive to do the deeds; who knows great enthusiasms, the great devotions; who spends himself in a worthy cause; who at the best knows in the end the triumph of high achievement, and who at the worst, if he fails, at least fails while daring greatly, so that his place shall never be with those cold and timid souls who neither know victory nor defeat. "

	//NewScanner uzima Reader - Newreader uzima string i vraca reader
	//vraca pokazivac na scanner
	scanner := bufio.NewScanner(strings.NewReader(tekstic))

	// podesavamo split funciju da bi mogli da skeniramo reci.
	// ScanWords je split funkcija za Scanner koja vraca
	// svaku rec odvojenu space-om i brise sve space-ove oko nje.
	scanner.Split(bufio.ScanWords)

	// Prebroj reci
	//Scan() vraca false kada se skeniranje zavrsi
	//Err metoda vraca bilo koju gresku koja se pojavila kod skeniranja
	// osim ako je u pitanju io.EOF, tad Err vraca nil
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
