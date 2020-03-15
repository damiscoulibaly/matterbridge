// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "time"

// SiteActivitySummary undocumented
type SiteActivitySummary struct {
	// Entity is the base model of SiteActivitySummary
	Entity
	// ReportRefreshDate undocumented
	ReportRefreshDate *time.Time `json:"reportRefreshDate,omitempty"`
	// ViewedOrEdited undocumented
	ViewedOrEdited *int `json:"viewedOrEdited,omitempty"`
	// Synced undocumented
	Synced *int `json:"synced,omitempty"`
	// SharedInternally undocumented
	SharedInternally *int `json:"sharedInternally,omitempty"`
	// SharedExternally undocumented
	SharedExternally *int `json:"sharedExternally,omitempty"`
	// ReportDate undocumented
	ReportDate *time.Time `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
}