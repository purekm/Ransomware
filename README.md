
---

# AES-GCM Ransomware Decryptor (Python)

이 스크립트는 AES-GCM 방식으로 암호화된 랜섬웨어 파일을 복호화하기 위한 Python 예제입니다.
예시 파일: `C:/Users/Alpha/Desktop/FLAG.txt.ryk`

---

## 기능

* AES-GCM(cryptography 라이브러리 활용) 복호화
* 파일 형식: Nonce(12Bytes) + Ciphertext + Tag(16Bytes)
* AES 키: 32바이트(256비트) hex 문자열을 사용

---

## 실행 환경

* [cryptography](https://pypi.org/project/cryptography/) 라이브러리 설치 필요

```bash
pip install cryptography
```

---

## 사용 방법

1. 코드 안의 **파일 경로**와 \*\*AES 키(hexkey)\*\*를 실제 환경에 맞게 수정합니다.
2. 스크립트를 실행합니다.

```bash
python decrypt.py
```

3. 복호화에 성공하면 평문 데이터가 출력됩니다.

   * 성공: `b'원본 데이터 내용...'`
   * 실패: `Decryption failed: ...` 에러 메시지 출력

---

## 코드 설명

```python
nonce = enc_data[:12]    # AES-GCM Nonce (12바이트)
data  = enc_data[12:]    # Ciphertext + Tag(16바이트)
plaintext = gcm.decrypt(nonce, data, None)  # 복호화 및 태그 검증
```

* `AESGCM.decrypt()` 함수가 \*\*무결성 검증(Tag)\*\*까지 수행
* 키, nonce, 데이터가 일치하지 않으면 복호화 실패

---

