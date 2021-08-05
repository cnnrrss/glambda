package main

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"
)

type SuccessResponse struct {
	Success bool `json:"success"`
}

func handler(_ context.Context, request json.RawMessage) (interface{}, error) {
	log.Info().Str("env", environment).Interface("request", request).Msg("lambda invoke request body")
	response := SuccessResponse{Success: true}
	b, _ := json.Marshal(response)
	return b, nil
}