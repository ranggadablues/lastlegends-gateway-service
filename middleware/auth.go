package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ranggadablues/gosok/auth"
	"github.com/ranggadablues/gosok/common"
	"github.com/ranggadablues/lastlegends-proto-library/user-proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		claims, err := auth.ValidateAccessToken(token)
		fmt.Println("cek claims >>", common.ToJSON(claims), "==>", err)
		if err != nil {
			// Handle token expiration
			if err == auth.ErrTokenExpired || err.Error() == "token is expired" {
				refreshToken := r.Header.Get("X-Refresh-Token")
				newAccess, newRefresh, err := claimCheck(refreshToken)
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}

				// Add new tokens to response header for client to update locally
				w.Header().Set("X-New-Access-Token", newAccess)
				w.Header().Set("X-New-Refresh-Token", newRefresh)

				// Optionally, continue request using new access token
				r.Header.Set("Authorization", "Bearer "+newAccess)
				// proceed to next handler
				next.ServeHTTP(w, r)
				return
			}

			http.Error(w, "invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// ✅ token is valid
		fmt.Println("cek claims 2 >>", common.ToJSON(claims))
		ctx := context.WithValue(r.Context(), auth.ClaimsContextKey, claims) // you can attach claims to context if needed
		fmt.Println("cek ctx >>", ctx)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func claimCheck(refreshToken string) (string, string, error) {
	if refreshToken == "" {
		return "", "", common.Error("token expired and no refresh token provided", nil)
	}

	newAccess, newRefresh, err := refreshTokens(refreshToken)
	if err != nil {
		return "", "", common.Error("refresh failed", err)
	}

	return newAccess, newRefresh, nil
}

// --------------------------
// Helper: call user-service.RefreshToken()
// --------------------------
func refreshTokens(refreshToken string) (string, string, error) {
	addr := os.Getenv("USER_SERVICE_PORT")
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", "", err
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	resp, err := client.RefreshToken(context.Background(), &pb.RefreshRequest{RefreshToken: refreshToken})
	if err != nil {
		return "", "", err
	}

	return resp.Token, resp.RefreshToken, nil
}
