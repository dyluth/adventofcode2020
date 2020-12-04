package passport

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	EyeColour      *string `json:"ecl"`
	PassportID     *string `json:"pid"`
	ExpirationYear *string `json:"eyr"`
	HairColor      *string `json:"hcl"`
	BirthYear      *string `json:"byr"`
	IssueYear      *string `json:"iyr"`
	CountryID      *string `json:"cid"`
	Height         *string `json:"hgt"`
}

func (p *Passport) Validate() error {
	err := validateNumber(4, 1920, 2002, *p.BirthYear)
	if err != nil {
		return err
	}
	err = validateNumber(4, 2010, 2020, *p.IssueYear)
	if err != nil {
		return err
	}
	err = validateNumber(4, 2020, 2030, *p.ExpirationYear)
	if err != nil {
		return err
	}
	err = validateHeight(*p.Height)
	if err != nil {
		return err
	}
	err = validateHairColour(*p.HairColor)
	if err != nil {
		return err
	}
	err = validateEyeColour(*p.EyeColour)
	if err != nil {
		return err
	}
	err = validateHPassportID(*p.PassportID)
	if err != nil {
		return err
	}
	return nil
}

func validateHPassportID(in string) error {
	re := regexp.MustCompile(`^\d\d\d\d\d\d\d\d\d$`)
	if re.MatchString(in) {
		return nil
	}
	return fmt.Errorf("%v is not a valid colour", in)

}

func validateEyeColour(in string) error {
	eyes := make(map[string]bool)
	eyes["amb"] = true
	eyes["blu"] = true
	eyes["brn"] = true
	eyes["gry"] = true
	eyes["grn"] = true
	eyes["hzl"] = true
	eyes["oth"] = true

	_, ok := eyes[in]
	if !ok {
		return fmt.Errorf("NOT EYES")
	}
	return nil
}

func validateHairColour(in string) error {
	re := regexp.MustCompile(`^#\S\S\S\S\S\S$`)
	if re.MatchString(in) {
		return nil
	}
	return fmt.Errorf("%v is not a valid colour", in)

}

func validateHeight(in string) error {
	re := regexp.MustCompile(`(\d+)(cm|in)`)
	match := re.FindStringSubmatch(in)
	if len(match) != 3 {
		return fmt.Errorf("not the right number of parts (found %v)", len(match))
	}

	switch match[2] {
	case "cm":
		return validateNumber(3, 150, 193, match[1])
	case "in":
		return validateNumber(2, 59, 77, match[1])
	default:
		return fmt.Errorf("%v not in or cm", match[2])
	}
}

func validateNumber(digits, min, max int, in string) error {
	if len(in) != digits {
		return fmt.Errorf("wrong number of digits (%v, should be %v)", len(in), digits)
	}
	value, err := strconv.Atoi(in)
	if err != nil {
		return err
	}
	return validateRange(min, max, value)
}

func validateRange(min, max, value int) error {
	if value > max || value < min {
		return fmt.Errorf("value %v not in range [%v to %v]", value, min, max)
	}
	return nil
}

func ParsePassport(in string) (Passport, error) {

	// convert string into yaml format
	// remove newlines
	in = strings.ReplaceAll(in, "\n", " ")
	// split on whitespace
	lines := strings.Fields(in)
	newLines := []string{}
	for _, line := range lines {
		l := strings.TrimSpace(line)
		parts := strings.Split(l, ":")
		if len(parts) != 2 {
			return Passport{}, fmt.Errorf("not nice pair: %v", l)
		}
		newLines = append(newLines, fmt.Sprintf(`"%v":"%v"`, parts[0], parts[1]))

	}
	// join with newlines
	in = fmt.Sprintf("{%v}", strings.Join(newLines, ","))
	fmt.Printf("lines:\n %+v\n", in)
	// parse as yaml
	pass := Passport{}
	err := json.Unmarshal([]byte(in), &pass)
	// yaml.Unmarshal(in, &pass)

	// check for mandatory fields
	if pass.BirthYear == nil || pass.ExpirationYear == nil || pass.EyeColour == nil || pass.HairColor == nil || pass.Height == nil || pass.IssueYear == nil || pass.PassportID == nil {
		return pass, fmt.Errorf("missing mandatory field")
	}

	return pass, err
}

func SplitInput(input []string) []string {
	inputString := strings.Join(input, "\n")
	// split on empty lines
	re := regexp.MustCompile(`\n\s*\n`)
	passList := re.Split(inputString, -1)
	return passList

}
