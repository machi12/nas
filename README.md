## 修改

- nasMessage/NAS_AuthenticationRequest.go：在认证请求消息中增加了SNMAC以及相关的编解码操作。
- nasType/NAS_AuthenticationParameterSNMAC.go：增加了SNMAC结构体及其相关操作。
- nasType/NAS_N.go：增加了N结构体及其相关操作。
- nasMessage/NAS_RegistrationRequest.go：在注册请求消息中增加了N以及相关的编解码操作。
- nasType/NAS_AuthenticationParameterAUTN.go：对AUTN的长度进行修改，5G-ESAKA中AUTN的长度为64bits
