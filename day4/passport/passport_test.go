package passport

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\n" +
	"byr:1937 iyr:2017 cid:147 hgt:183cm\n" +
	"\n" +
	"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\n" +
	"hcl:#cfa07d byr:1929\n" +
	"\n" +
	"hcl:#ae17e1 iyr:2013\n" +
	"eyr:2024\n" +
	"ecl:brn pid:760753108 byr:1931\n" +
	"hgt:179cm\n" +
	"\n" +
	"hcl:#cfa07d eyr:2025 pid:166559648\n" +
	"iyr:2011 ecl:brn hgt:59in\n"

var (
	i1 = "ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\n" +
		"byr:1937 iyr:2017 cid:147 hgt:183cm\n"

	i2 = "iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884\n" +
		"hcl:#cfa07d byr:1929\n"

	i3 = "hcl:#ae17e1 iyr:2013\n" +
		"eyr:2024\n" +
		"ecl:brn pid:760753108 byr:1931\n" +
		"hgt:179cm\n"

	i4 = "hcl:#cfa07d eyr:2025 pid:166559648\n" +
		"iyr:2011 ecl:brn hgt:59in\n"
)

func TestParsePassport(t *testing.T) {
	_, err := ParsePassport(i1)
	require.NoError(t, err)

	_, err = ParsePassport(i2)
	require.Error(t, err)

	_, err = ParsePassport(i3)
	require.NoError(t, err)

	_, err = ParsePassport(i4)
	require.Error(t, err)

}

func TestSplitInput(t *testing.T) {
	in := strings.Split(input, "\n")
	p := SplitInput(in)
	fmt.Printf("length: %v\n", len(p))
	require.Len(t, p, 4)
}

func TestValidateNumber(t *testing.T) {
	require.NoError(t, validateNumber(4, 2010, 2020, "2010"))
	require.NoError(t, validateNumber(4, 2010, 2020, "2020"))
	require.NoError(t, validateNumber(4, 1009, 2020, "1009"))
	require.NoError(t, validateNumber(4, 2010, 2020, "2015"))

	require.Error(t, validateNumber(4, 2, 2020, "201"))
	require.Error(t, validateNumber(4, 2010, 20200, "20001"))
	require.Error(t, validateNumber(4, 2010, 2020, "2001"))
	require.Error(t, validateNumber(4, 2010, 2020, "2009"))
	require.Error(t, validateNumber(4, 2010, 2020, "2021"))
	require.Error(t, validateNumber(4, 2010, 2020, "PIES"))
}

func TestValidateHeight(t *testing.T) {
	require.NoError(t, validateHeight("60in"))
	require.NoError(t, validateHeight("190cm"))

	require.Error(t, validateHeight("190in"))
	require.Error(t, validateHeight("190"))
}

func TestValidateHairColour(t *testing.T) {
	require.NoError(t, validateHairColour("#123abc"))
	require.NoError(t, validateHairColour("#123456"))
	require.NoError(t, validateHairColour("#abcdEF"))

	require.Error(t, validateHairColour("abcdEF"))
	require.Error(t, validateHairColour("#adEF"))
	require.Error(t, validateHairColour("#adEFDSFAS"))
	require.Error(t, validateHairColour("so #abcdEF"))
	require.Error(t, validateHairColour("#abcdEF os"))
}

func TestValidateEyeColour(t *testing.T) {
	require.NoError(t, validateEyeColour("grn"))
	require.NoError(t, validateEyeColour("brn"))
	require.Error(t, validateEyeColour("bru"))
	require.Error(t, validateEyeColour("erg"))
}

func TestValidatePassportID(t *testing.T) {
	require.NoError(t, validateHPassportID("123456789"))
	require.Error(t, validateHPassportID("1234567891"))
	require.Error(t, validateHPassportID("123d56789"))
}
