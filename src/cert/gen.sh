# https://dev.to/techschoolguru/how-to-create-sign-ssl-tls-certificates-2aai

# 字段顺序无关性
# 各字段顺序可自由调整（如 /C=FR/ST=... 或 /ST=.../C=FR），但需确保键值对完整1。

# 特殊字符处理
# 若字段值包含空格或特殊符号（如 O=Tech School），需用引号包裹整个 -subj 参数4。

# 必填与选填字段

# 必填字段：/C、/ST、/L、/O、/CN（根据证书用途，如 HTTPS 证书必须设置域名到 CN 或 SAN）14。
# 选填字段：/OU、/emailAddress 等

# 字段标识	含义	示例值	用途说明
# /C	Country（国家）	FR	符合 ISO 3166 国家代码，标识证书持有者的国家
# /ST	State/Province（州/省）	Occitanie	证书持有者所在行政区域（如法国的大区
# /L	Locality（城市）	Toulouse	持有者所在城市或地区
# /O	Organization（组织）	Tech School	标识证书所属的机构或公司
# /OU	Organizational Unit（部门）	Education	组织内的具体部门或分支
# /CN	Common Name（通用名称）	*.techschool.guru	关键字段，用于标识域名（如 HTTPS 证书）或服务名称
# /emailAddress	邮箱地址	techschool.guru@gmail.com	可选字段，标识联系邮箱

rm *.pem

# -nodes 并非英文单词 "nodes"，而是缩写 "no DES"，表示生成的私钥文件不通过 DES（或其他加密算法）进行加密

# 功能说明

# 使用此参数时，OpenSSL 生成的私钥文件（如 .key）无需设置密码，可直接被程序读取
# 若省略 -nodes，私钥默认会以 3DES-CBC 算法加密，使用时需手动输入密码
# 在 OpenSSL 中，私钥加密算法可通过命令行参数灵活指定，支持的算法包括但不限于：

# AES 系列：aes-128-cbc、aes-256-cfb、aes-256-ofb 等（支持 128/192/256 位密钥）1
# Camellia：如 camellia-256-cbc
# SEED、DESX、BF（Blowfish）等1
# 二、指定加密算法的命令行方法
# 通过 OpenSSL 生成或转换私钥时，可通过 -<algorithm> 参数指定加密方式：

# 生成新私钥时指定算法
# openssl genrsa -aes256 -out private.key  2048  # 使用 AES-256-CBC 加密私钥 
# openssl genpkey -algorithm RSA -aes-128-cbc -out private.pem   # 使用 AES-128-CBC 
# 转换现有私钥的加密算法
# openssl rsa -in private.key  -aes-128-cfb -out new_private.key   # 将原私钥转换为 AES-128-CFB 加密 
# 直接使用 enc 命令加密/解密
# openssl enc -aes-256-cbc -in private.key  -out encrypted.key   # 用 AES-256-CBC 加密文件 
# openssl enc -d -aes-256-cbc -in encrypted.key  -out decrypted.key   # 解密 
# 三、注意事项
# 算法兼容性
# 不同版本的 OpenSSL 支持的算法可能不同（如旧版本可能不支持 AES-256-GCM）
# 使用前可通过 openssl enc -list 查看当前支持的算法列表。
# 安全性建议
# 优先选择 AES-256 或 Camellia-256 等强加密算法，避免使用 DES/3DES 等弱算法。
# 避免使用 ECB 模式（如 aes-256-ecb），因其安全性较低1。
# 密码保护
# 无论选择何种算法，均需通过 -pass 参数或交互式输入设置强密码，例如：

# openssl genrsa -aes256 -passout pass:your_password -out private.key  


# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=FR/ST=Occitanie/L=Toulouse/O=Tech School/OU=Education/CN=*commonnameCA"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=FR/ST=Ile de France/L=Paris/O=PC Book/OU=Computer/CN=*.commonnameServer"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem -subj "/C=FR/ST=Alsace/L=Strasbourg/O=PC Client/OU=Computer/CN=*.commonnameClient"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -in client-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf

echo "Client's signed certificate"
openssl x509 -in client-cert.pem -noout -text
