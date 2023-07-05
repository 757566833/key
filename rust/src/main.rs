use std::str::FromStr;

use bip39::Mnemonic;
use bitcoin::{
    address::Address,
    bip32,
    // Address,
    Network,
    PublicKey,
};
use secp256k1::{Secp256k1, XOnlyPublicKey};
// use rand::prelude::*;
// fn generate_entropy(bits: usize) -> Vec<u8> {
//     let bytes = bits / 8;
//     let mut entropy = vec![0u8; bytes];
//     let mut rng = rand::thread_rng();
//     rng.fill_bytes(&mut entropy);
//     entropy
// }
static P2_WPKH_PATH: &str = "m/84'/0'/0'/0/";
static P2_SH_PATH: &str = "m/49'/0'/0'/0/";
static P2_TR_PATH: &str = "m/86'/0'/0'/0/";
static P2_PKH_PATH: &str = "m/44'/0'/0'/0/";

fn main() {
    // let entropy = generate_entropy(256);
    // let mnemonic =Mnemonic::from_entropy(&entropy).unwrap();
    // let mnemonic_string = mnemonic.to_string();
    // println!("{}",mnemonic_string);

    let mnemonic_string =
        "thrive member govern lake wealth alley theory divert roast screen poet maximum";
    let mnemonic = Mnemonic::from_str(mnemonic_string).unwrap();
    let seed = mnemonic.to_seed("");
    let root_key = bip32::ExtendedPrivKey::new_master(Network::Bitcoin, &seed).unwrap();


    // p2wpkh
    let p2wpkh_path = bip32::DerivationPath::from_str(&(P2_WPKH_PATH.to_owned() + "0")).unwrap();
    let p2wpkh_secp = Secp256k1::new();
    let p2wpkh_child_key = root_key
        .derive_priv(&p2wpkh_secp, &p2wpkh_path)
        .unwrap();
    let p2wpkh_public_key = bip32::ExtendedPubKey::from_priv(&p2wpkh_secp, &p2wpkh_child_key);
    println!("p2wpkh格式生成的公钥为： {}", p2wpkh_public_key.public_key);
    let p2wpkh_bitcoin_public_key: PublicKey = PublicKey {
        compressed: true,
        inner: p2wpkh_public_key.public_key,
    };
    let p2wpkh_address = Address::p2wpkh(&p2wpkh_bitcoin_public_key, Network::Bitcoin).unwrap();
    let p2wpkh_address_str = p2wpkh_address.to_string();
    println!("p2wpkh格式生成的地址为：{}", p2wpkh_address_str);
    println!("---");


    // p2sh-p2wpkh
    let p2sh_p2wpkh_path = bip32::DerivationPath::from_str(&(P2_SH_PATH.to_owned() + "0")).unwrap();
    let p2sh_p2wpkh_secp = Secp256k1::new();
    let p2sh_p2wpkh_child_key = root_key
        .derive_priv(&p2sh_p2wpkh_secp, &p2sh_p2wpkh_path)
        .unwrap();
    let p2sh_p2wpkh_public_key =
        bip32::ExtendedPubKey::from_priv(&p2sh_p2wpkh_secp, &p2sh_p2wpkh_child_key);
    println!(
        "p2sh_p2wpkh格式生成的公钥为： {}",
        p2sh_p2wpkh_public_key.public_key
    );

    let p2sh_p2wpkh_bitcoin_public_key: PublicKey = PublicKey {
        compressed: true,
        inner: p2sh_p2wpkh_public_key.public_key,
    };

    let inner_p2sh_p2wpkh_address =
        Address::p2wpkh(&p2sh_p2wpkh_bitcoin_public_key, Network::Bitcoin).unwrap();
    let binding = inner_p2sh_p2wpkh_address.script_pubkey();
    let script = binding.as_script();
    let a = Address::p2sh(script, Network::Bitcoin).unwrap();
    let p2sh_p2wpkh_address_str = a.to_string();
    println!("p2sh_p2wpkh格式生成的地址为：{}", p2sh_p2wpkh_address_str);
    println!("---");


    // p2tr
    let p2tr_path = bip32::DerivationPath::from_str(&(P2_TR_PATH.to_owned() + "0")).unwrap();
    let p2tr_secp = Secp256k1::new();
    let p2tr_child_key = root_key.derive_priv(&p2tr_secp, &p2tr_path).unwrap();
    let p2tr_public_key = bip32::ExtendedPubKey::from_priv(&p2tr_secp, &p2tr_child_key);
    println!("p2tr格式生成的公钥为： {}", p2tr_public_key.public_key);
    let p2tr_address = Address::p2tr(
        &p2tr_secp,
        XOnlyPublicKey::from(p2tr_public_key.public_key),
        None,
        Network::Bitcoin,
    );
    let p2tr_address_str = p2tr_address.to_string();
    println!("p2tr格式生成的地址为：{}", p2tr_address_str);
    println!("---");


    // p2pkh
    let p2pkh_path = bip32::DerivationPath::from_str(&(P2_PKH_PATH.to_owned() + "0")).unwrap();
    let p2pkh_secp = Secp256k1::new();
    let p2pkh_child_key = root_key
        .derive_priv(&p2pkh_secp, &p2pkh_path)
        .unwrap();
    let p2pkh_public_key = bip32::ExtendedPubKey::from_priv(&p2pkh_secp, &p2pkh_child_key);
    println!("p2pkh格式生成的公钥为： {}", p2pkh_public_key.public_key);
    let p2pkh_bitcoin_public_key: PublicKey = PublicKey {
        compressed: true,
        inner: p2pkh_public_key.public_key,
    };
    let p2pkh_address = Address::p2pkh(&p2pkh_bitcoin_public_key, Network::Bitcoin);
    let p2pkh_address_str = p2pkh_address.to_string();
    println!("p2pkh格式生成的地址为：{}", p2pkh_address_str);
}
