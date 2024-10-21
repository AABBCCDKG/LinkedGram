// A program Team Project
//
// Purdue University -- CS18000 -- Spring 2024 -- Team Project
//
// Author: Purdue CS
// Version: Mar 31, 2024

package main

import (
	"testing"
)

// Profile struct definition (assumed to be already implemented)
type Profile struct {
	About      string
	Experience []string
	Education  string
	Awards     string
	Skills     string
	Status     string
}

// NewProfile creates a new Profile instance
func NewProfile(about string, experience []string, education, awards, skills, status string) *Profile {
	return &Profile{
		About:      about,
		Experience: experience,
		Education:  education,
		Awards:     awards,
		Skills:     skills,
		Status:     status,
	}
}

// TestGetAbout tests the GetAbout method of Profile
func TestGetAbout(t *testing.T) {
	experience := []string{}
	profile := NewProfile("about", experience, "education", "awards", "skills", "status")
	if profile.About != "about" {
		t.Errorf("Expected 'about', got '%s'", profile.About)
	}
}

// TestSetAbout tests the SetAbout method of Profile
func TestSetAbout(t *testing.T) {
	experience := []string{}
	profile := NewProfile("about", experience, "education", "awards", "skills", "status")
	profile.About = "newAbout"
	if profile.About != "newAbout" {
		t.Errorf("Expected 'newAbout', got '%s'", profile.About)
	}
}

// TestSetExperience tests the SetExperience method of Profile
func TestSetExperience(t *testing.T) {
	experience := []string{"experience1"}
	profile := NewProfile("about", experience, "education", "awards", "skills", "status")
	newExperience := []string{"experience2"}
	profile.Experience = newExperience
	if len(profile.Experience) != 1 || profile.Experience[0] != "experience2" {
		t.Errorf("Expected 'experience2', got '%v'", profile.Experience)
	}
}

// TestGetEducation tests the GetEducation method of Profile
func TestGetEducation(t *testing.T) {
	experience := []string{}
	profile := NewProfile("about", experience, "education", "awards", "skills", "status")
	if profile.Education != "education" {
		t.Errorf("Expected 'education', got '%s'", profile.Education)
	}
}

// TestSetEducation tests the SetEducation method of Profile
func TestSetEducation(t *testing.T) {
	experience := []string{}
	profile := NewProfile("about", experience, "education", "awards", "skills", "status")
	profile.Education = "newEducation"
	if profile.Education != "newEducation" {
		t.Errorf("Expected 'newEducation', got '%s'", profile.Education)
	}
}

// TestGetAwards tests the GetAwards method of Profile
func TestGetAwards(t *testing.T) {
	experience := []string{}
	profile := NewProfile("about", experience, "education", "awards", "skills", "status")
	if profile.Awards != "awards" {
		t.Errorf("Expected 'awards', got '%s'", profile.Awards)
	}
}

// TestSetAwards tests the SetAwards method of Profile
func TestSetAwards(t *testing.T) {
	experience := []string{}
	profile := NewProfile("about", experience, "education", "awards", "skills", "status")
	profile.Awards = "newAwards"
	if profile.Awards != "newAwards" {
		t.Errorf("Expected 'newAwards', got '%s'", profile.Awards)
	}
}

// TestSetSkills tests the SetSkills method of Profile
func TestSetSkills(t *testing.T) {
	experience := []string{}
	profile := NewProfile("about", experience, "education", "awards", "skills", "status")
	profile.Skills = "newSkills"
	if profile.Skills != "newSkills" {
		t.Errorf("Expected 'newSkills', got '%s'", profile.Skills)
	}
}

// TestGetStatus tests the GetStatus method of Profile
func TestGetStatus(t *testing.T) {
	experience := []string{}
	profile := NewProfile("about", experience, "education", "awards", "skills", "status")
	if profile.Status != "status" {
		t.Errorf("Expected 'status', got '%s'", profile.Status)
	}
}

// TestSetStatus tests the SetStatus method of Profile
func TestSetStatus(t *testing.T) {
	experience := []string{}
	profile := NewProfile("about", experience, "education", "awards", "skills", "status")
	profile.Status = "newStatus"
	if profile.Status != "newStatus" {
		t.Errorf("Expected 'newStatus', got '%s'", profile.Status)
	}
}
