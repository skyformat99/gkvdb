package gkvdb

// 设置是否开启缓存
func (db *DB) SetCache(enabled bool) {
    if enabled {
        db.setCache(1)
    } else {
        db.setCache(0)
    }
}

// 关闭数据库链接
func (db *DB) Close() {
    db.close()
}

// 查询KV数据
func (db *DB) Get(key []byte) []byte {
    if v, ok := db.memt.get(key); ok {
        return v
    }
    return db.get(key)
}

// 设置KV数据
func (db *DB) Set(key []byte, value []byte) error {
    if db.isCacheEnabled() {
        if err := db.memt.set(key, value); err != nil {
            return err
        }
        return nil
    }
    return db.set(key, value)
}

// 删除KV数据
func (db *DB) Remove(key []byte) error {
    if db.isCacheEnabled() {
        if err := db.memt.remove(key); err != nil {
            return err
        }
        return nil
    }
    return db.remove(key)
}

// 设置KV数据(强制不使用缓存)
func (db *DB) SetWithoutCache(key []byte, value []byte) error {
    return db.set(key, value)
}

// 获取max条随机键值对，max=-1时获取所有数据返回
// 该方法会强制性遍历整个数据库
func (db *DB) Items(max int) map[string][]byte {
    return db.items(max)
}

// 获取最多max个随机键名，构成列表返回
func (db *DB) Keys(max int) []string {
    m    := db.Items(max)
    keys := make([]string, 0)
    for k, _ := range m {
        keys = append(keys, k)
    }
    return keys
}

// 获取最多max个随机键值，构成列表返回
func (db *DB) Values(max int) [][]byte {
    m      := db.Items(max)
    values := make([][]byte, 0)
    for _, v := range m {
        values = append(values, v)
    }
    return values
}

// 打印数据库状态(调试使用)
//func (db *DB) PrintState() {
//    mtblocks := db.mtsp.GetAllBlocks()
//    dbblocks := db.dbsp.GetAllBlocks()
//    fmt.Println("meta pieces:")
//    fmt.Println("       size:", len(mtblocks))
//    fmt.Println("       list:", mtblocks)
//
//    fmt.Println("data pieces:")
//    fmt.Println("       size:", len(dbblocks))
//    fmt.Println("       list:", dbblocks)
//
//    fmt.Println("=======================================")
//}


