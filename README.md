# triam-adapter

triam-adapter适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。
本适配器可以改造为stellar(xlm)恒星币

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建RIA.ini文件，编辑如下内容：

```ini

# triam service
ServerAPI = https://testnet-horizon.triamnetwork.com
Network = "SAAK5654--ARM-NETWORK--BHC3SQOHPO2GGI--BY-B.A.P--CNEMJQCWPTA--RUBY-AND-BLOCKCHAIN--3KECMPY5L7W--THANKYOU-CS--S542ZHDVHLFV
AddressRetainAmount = 20

```

## 资料介绍

### 官网
https://triamnetwork.com/

### 区块浏览器

https://dashboard.triamnetwork.com/

### github

https://github.com/triamnetwork/triam-core