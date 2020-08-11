# stellar-adapter

stellar-adapter适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建XLM.ini文件，编辑如下内容：

```ini

# stellar service
isScan = true
ServerAPI = https://horizon-testnet.stellar.org/
Network = "Test SDF Network ; September 2015"
#提币是否创建不存在的账户
IsCreateNotExistsAccount = true
#账户保留余额，链限制
AddressRetainAmount = 1

```


### 说明

钱包创建的地址
1 可以通过币安交易所提币创建 
2 通过 tx_decode_test.go 里面的TestBuildCreateAccountTransaction 方法创建
3 有密钥对可以通过 https://laboratory.stellar.org/  创建 


### 官网

https://www.stellar.org/

### stellar 功能实验室 

https://laboratory.stellar.org/

### 区块浏览器

https://stellar.expert/explorer/public/

### github

https://github.com/stellar/stellar-core

