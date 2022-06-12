package cache

import "time"

type Cache struct {
	m map[string]string
}

func NewCache() Cache {
	return Cache{
		m: make(map[string]string),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	v, ok := c.m[key]
	return v, ok

}

func (c *Cache) Put(key, value string) {
	c.m[key] = value
}

func (c *Cache) Keys() []string {
	var ss []string
	for k, _ := range c.m {
		ss = append(ss, k)
	}
	return ss
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.m[key] = value
	//c.PutTill("s", "v", time.Now())

	go c.Delete(key, deadline)
}

func (c *Cache) Delete(k string, deadline time.Time) {
	d := deadline.Sub(time.Now())
	time.Sleep(d)
	delete(c.m, k)
}
