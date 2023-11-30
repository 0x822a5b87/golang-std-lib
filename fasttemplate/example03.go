package main

import (
	"fmt"
	"github.com/valyala/fasttemplate"
	"io"
	"net/url"
)

func example03() {
	template := "https://{{host}}/?foo={{bar}}{{bar}}&q={{query}}&baz={{baz}}"
	t := fasttemplate.New(template, "{{", "}}")

	// Substitution map.
	// Since "baz" tag is missing in the map, it will be substituted
	// by an empty string.
	m := map[string]interface{}{
		"host": "google.com",     // string - convenient
		"bar":  []byte("foobar"), // byte slice - the fastest

		// TagFunc - flexible value. TagFunc is called only if the given
		// tag exists in the template.
		"query": fasttemplate.TagFunc(func(w io.Writer, tag string) (int, error) {
			return w.Write([]byte(url.QueryEscape(tag + "=world")))
		}),
		"baz": "BAZ",
	}

	s := t.ExecuteString(m)
	fmt.Printf("%s", s)
}
