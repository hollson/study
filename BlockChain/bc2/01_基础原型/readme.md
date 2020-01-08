# 定义

// 区块
Block{
PrevBlockHash []byte
Hash []byte
Timestamp int64
Height int64 //高度
Merkle []byte
Nonce 

}

//唯一方法，关键方法 设置哈希
func SetHash(prev []byte) []byte{
new一个对象
  bc：=&Block「{

  }
  
  // 将成员值的字节数据拼接
  arr=bytes.jion(,)


return arrs
}

  //转为二进制字节数组
func ToHex(n int) []byte{
var buf []byte
 binary.write(buf，高位对齐，num)
 ret buf
}

//创世区块
func GenerantBlock(){
  
}


// 区块链，就是一个大数组
type BlockChain struct{
 []Block
}




