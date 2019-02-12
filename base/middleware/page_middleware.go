package middleware

import (
    "strconv"
    "context"
	"net/http"
)

// Middleware used to extract the `page` paramter from the URL and save it
// in the context.
func PaginationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Setup our variables for the paginator.
        pageString := r.FormValue("page")
        pageIndex, err := strconv.ParseUint(pageString, 10, 64)
        if err != nil {
            pageIndex = 0
        }

        paginateByString := r.FormValue("paginate_by")
        paginateBy, err := strconv.ParseUint(paginateByString, 10, 64)
        if err != nil {
            paginateBy = 25 // Default value.
        }
        if paginateBy > 100 {
            http.Error(w, "Please pick a pagination number less then or equal to 100.", http.StatusBadRequest)
    		return
        }

        // Attach the 'page' & 'paginate_by' parameter values to our context
        // to be used.
		ctx := context.WithValue(r.Context(), "pageIndex", pageIndex)
        ctx = context.WithValue(ctx, "paginateBy", paginateBy)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
