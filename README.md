# Kuri API Encryption

## Spesifikasi Tugas
1. API can do encryption and decryption text
2. There are _caesarcipher_ and _vigenerecipher_ for encrypt and decrypt text
3. This API using golang and gin as a router

## API Spesification

### Caesar Cipher
1. An endpoint contains _string_ plaintext and _int_ key, and change it into ciphertext. Endpoint will receive **POST** request with payload:
```JSON
{
  "plaintext": InsertPlaintextHere,
  "key": 12
}
```

2. An endpoint contains _string_ ciphertext and _int_ key, and change it into plaintext. Endpoint will receive **POST** request with payload:
```JSON
{
  "ciphertext": InsertCiphertextHere,
  "key": 12
}
```

**Request:**
```
http://0.0.0.0:8080/caesarcipher
```

### Vigenere Cipher
1. An endpoint contains _string_ plaintext and _string_ key, and change it into ciphertext. Endpoint will receive **POST** request with payload:
```JSON
{
  "plaintext": InsertPlaintextHere,
  "key": KURI
}
```

2. An endpoint contains _string_ ciphertext and _string_ key, and change it into plaintext. Endpoint will receive **POST** request with payload:
```JSON
{
  "ciphertext": InsertCiphertextHere,
  "key": KURI
}
```

**Request:**
```
http://0.0.0.0:8080/vigenerecipher
```

**Make sure plaintext and ciphertext only contains char 'A'-'Z' and 'a'-'z'**