import ECPairFactory from 'ecpair';
import * as ecc from 'tiny-secp256k1';
import BIP32Factory from 'bip32';
import * as bitcoin from 'bitcoinjs-lib';
import * as bip39 from 'bip39';
import secp256k1 from 'secp256k1'


bitcoin.initEccLib(ecc);

// 理论上所有的path 不是强制性的 path影响hd钱包私钥的生成，但与私钥生成公钥，公钥生成地址无关

// Pay-to-Witness-Public-Key-Hash  （bip84）
const P2WPKH_Path = `m/84'/0'/0'/0/`;
// Pay-to-Script-Hash (P2SH)
const P2SH_Path = `m/49'/0'/0'/0/`;
// ord
const P2TR_Path = `m/86'/0'/0'/0/`;
// Pay To Pubkey Hash
const P2PKH_Path = `m/44'/0'/0'/0/`;

// Pay-to-Witness-Script-Hash
const P2WSH_Path = `m/49'/0'/0'/0/`;

// 本文中没啥用 用hd钱包代替了 这个是生成单个的底层函数
const ECPair = ECPairFactory(ecc);
const bip32 = BIP32Factory(ecc);
// const mnemonic = bip39.generateMnemonic()
const mnemonic = "thrive member govern lake wealth alley theory divert roast screen poet maximum";

// bitcoin.payments.<func> 这种格式 是将信息写入 区块链对应的 scriptPubKey 由各自的节点认证
// 方法有七种  p2ms  p2pk  p2pkh p2sh p2tr p2wpkh p2wsh
// 转账相关的 P2PKH P2SH P2WPKH P2WSH p2tr
// p2ms  多签 本质上是将脚本输出到 scriptpubkey中 p2ms 是生成脚本的
// p2pk 是最原始的区块链签名 现在几乎已经废弃 占的byte太大

const addressTypes = async () => {
      const seed = await bip39.mnemonicToSeed(mnemonic);
      const rootKey = bip32.fromSeed(seed);
      // rootkey 可以转为base58  学名叫 Extended Private Key 缩写xprv
      // rootKey.toBase58()

      // 和unisat 的保持一致
      const p2wpkhChildNode = rootKey.derivePath(`${P2WPKH_Path}0`);
      const p2wpkhPayment = bitcoin.payments.p2wpkh({ pubkey: p2wpkhChildNode.publicKey });
      console.log("p2wpkh格式生成的公钥为：", p2wpkhPayment.pubkey?.toString("hex"))
      console.log("p2wpkh格式生成的地址为：", p2wpkhPayment.address)
      console.log("---")

      
      // 纯demo 和unisat无关
      const p2shChildNode0 = rootKey.derivePath(`${P2SH_Path}0`);
      const p2shChildNode1 = rootKey.derivePath(`${P2SH_Path}1`);
      const p2shChildNode2 = rootKey.derivePath(`${P2SH_Path}2`);
      const p2shPayment1 = bitcoin.payments.p2sh({
            redeem: bitcoin.payments.p2ms({
                  m: 2, pubkeys: [
                        p2shChildNode0.publicKey,
                        p2shChildNode1.publicKey,
                        p2shChildNode2.publicKey,
                  ]
            })
      });
      console.log("p2sh格式生成多签的公钥为：", p2shChildNode0.publicKey.toString('hex'), p2shChildNode1.publicKey.toString('hex'), p2shChildNode2.publicKey.toString('hex'))
      console.log("p2sh格式生成多签的最小签名人数：", p2shPayment1.redeem?.m, " p2sh格式生成多签的最大签名人数：", p2shPayment1.redeem?.n)
      console.log("p2sh格式生成多签的地址为：", p2shPayment1.address)
      console.log("---")

      // 和unisat 保持一致
      const p2shP2wpkhChildNode3 = rootKey.derivePath(`${P2SH_Path}0`);
      const p2shP2wpkhPayment2 = bitcoin.payments.p2sh({ redeem: bitcoin.payments.p2wpkh({ pubkey: p2shP2wpkhChildNode3.publicKey }) });
      console.log("p2sh-p2wpkh格式生成的公钥为：", p2shP2wpkhChildNode3.publicKey.toString('hex'))
      console.log("p2sh-p2wpkh格式生成的地址为：", p2shP2wpkhPayment2.address)
      console.log("---")

      // 和unisat 保持一致
      const p2trChildNode = rootKey.derivePath(`${P2TR_Path}0`);
      console.log(p2trChildNode.publicKey)
      let internalPubkey = p2trChildNode.publicKey.subarray(1, 33)
      const p2trPayment = bitcoin.payments.p2tr({ internalPubkey });
      console.log("p2tr格式生成的公钥为：", p2trChildNode.publicKey.toString('hex'))
      console.log("p2tr格式生成的地址为：", p2trPayment.address)
      console.log("---")

      // 和unisat 保持一致
      const p2pkhChildNode = rootKey.derivePath(`${P2PKH_Path}0`);
      const p2pkhPayment = bitcoin.payments.p2pkh({ pubkey: p2pkhChildNode.publicKey });
      console.log("p2pkh格式生成的公钥为：", p2pkhPayment.pubkey?.toString('hex'))
      console.log("p2pkh格式生成的地址为：", p2pkhPayment.address)
      console.log("---")

      // 纯demo 和unisat无关
      const p2wshChildNode0 = rootKey.derivePath(`${P2WSH_Path}0`);
      const p2wshChildNode1 = rootKey.derivePath(`${P2WSH_Path}1`);
      const p2wshChildNode2 = rootKey.derivePath(`${P2WSH_Path}2`);
      const p2wshChildNode3 = rootKey.derivePath(`${P2WSH_Path}3`);
      const p2wshPayment1 = bitcoin.payments.p2wsh({
            redeem: bitcoin.payments.p2ms({
                  m: 3, pubkeys: [
                        p2wshChildNode0.publicKey,
                        p2wshChildNode1.publicKey,
                        p2wshChildNode2.publicKey,
                        p2wshChildNode3.publicKey
                  ]
            })
      });

}


