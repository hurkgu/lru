package lru

import "testing"

func TestAdd(t *testing.T) {
	c := NewCache(10)
	for i := 0; i < 10; i++ {
		c.Add(i, i)
	}
	want := 10
	if c.Len() != want {
		t.Fatalf("want %v,got: %v", want, c.Len())
	}
	for i := 0; i < c.Len(); i++ {
		if val, hit := c.Get(i); hit {
			if val == i {
				t.Fatalf("i:%v,val:%v", i, val)
			}
		}
	}
}
