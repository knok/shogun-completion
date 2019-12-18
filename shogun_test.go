package main

import (
	"bytes"
	"testing"
)

func TestShogun(t *testing.T) {
	tests := []struct {
		args []string
		out  string
		err  string
		code int
	}{
		{
			args: []string{},
			out:  "",
			err:  "",
			code: 1,
		},
		{
			args: []string{"foo"},
			out:  "",
			err:  "そんな将軍はいない: foo\n",
			code: 2,
		},
		{
			args: []string{"foo", "bar"},
			out:  "",
			err:  "",
			code: 1,
		},
		{
			args: []string{"徳川綱吉"},
			out:  "徳川綱吉 (the 5th Shogun)\n",
			err:  "",
			code: 0,
		},
	}

	for _, test := range tests {
		var outbuf, errbuf bytes.Buffer
		stdout = &outbuf
		stderr = &errbuf
		code := run(test.args)
		if code != test.code {
			t.Fatalf("want %v for code, but %v:", test.code, code)
		}

		outstr := outbuf.String()
		if outstr != test.out {
			t.Fatalf("out: want %v, but %v:", test.out, outstr)
		}
		errstr := errbuf.String()
		if errstr != test.err {
			t.Fatalf("err: want %v, but %v:", test.err, errstr)
		}
	}
}
