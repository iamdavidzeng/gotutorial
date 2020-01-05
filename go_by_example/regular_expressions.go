package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	fmt.Println(r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+ch)")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

	sample := "{% macro products_list(products) %}{% for product in products %}productsList{% endfor %}{% endmacro %}"
	var re = regexp.MustCompile(`(^|[^_)])\bproducts\b([^_]|$)`)
	s := re.ReplaceAllString(sample, `$1.$2`)
	fmt.Println(s)

	count := 1
	search := regexp.MustCompile(`\$\{([^}:]+):?([^}]+)?\}`)
	body := []byte("DB_URIS:'payments:Base': mysql+mysqlconnector://${DB_USER:root}:${DB_PASS:}@${DB_SERVER:localhost}/${DB_NAME:payments}")

	body = search.ReplaceAllFunc(body, func(s []byte) []byte {
		fmt.Println(string(s))
		// How can I access the capture group here?
		s1 := search.ReplaceAllString(string(s), `$1`)
		s2 := search.ReplaceAllString(string(s), `$2`)
		envS1 := os.Getenv(s1)
		fmt.Println(string(s1), string(s2))
		count++
		if len(envS1) > 0 {
			return []byte(envS1)
		}
		return []byte(s2)
	})

	fmt.Println(string(body))
}
