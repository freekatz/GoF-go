package iterator

import "fmt"

// ===
// 定义实现了 Iterator 接口的测试 ObjectIterator, 由 ObjectGroup 生成
// ===

type ObjectIterator struct {
	objectGroup *ObjectGroup // 想要实现 Iterator 要求 ObjectGroup 提供 Getter 方法
	objectIndex int          // 想要实现 Iterator 要求 ObjectGroup 可索引
}

func (iter *ObjectIterator) HasNext() bool {
	return iter.objectIndex < iter.objectGroup.Size()
}

func (iter *ObjectIterator) Next() (*Object, error) {
	if !iter.HasNext() {
		return nil, fmt.Errorf("There has no next object.\n")
	}
	o, _ := iter.objectGroup.GetObject(iter.objectIndex)
	iter.objectIndex++
	return o, nil
}

// 小写非导出, 确保只有实现了 Aggregate 接口的实例, 才可生成 Iterator
func newObjectIterator(og *ObjectGroup, i int) *ObjectIterator {
	return &ObjectIterator{
		objectGroup: og,
		objectIndex: i,
	}
}
