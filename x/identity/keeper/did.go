package keeper

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// CreateDIDHandler creates a DID using the walt.id API
func (k Keeper) CreateDIDHandler(ctx context.Context, walletUUID string) (string, error) {
	baseURL := "https://wallet.walt.id"
	url := fmt.Sprintf("%s/wallet-api/wallet/%s/dids/create/jwk", baseURL, walletUUID)

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Log the request details
	log.Printf("Sending request: %s %s", req.Method, req.URL)
	for key, value := range req.Header {
		log.Printf("Header: %s=%s", key, value)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	log.Printf("Received response status: %s", resp.Status)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Response body: %s", string(body))
		return "", fmt.Errorf("received non-200 response status: %s, body: %s", resp.Status, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	log.Printf("Response body: %s", string(body))

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	did, ok := result["did"].(string)
	if !ok {
		return "", fmt.Errorf("unexpected response format: %s", string(body))
	}

	return did, nil
}
