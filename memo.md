* func method(long maxSplitSize)
* 加载文件
* 获取文件大小
* 计算分片数量
    * 对文件进行逻辑划分
* go func 读取各个分片
    
    * lineRecordReader --读取start～end的内容
        * LineReader --从缓存中读取某一行
        * 读取的每一行内容放入chan
    * Mapper
        * 转化为csv格式
        * 构造key value

********************************************
