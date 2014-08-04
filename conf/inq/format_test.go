package inq

import (
	"github.com/stretchr/testify/assert"
	"inspeqtor/conf/inq/ast"
	"inspeqtor/conf/inq/lexer"
	"inspeqtor/conf/inq/parser"
	"io/ioutil"
	"testing"
)

func TestBasicServiceParsing(t *testing.T) {
	data, err := ioutil.ReadFile("conf.d/memcache.inq")
	if err != nil {
		t.Fatal(err)
	}

	s := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	obj, err := p.Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	check := obj.(*ast.ProcessCheck)
	assert.Equal(t, check.Parameters["key"], "value")
	assert.Equal(t, check.Parameters["foo"], "bar")
}

func TestBasicHostParsing(t *testing.T) {
	data, err := ioutil.ReadFile("conf.d/system.inq")
	if err != nil {
		t.Fatal(err)
	}

	s := lexer.NewLexer([]byte(data))
	p := parser.NewParser()
	obj, err := p.Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	check := obj.(*ast.HostCheck)
	assert.Equal(t, len(check.Rules), 2)
}
