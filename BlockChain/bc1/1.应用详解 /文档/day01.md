[TOC]

1.  区块链产生
   1. 密码朋克：通过匿名性来保护隐私安全
   2. 不可篡改文件记录：加盖时间戳到文件内容中，证实数据真实存在，保证安全
   3. 数字现金
2.  发展历史
   1. 09，比特币诞生 1.0
   2. 14年，以太坊诞生 2.0 
   3. 18中期，EOS主网上线 3.0
   4. Fabric
3.  区块链场景概念
   1. 数字货币：可以与现金流进行等价交换，拥有现有现金流的属性和功能
   2. ICO：众筹
   3. 电子钱包：在商的购物中常用的支付工具
4.  什么是区块链(定义)
   1. 区块是一种集成了点对点传输协议，现代密码学，共识算法，分布数据存储等成型技术新型的应用模型
5.  区块链特点
   1. 可追溯
   2. 不可篡改
   3. 去中心化
   4. 完整备份：区块链具有完整的分布式存储特性
   5. 历史记录：被存储的数据拥有完整的历史记录，可以快速查看，复原
   6. 交易广播：一次交易分发给网络中的其它节点，同步进行接收
6.  区块链加密货币特点
   1. 独立性：所有货币都是独立存在的
   2. 唯一性：地址、交易都具有不重复的唯一性
   3. 匿名性：账户信息和个人信息没有关联，整个交易过程全程加密
   4. 不可伪造
7.  区块链核心技术
   1. 点对点传输协议：在网络中的数据流通方式
   2. 现代密码学：在区块链中的应用：公私钥签名、哈希算法
   3. 共识算法：数据一致性
   4. 分布式数据存储：实现去中心化的重要技术依据
8.  区块链核心概念
   1. 区块链
      1. 本质：一个分布式**账本**，通过共识算法来决定谁能抢到当前的**记账权**。区块链以**区块**为单位，以区块产生的时间顺序去进行**连接**。
   2. 区块
      1. 概念：区块链的基本组成单位
      2. 区块头
         1. 时间戳
         2. 当前区块哈希
         3. 父区块哈希
         4. 随机数
         5. Merkle树
         6. 区块号码
      3. 区块体：交易数据
   3. 分布式数据库：区块链中的区块数据都存储在每一个节点中，所有节点组成一个分布式数据库。任何一个或者多个节点退出，都不会影响其它节点，因为每个节点都保存了完整的数据
   4. 节点
      1. 可以理解成为一个运行区块软件的计算机
      2. 分类
         1. 全节点：保存了完整的区块链副本，安全性极高，效率不高
         2. 轻节点：不保存所有的区块，需要依赖全节点进行验证，效率更高，安全不如全节点高
         3. 挖矿节点：带有挖矿功能的全节点，专门处理交易验证
   5. 挖矿
      1. 对交易进行验证处理(记账)，区块就是通过挖矿产生的
      2. 穷举随机数算法，生成哈希，与目标哈希进行比较，成功则说明挖矿成功
   6. 分叉
      1. 升级分叉
         1. 矿工遵从不同的机制(规则)导致分叉
         2. 硬分叉：如果区块链共识规则改变之后，不允许前向兼容，旧节点没有办法认可新节点产生的区块
         3. 软分叉：如果区块链共识规则改变之后，允许前向兼容，旧节点可以兼容新节点产生的区块
      2. 挖矿分叉
         1. 现象：两个或者多个矿工，同时完成了工作量证明，就会产生两个新的区块，形成分叉
         2. 解决方案：不同矿工跟随了不同的区块，但是不同链算力会有区别，矿工的数量一样，链的增长速度就不会相同，最终会出现一个条链更长，这条就会变成主链。
   7. 51%攻击
   8. 交易
      1. 概念：一笔资产在参与都之间的转移
      2. 内容
         1. 金额
         2. 发送者
         3. 接收者
         4. 交易ID(HASH)
   9. 双花
      1. 概念：复用数字货币的数字特性，可以完成两次或者多次支付
      2. 传统虚拟货币之所以可以避免双花是因为有可依赖的第三方机构提供保证
      3. 区块中需要达成只通过分布式节点之间的相互校验与共识机制来避免双花，同时完成价值转移 
   10. UTXO(unspent transaction output)交易模式
       1. 是比特币独有的交易模式，比特币交易过程中的基本单位，主要就是为了避免双花
   11. 哈希
       1. 将任意的原始数据(交易记录)通过指定哈希函数，编码为特定长度的字符串
       2. 在区块链中的使用：生成地址，交易验证
       3. 特点
          1. 不可逆
          2. 随机性
          3. 时间正相关：输入的源数据越长，哈希的处理时间就更长
   12. 加密算法
       1. 对称加密：加密与解密都使用相同的密钥
       2. 非对称加密
          1. 采用公钥和私钥进行加密
          2. 无法用公钥反推私钥
   13. 数字签名
   14. Merkle树
       1. Merlkle树可以二叉树，也可以是多叉树，它具有树的所有特点
       2. 在区块链中的作用：快速校验、归纳交易数据的完整性
       3. 在区块链中，Merkle树可以极大的提高查询效率，区块头只需要保存一个Merkle根的hash
       4. Merkle支持SPV
   15. P2P
       1. 通过对等网络来分配工作任务的分布式应用架构
       2. 迅雷采用的就是P2P
       3. 由于在P2P中，所有网络节点的地位是对等的，不存在任何一个中心化节点，也不存在所谓的层级结构，所以每个节点都需要承担验证区块数据等功能
