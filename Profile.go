package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Profile represents a user's profile
type Profile struct {
	About      string   `json:"about"`
	Experience []string `json:"experience"`
	Education  string   `json:"education"`
	Awards     string   `json:"awards"`
	Skills     string   `json:"skills"`
	Status     string   `json:"status"`
}

// NewProfile creates a new Profile and validates it
func NewProfile(about, education, awards, skills, status string, experience []string) (*Profile, error) {
	prof := &Profile{
		About:      about,
		Experience: experience,
		Education:  education,
		Awards:     awards,
		Skills:     skills,
		Status:     status,
	}
	if err := prof.validateProfile(); err != nil {
		return nil, err
	}
	return prof, nil
}

// validateProfile checks if the profile is complete
func (p *Profile) validateProfile() error {
	if p.About == "" || p.Education == "" || p.Awards == "" || p.Skills == "" || p.Status == "" {
		return errors.New("profile is incomplete")
	}
	return nil
}

// String returns a string representation of the profile
func (p *Profile) String() string {
	return fmt.Sprintf("About: %s\nExperience: %v\nEducation: %s\nAwards: %s\nSkills: %s\nStatus: %s",
		p.About, p.Experience, p.Education, p.Awards, p.Skills, p.Status)
}

// MarshalJSON allows custom serialization of the Profile
func (p *Profile) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		About      string   `json:"about"`
		Experience []string `json:"experience"`
		Education  string   `json:"education"`
		Awards     string   `json:"awards"`
		Skills     string   `json:"skills"`
		Status     string   `json:"status"`
	}{
		About:      p.About,
		Experience: p.Experience,
		Education:  p.Education,
		Awards:     p.Awards,
		Skills:     p.Skills,
		Status:     p.Status,
	})
}

// UnmarshalJSON allows custom deserialization of the Profile
func (p *Profile) UnmarshalJSON(data []byte) error {
	temp := struct {
		About      string   `json:"about"`
		Experience []string `json:"experience"`
		Education  string   `json:"education"`
		Awards     string   `json:"awards"`
		Skills     string   `json:"skills"`
		Status     string   `json:"status"`
	}{}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	p.About = temp.About
	p.Experience = temp.Experience
	p.Education = temp.Education
	p.Awards = temp.Awards
	p.Skills = temp.Skills
	p.Status = temp.Status
	if err := p.validateProfile(); err != nil {
		return err
	}
	return nil
}

func main() {
	prof, err := NewProfile("I am a software developer.", "Purdue University", "Dean's List", "Go, Java", "Active", []string{"Software Developer at XYZ"})
	if err != nil {
		fmt.Printf("Error creating profile: %v\n", err)
		return
	}
	fmt.Println(prof)

	jsonData, err := json.Marshal(prof)
	if err != nil {
		fmt.Printf("Error marshaling profile: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))

	newProf := &Profile{}
	if err := json.Unmarshal(jsonData, newProf); err != nil {
		fmt.Printf("Error unmarshaling profile: %v\n", err)
		return
	}
	fmt.Println(newProf)
}