package encryption

import (
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/gogf/gf/v2/text/gstr"
	"testing"
)

func Test_Md5Encode(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		t.Assert(Md5Encode("123456"), "e10adc3949ba59abbe56e057f20f883e")
		t.Assert(gstr.Equal(gstr.ToLower("E10ADC3949BA59ABBE56E057F20F883E"), Md5Encode(gstr.Trim("123456"))), true)

		t.Assert(ComputeHmacSha256("123456", "123456"), "b8ad08a3a547e35829b821b75370301dd8c4b06bdd7771f9b541a75914068718")

	})
}