9.  区块链分类
   1. 公有链：真正意义上的去中心化分布式区块链，任何一个节点都可以随时加入/退出网络中。
   2. 私有链：部分中心化的区块链，具有一定的分布式特点，但是有一个中心节点，可以指定参与者
   3. 联盟链：部分去中心化的区块链，拥有权限控制的功能
      1. 代表：Fabric
10.  区块链架构特点
   1.  去中心化：基于分布式系统，整个网络中没有中心机构存在
   2.  可靠数据库：分布式存储，参与系统的节点越多，数据库的安全性就越高
   3.  开源可编程：区块链提供了灵活的脚本系统甚至于完善的开发平台，支持用户创建更加高级的应用
   4.  集体维护：区块链中的数据，由整个系统中所有具有记账功能的节点进行维护
   5.  安全可信：通过现代密码学实现
   6.  准匿名性：采用与身份信息无关的哈希作为哈希地址与交易ID
11.  典型应用分析 
    1.  比特币
       1.  特点
          1.  问题2100W，永不增发
          2.  挖矿奖励，逐年减半
       2.  架构
          1.  前端
             1.  钱包：保存用户私钥，管理用户余额，提供比特币交易(支持、转账)
             2.  钱包分类
                1.  决定性钱包：所有的私钥都由一个私钥种子通过单向哈希算法生成
                   1.  普通决定性钱包：由私钥种子一次性生成所有私钥
                   2.  层级决定性钱包：由私钥种子生成父私钥，父私钥生成子私钥
                2.  非决定性钱包：直接保存私钥，私钥直接放在DB上面
             3.  展示方式的分类
                1.  桌面钱包
                   1.  厚钱包：下载整条区块链，可以完整交易，安全性高，验证成本高
                   2.  薄钱包：不会下载整条区块链，采用部分存储+节点请求验证的方式
                   3.  离线钱包：USB设备，纸钱包，可以有效防范网络攻击
                2.  HTTP/JSON RPC API
                   1.  比特币提供的接口，可以使外部通过该接口访问或者控制比特币节点
                3.  命令行接口
                   1.  通过命令行的方式实现类似钱包的功能
                4.  浏览器
                   1.  访问区块链的区块数据等信息
          2.  节点后台：负责参数比特币网络的通信，区块链维护，验证，交易...
          3.  比特币地址
             1.  基本概念：由哈希生成
             2.  生成过程
                1.  随机数生成私钥
                2.  采用Secp256k1椭圆加密算法生成公钥
             3.  生成地址
                1.  以公钥作为输入，进行SHA256，再进行RIPEMD160，最后通过base58编码生成比特币地址
          4.  比特币区块校验：确保确实完成了工作量证明
             1.  校验内容：格式 、难度、时间戳、大小、交易
          5.  比特币交易(UTXO)
             1.  交易结构
                1.  输入(input)
                2.  输出(output)
                3.  交易ID(HASH)
             2.  说明
                1.  每一笔交易的输入来自于前面交易的输出
             3.  UTXO数据库：专门用于存储当前比特中未被花费的输出
    2.  以太坊
       1.  以太坊基本概念：一个用于开发去中心化DAPP的分布式平台
       2.  智能合约：一个拥有自我校验与自我执行的协议
       3.  以太坊为什么可以被称为区块链2.0：以太坊提供了非常方便的应用开发平台，对底层做了完善的封装，让开发者只需要关注于上层应用
       4.  共识 ： POW(ETHASH)
    3.  EOS
       1.  EOS基本概念：一个用于开发去中心化DAPP的分布式平台 
       2.  与以太坊的比较：
          1.  TPS有了极大的提高，能够达到百万级的TPS处理量
          2.  以太坊本身是一条公链，其中每个DAPP会消耗整条链上的资源
          3.  EOS本身不再是一条单纯的公链，它是一个区块链的基础架构，开发者可以自由在EOS上面创建公链，链与链之间不会影响彼此的资源使用，因此，不会出现单个应用占用的资源太多使得整个网络拥堵
       3.  共识算法：DPOS
       4.  EOS上的智能合约调用不需要手续费
    4.  Fabric(联盟链)
       1.  联盟链代表，与公链最大的区别在于不发行虚拟货币
       2.  目标：实现一个通过权限管理区块链的底层基础框架
       3.  架构分类
          1.  身份服务：Fabric具有身份识别能力，加入/退出Fabric网络需要权限认证
          2.  策略服务：提供访问控制，授权等一系列功能
          3.  区块链服务：提供构建分布式账本的基础能力以及数据传输，共识达成的底层功能
          4.  智能合约：验证节点上运行的分布式程序，用于自动执行特定的业务规则