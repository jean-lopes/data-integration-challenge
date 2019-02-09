package companies

import (
	"errors"
	"net/url"
	"regexp"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// Company entity
type Company struct {
	ID      *uuid.UUID `json:"id,omitempty"`
	Name    string     `json:"name"`
	Zip     string     `json:"zip"`
	Website *string    `json:"website,omitempty"`
}

var (
	// ErrEmptyName empty name error
	ErrEmptyName = errors.New("Empty company name")
	// ErrEmptyZip empty zip error
	ErrEmptyZip = errors.New("Empty company zip code")
	// ErrInvalidZip Invalid zip error
	ErrInvalidZip = errors.New("Company zip code must have exacly 5 (five) digits")
	// ErrInvalidWebsite malformed URI
	ErrInvalidWebsite = errors.New("Invalid website")
	// Nil represents an Empty company
	Nil   = Company{}
	zipRE = regexp.MustCompile("[0-9]{5}")
)

// Validate business constraints
func (company Company) Validate() []error {
	nameError := validateName(company.Name)
	zipError := validateZip(company.Zip)
	websiteError := validateWebsite(company.Website)

	errors := make([]error, 0)
	errors = appendError(errors, nameError)
	errors = appendError(errors, zipError)
	errors = appendError(errors, websiteError)

	if len(errors) == 0 {
		return nil
	}

	return errors
}

// HasID verify if the company has ID
func (company Company) HasID() bool {
	return company.ID != nil && !uuid.Equal(*company.ID, uuid.Nil)
}

// IsEmpty checks if an instance of company is empty
func (company Company) IsEmpty() bool {
	return company.ID == nil &&
		company.Name == "" &&
		company.Zip == "" &&
		company.Website == nil
}

// Equal Checks if two companies instances have the same ID
func (company Company) Equal(other Company) bool {
	return company.ID != nil && other.ID != nil && uuid.Equal(*company.ID, *other.ID)
}

func isBlank(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func validateName(name string) error {
	if isBlank(name) {
		return ErrEmptyName
	}

	return nil
}

func validateZip(zip string) error {
	if isBlank(zip) {
		return ErrEmptyZip
	}

	if !zipRE.MatchString(zip) {
		return ErrInvalidZip
	}

	return nil
}

func validateWebsite(website *string) error {
	if website != nil && !isBlank(*website) {
		_, err := url.ParseRequestURI(*website)
		if err != nil {
			return ErrInvalidWebsite
		}
	}

	return nil
}

func appendError(errors []error, err error) []error {
	if err != nil {
		errors = append(errors, err)
	}

	return errors
}