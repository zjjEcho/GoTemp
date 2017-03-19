package main

import (
	"fmt"
	"github.com/cupcake/rdb"
	"github.com/cupcake/rdb/nopdecoder"
	"os"
)

type Decoder struct {
	m MemProfiler
	nopdecoder.NopDecoder
}

// Set is called once for each string key.
func (d *Decoder) Set(key, value []byte, expiry int64) {
	keyStr := string(key)
	bytes := d.m.SizeofString(key)
	bytes += d.m.SizeofString(value)
	bytes += d.m.TopLevelObjOverhead()
	bytes += 2 * d.m.RobjOverhead()
	bytes += d.m.KeyExpiryOverhead(expiry)
	fmt.Printf("%q -> %d\n", keyStr, bytes)

}

// // StartHash is called at the beginning of a hash.
// // Hset will be called exactly length times before EndHash.
// func (d *Decoder) StartHash(key []byte, length, expiry int64) {

// }

// // Hset is called once for each field=value pair in a hash.
// func (d *Decoder) Hset(key, field, value []byte) {

// }

// // EndHash is called when there are no more fields in a hash.
// func (d *Decoder) EndHash(key []byte) {

// }

// // StartSet is called at the beginning of a set.
// // Sadd will be called exactly cardinality times before EndSet.
// func (d *Decoder) StartSet(key []byte, cardinality, expiry int64) {

// }

// // Sadd is called once for each member of a set.
// func (d *Decoder) Sadd(key, member []byte) {

// }

// // EndSet is called when there are no more fields in a set.
// // Same as EndHash
// func (d *Decoder) EndSet(key []byte) {
// }

// // StartList is called at the beginning of a list.
// // Rpush will be called exactly length times before EndList.
// // If length of the list is not known, then length is -1
// func (d *Decoder) StartList(key []byte, length, expiry int64) {
// }

// // Rpush is called once for each value in a list.
// func (d *Decoder) Rpush(key, value []byte) {
// }

// // EndList is called when there are no more values in a list.
// func (d *Decoder) EndList(key []byte) {
// }

// // StartZSet is called at the beginning of a sorted set.
// // Zadd will be called exactly cardinality times before EndZSet.
// func (d *Decoder) StartZSet(key []byte, cardinality, expiry int64) {
// }

// // Zadd is called once for each member of a sorted set.
// func (d *Decoder) Zadd(key []byte, score float64, member []byte) {
// }

// // EndZSet is called when there are no more members in a sorted set.
// func (d *Decoder) EndZSet(key []byte) {
// }

// // EndRDB is called when parsing of the RDB file is complete.
// func (d *Decoder) EndRDB() {
// }

func main() {
	f, err := os.Open(os.Args[1])
	fmt.Println(err)
	dr := &Decoder{
		m: MemProfiler{},
	}
	err = rdb.Decode(f, dr)
	fmt.Println(err)
}
