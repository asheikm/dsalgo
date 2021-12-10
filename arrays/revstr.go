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
	fmt.Println(reverseString(s))
}

func reverseString(s string) string {
	a := []byte(s)
	for i, j := 0, len(s)-1; i < j; i++ {
		a[i], a[j] = a[j], a[i]
		j--
	}
	return string(a)
}
