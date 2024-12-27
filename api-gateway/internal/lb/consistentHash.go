package lb

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	// hash cricle
	keys    []int
	hashMap map[int]string
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// 增加节点到hash circle中
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

func (m *Map) AddOne(key string) {
	for i := 0; i < m.replicas; i++ {
		hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
		m.keys = append(m.keys, hash)
		m.hashMap[hash] = key
	}
	sort.Ints(m.keys)
}

func (m *Map) Delete(key string) {
	for i := 0; i < m.replicas; i++ {
		hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
		idx := sort.Search(len(m.keys), func(i int) bool {
			return m.keys[i] >= hash
		})
		if m.keys[idx] == hash {
			m.keys = append(m.keys[:idx], m.keys[idx+1:]...)
			delete(m.hashMap, hash)
		}
	}
}

func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	// 从hash circle中正向寻找到第一个大于hash的节点
	idx := sort.Search(len(m.keys), func(i int) bool {
		return m.keys[i] >= hash
	})
	// if hash is larger than the last node, return the first node idx%len(m.keys)
	return m.hashMap[m.keys[idx%len(m.keys)]]
}

func ConsistentHash(serviceNames []string, key string) string {
	m := New(5, nil)
	m.Add(serviceNames...)
	return m.Get(key)
}
