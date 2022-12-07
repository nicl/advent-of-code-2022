package main

import "testing"

func TestAdd(t *testing.T) {
	node := Node{
		name: "/",
		children: []*Node{
			{name: "foo.txt", size: 123},
			{name: "bar.txt", size: 345},
			{name: "tmp"},
		},
	}

	want := &Node{name: "baz.txt", size: 456}
	node.Add("/tmp", want)
	got, ok := node.Find("/tmp/baz.txt")
	if !ok {
		t.Errorf("unable to find /tmp/baz.txt")
		return
	}

	if got != want {
		t.Errorf("got %v; want %v", *got, *want)
	}

}
