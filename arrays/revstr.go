package main

import "fmt"

/*
Posted this to a stackoveflow string reverse question although it has 30+ answers already, intrestingly no one tried using byte array approach

Strings are immutable object in golang, unlike C inplace reverse is not possible with golang. With C , you can do something like,

void reverseString(char *str) {
  int length = strlen(str)
  for(int i = 0, j = length-1; i < length/2; i++, j--)
  {
    char tmp = str[i];
    str[i] = str[j];
    str[j] = tmp;
  }
}
But with golang, following one, uses byte to convert the input into bytes first and then reverses the byte array once it is reversed,
convert back to string before returning. works only with non unicode type string.
*/
func main() {
	s := "test123 4"
	fmt.Println(reverseString1(s))
	fmt.Println(reverseString2(s))
}

func reverseString1(s string) string {
	b := []byte(s)
	for i, j := 0, len(s)-1; i < j; i++ {
		b[i], b[j] = b[j], b[i]
		j--
	}
	return string(b)
}

/* String Reverse using rune *
   Rune is to hanle not only chars but also unicode chars
   rune is considered int32 equavalent */
func reverseString2(s string) string {
	r := []rune(s) // convert string to rune
	for i, j := 0, len(s)-1; i < j; i++ {
		r[i], r[j] = r[j], r[i] // Swap
		j--
	}
	return string(r) // convert back to string

}
