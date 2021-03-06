package middleware

import (
    "context"
	"net/http"
    
    "github.com/mikasoftware/mikapost-go/model_manager"
	"github.com/mikasoftware/mikapost-go/base/service"
)

func ProfileCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user_id := service.GetUserIDFromContext(r.Context())
    	if user_id == 0 {
    		http.Error(w, "User ID not inputted.", http.StatusUnauthorized)
    		return
    	}
        user, count := model_manager.UserManagerInstance().GetByID(user_id)
        if count == 0  {
    		http.Error(w, "No User found with ID.", http.StatusUnauthorized)
    		return
    	}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
