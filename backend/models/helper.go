package models

import "context"

type PaginationFilter struct {
	PageNumber uint
	PageSize   uint
}

type PaginationMetadata struct {
	TotalRecords    uint `json:"total_records"`
	TotalPages      uint `json:"total_pages"`
	CurrentLimit    uint `json:"page_size"`
	CurrentPage     uint `json:"current_page"`
	HasNextPage     bool `json:"has_next_page"`
	HasPreviousPage bool `json:"has_previous_page"`
}

func CalculateMetadata(ctx context.Context, totalRecords int64, metadata *PaginationMetadata, filter *PaginationFilter) error {
	metadata.TotalRecords = uint(totalRecords)
	metadata.CurrentLimit = filter.PageSize
	metadata.CurrentPage = filter.PageNumber
	metadata.TotalPages = (metadata.TotalRecords + metadata.CurrentLimit - 1) / metadata.CurrentLimit
	metadata.HasNextPage = (metadata.CurrentPage < metadata.TotalPages)
	metadata.HasPreviousPage = (metadata.CurrentPage > 1)

	return nil
}
