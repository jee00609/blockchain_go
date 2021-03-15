package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

//루프는 maxNonce만큼만 실행되며 이 값은 math.MaxInt64으로 정해준다.
var (
	maxNonce = math.MaxInt64
)

// 채굴되는 난이도 (목표 비트) 구하기
const targetBits = 24

// 블록 포인터와 타겟 포인터를 가진 ProofOfWork 구조체
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// NewProofOfWork는 ProofOfWork를 빌드하고 반환합니다.
func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1) //bit.Int을 1로 초기화

	//256은 SHA-256 해시의 비트 길이
	//SHA-256이 사용할 해시 알고리즘이다.
	target.Lsh(target, uint(256-targetBits)) //256 - targetBits 비트만큼 좌측 시프트 연산

	pow := &ProofOfWork{b, target}

	return pow
}

// 블록의 필드값들과 타겟 및 논스값을 병합
// 논스란 해시캐시에서의 카운터와 동일한 역할을 하는 암호학 용어다.
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

// Run performs a proof-of-work
// PoW 알고리즘의 핵심 코드를 구현
func (pow *ProofOfWork) Run() (int, []byte) {
	// 변수 초기화
	var hashInt big.Int //hash의 정수 표현값
	var hash [32]byte
	nonce := 0 //카운터

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	//루프 횟수를 제한하는 이유는 nonce의 오버플로우를 막기 위해서이다.
	//루프에서의 작업
	// 1. 데이터 준비 (생성)
	// 2. SHA-256 해싱
	// 3. 해시값의 큰 정수로의 변환
	// 4. 정수값과 타겟값 비교
	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Print("\n\n")

	return nonce, hash[:]
}

// Validate validates block's PoW
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
