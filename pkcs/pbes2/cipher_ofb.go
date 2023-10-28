package pbes2

import (
    "errors"
    "crypto/rand"
    "crypto/cipher"
    "encoding/asn1"
)

// OFB 模式加密参数
type ofbParams []byte

// OFB 模式加密
type CipherOFB struct {
    cipherFunc func(key []byte) (cipher.Block, error)
    keySize    int
    blockSize  int
    identifier asn1.ObjectIdentifier
}

// 值大小
func (this CipherOFB) KeySize() int {
    return this.keySize
}

// oid
func (this CipherOFB) OID() asn1.ObjectIdentifier {
    return this.identifier
}

// 加密
func (this CipherOFB) Encrypt(key, plaintext []byte) ([]byte, []byte, error) {
    block, err := this.cipherFunc(key)
    if err != nil {
        return nil, nil, errors.New("pkcs/cipher: failed to create cipher: " + err.Error())
    }

    // 随机生成 iv
    iv := make(ofbParams, this.blockSize)
    if _, err := rand.Read(iv); err != nil {
        return nil, nil, errors.New("pkcs/cipher: failed to generate IV: " + err.Error())
    }

    // 需要保存的加密数据
    encrypted := make([]byte, len(plaintext))

    enc := cipher.NewOFB(block, iv)
    enc.XORKeyStream(encrypted, plaintext)

    // 编码 iv
    paramBytes, err := asn1.Marshal(iv)
    if err != nil {
        return nil, nil, err
    }

    return encrypted, paramBytes, nil
}

// 解密
func (this CipherOFB) Decrypt(key, params, ciphertext []byte) ([]byte, error) {
    // 解析出 iv
    var iv ofbParams
    if _, err := asn1.Unmarshal(params, &iv); err != nil {
        return nil, errors.New("pkcs/cipher: invalid iv parameters")
    }

    block, err := this.cipherFunc(key)
    if err != nil {
        return nil, err
    }

    if len(iv) != block.BlockSize() {
        return nil, errors.New("pkcs/cipher: incorrect IV size")
    }

    plaintext := make([]byte, len(ciphertext))

    mode := cipher.NewOFB(block, iv)
    mode.XORKeyStream(plaintext, ciphertext)

    return plaintext, nil
}
