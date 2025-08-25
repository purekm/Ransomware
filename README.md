
---

# AES-GCM Ransomware Decryptor (Python)

이 스크립트는 AES-GCM 방식으로 암호화된 랜섬웨어 파일을 복호화하기 위한 Python script 입니다.
테스트 파일: `C:/Users/Alpha/Desktop/FLAG.txt.ryk`

---

## 주요 기능

* AES-GCM(cryptography 라이브러리 활용) 복호화
* 파일 형식: Nonce(12Bytes) + Ciphertext + Tag(16Bytes)
* AES 키: 32바이트(256비트) hex 문자열을 사용

---

## 실행 환경

* cryptography 라이브러리 설치 필요

```bash
pip install cryptography
```

---

## 사용 방법

1. 코드 안의 cipherfile에 **파일 경로**와 hexkey에 **AES 키**를 실제 환경에 맞게 수정합니다.
2. 스크립트를 실행합니다.

```bash
python decrypt.py
```

3. 복호화에 성공하면 평문 데이터가 출력됩니다.

   * 성공: `b'원본 데이터 내용...'`
   * 실패: `Decryption failed: ...` 에러 메시지 출력

---

## 작동방식

1. 암호화된 파일 읽기 → .ryk 파일을 바이트 단위로 불러옴
2. AES 키 준비 → hex 문자열을 32바이트 키(AES-256)로 변환
3. Nonce / Ciphertext 분리 → 파일 앞 12바이트는 Nonce, 이후는 암호문+태그
4. AES-GCM 복호화 & 무결성 검증 → AESGCM.decrypt()로 복호화 수행, 태그 불일치 시 실패
5. 결과 출력 → 성공 시 평문 데이터 출력, 실패 시 에러 메시지 출력
---

