package middleware

import (
	"net/http"
	
    "github.com/mikasoftware/mikapost-go/model"
)

func StaffCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract the current user from the request context. It is important
        // to note that this context must be AFTER the `ProfileCtx` middleware.
        ctx := r.Context()
        user := ctx.Value("user").(*model.User)

        // Verify the user belongs to the staff group.
        if user.GroupID != 2  {
    		http.Error(w, "You are not a staff member.", http.StatusUnauthorized)
    		return
    	}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
