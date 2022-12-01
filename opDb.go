package main

// //leveldb的key和value都是byte数组类型

// //根据key获取value
// func GetValueByKey(key string) string {
// 	data, err := DB.Get([]byte(key), nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return string(data)
// }

// //设置key和value
// func SetKeyAndValue(key string, value []byte) bool {
// 	err := DB.Put([]byte(key), value, nil)
// 	return err == nil
// }

// //根据key删除key value
// func DelKeyAndValue(key string) bool {
// 	_, err := DB.Get([]byte(key), nil)
// 	if err == leveldb.ErrNotFound {
// 		return false
// 	}
// 	err2 := DB.Delete([]byte(key), nil)
// 	return err2 == nil
// }