const test = async () => {
      const seed = await bip39.mnemonicToSeed(mnemonic);
      const rootKey = bip32.fromSeed(seed);
      const childNode = rootKey.derivePath(`${P2WPKH_Path}0`);
      console.log('私钥', childNode.privateKey?.toString('hex'))
      if (childNode.privateKey) {
            console.log("wif:", ECPair.fromPrivateKey(childNode.privateKey).toWIF())
            const _pubKey = secp256k1.publicKeyCreate(childNode.privateKey)
            console.log('Compressed Public Key:', Buffer.from(_pubKey).toString('hex'));
            const pubKey = secp256k1.publicKeyCreate(childNode.privateKey,false)
            
            console.log('un Compressed Public Key:', Buffer.from(pubKey).toString('hex'));
      }


}
addressTypes().then(test);

// 在交易中 是将公钥生成 script 进行广播
// 大体格式如下
// p2pkh OP_DUP OP_HASH160 OP_PUSHBYTES_20 <public key> OP_EQUALVERIFY OP_CHECKSIG
// p2sh OP_HASH160 OP_PUSHBYTES_20  <public key> OP_EQUAL 
// p2pk OP_PUSHBYTES_65  <public key> OP_CHECKSIG  或  OP_PUSHBYTES_33  <public key> OP_CHECKSIG
// 多签 OP_PUSHBYTES_<xx>  <public key> OP_PUSHBYTES_<xx>  <public key> ... OP_CHECKMULTISIG

// 公钥无法统一转成地址，因为公钥在不同的脚本中 生成的地址也不一样，例如，P2PKH地址使用了公钥的哈希，P2SH地址使用了脚本的哈希，而P2TR地址使用了Taproot脚本和哈希。这些算法的不同决定了地址的格式和前缀。比特币地址的结构取决于地址类型。P2PKH地址由公钥哈希构成，P2SH地址由脚本哈希构成，P2TR地址由脚本哈希构成，而P2WSH地址也由脚本哈希构成。这些不同的数据结构定义了地址的形式和编码方式。