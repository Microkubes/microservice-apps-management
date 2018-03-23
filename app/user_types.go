// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "apps-management": Application User Types
//
// Command:
// $ goagen
// --design=github.com/Microkubes/microservice-apps-management/design
// --out=$(GOPATH)/src/github.com/Microkubes/microservice-apps-management
// --version=v1.2.0-dirty

package app

import (
	"github.com/goadesign/goa"
	"unicode/utf8"
)

// App ID+secret credentials
type appCredentialsPayload struct {
	// The app ID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// The app secret
	Secret *string `form:"secret,omitempty" json:"secret,omitempty" xml:"secret,omitempty"`
}

// Validate validates the appCredentialsPayload type instance.
func (ut *appCredentialsPayload) Validate() (err error) {
	if ut.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "id"))
	}
	if ut.Secret == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "secret"))
	}
	return
}

// Publicize creates AppCredentialsPayload from appCredentialsPayload
func (ut *appCredentialsPayload) Publicize() *AppCredentialsPayload {
	var pub AppCredentialsPayload
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	if ut.Secret != nil {
		pub.Secret = *ut.Secret
	}
	return &pub
}

// App ID+secret credentials
type AppCredentialsPayload struct {
	// The app ID
	ID string `form:"id" json:"id" xml:"id"`
	// The app secret
	Secret string `form:"secret" json:"secret" xml:"secret"`
}

// Validate validates the AppCredentialsPayload type instance.
func (ut *AppCredentialsPayload) Validate() (err error) {
	if ut.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "id"))
	}
	if ut.Secret == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "secret"))
	}
	return
}

// Payload for the client apps
type appPayload struct {
	// Description of the app
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// App domain
	Domain *string `form:"domain,omitempty" json:"domain,omitempty" xml:"domain,omitempty"`
	// Name of the app
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the appPayload type instance.
func (ut *appPayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Description != nil {
		if utf8.RuneCountInString(*ut.Description) > 300 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.description`, *ut.Description, utf8.RuneCountInString(*ut.Description), 300, false))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 50, false))
		}
	}
	return
}

// Publicize creates AppPayload from appPayload
func (ut *appPayload) Publicize() *AppPayload {
	var pub AppPayload
	if ut.Description != nil {
		pub.Description = ut.Description
	}
	if ut.Domain != nil {
		pub.Domain = ut.Domain
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	return &pub
}

// Payload for the client apps
type AppPayload struct {
	// Description of the app
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// App domain
	Domain *string `form:"domain,omitempty" json:"domain,omitempty" xml:"domain,omitempty"`
	// Name of the app
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the AppPayload type instance.
func (ut *AppPayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Description != nil {
		if utf8.RuneCountInString(*ut.Description) > 300 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`type.description`, *ut.Description, utf8.RuneCountInString(*ut.Description), 300, false))
		}
	}
	if utf8.RuneCountInString(ut.Name) > 50 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.name`, ut.Name, utf8.RuneCountInString(ut.Name), 50, false))
	}
	return
}
