package models

// Profile is struct of profile
type User struct {
	Id        string `json:"id"`
	DispName  string `json:"disp_name"`
	PassHash  string `json:"pass_hash"`
	Email     string `json:"email"`
	Authority string `json:"authority"`
	CompanyId string `json:"company_id"`
}

/*
// GetProfile returns a profile
func GetProfile() *Profile {
	return &Profile{
		FamilyNameKanji: "道祖",
		GivenNameKanji:  "克理",
		Location:        "Tokyo",
	}
}

// FullNameKanji is full name in Japanese
func (p *Profile) FullNameKanji() string {
	return fmt.Sprintf("%s %s", p.FamilyNameKanji, p.GivenNameKanji)
}

// FullNameKana is kana in Japanese
func (p *Profile) FullNameKana() string {
	return fmt.Sprintf("%s %s", p.FamilyNameKana, p.GivenNameKana)
}

// FullNameEn is full name in English
func (p *Profile) FullNameEn() string {
	return fmt.Sprintf("%s %s", p.GivenNameEn, p.FamilyNameEn)
}
*/
