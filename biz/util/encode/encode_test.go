package encode

import "testing"

func TestEncodePassword(t *testing.T) {
	// 36d4b1a948112897fae64512a7554ee1a3d15248be69aab7a5ff4fb1e3ec2fa8
	t.Log(EncodePassword("NzHsbHoe9qdMM5DA", "1234546"))
	t.Log(EncodePassword("1234546", "NzHsbHoe9qdMM5DA"))
}
