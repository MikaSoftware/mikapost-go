package controller

import (
    "context"
    "net/http"
    "strconv"
    "github.com/go-chi/chi"
	"github.com/go-chi/render"

    "github.com/mikasoftware/mikapost-go/model"
	"github.com/mikasoftware/mikapost-go/model_manager"
    common_s "github.com/mikasoftware/mikapost-go/serializer"
    box_s "github.com/mikasoftware/mikapost-go/serializer/box_serializer"
)


//----------------------------------------------------------------------------//
//                                  CREATE                                    //
//----------------------------------------------------------------------------//


func CreateBoxFunc(w http.ResponseWriter, r *http.Request) {
    // Take the user POST data and serialize it.
    data := &box_s.BoxCreateRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, common_s.ErrInvalidRequest(err))
		return
	}

    // Create our user model in our database from the serialized data.
    box, _ := data.Save(r.Context())

    // Take newly created Box model data object and serialize it
    // to be returned as the result for this API endpoint.
    render.Status(r, http.StatusCreated)
	render.Render(w, r, box_s.NewBoxDetailResponse(box))
}


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//


func PaginatedBoxListCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract variables from our context.
        pageIndex := r.Context().Value("pageIndex").(uint64)
        paginateBy := r.Context().Value("paginateBy").(uint64)
        user := r.Context().Value("user").(*model.User)

        // Fetch all the model objects in our data based on user context.
        boxes, pagesCount, totalRecords := model_manager.BoxManagerInstance().FilterBy(user, pageIndex, paginateBy)

        // Validate our URL parameters.
        if pageIndex >= pagesCount {
            render.Render(w, r, common_s.ErrNotFound)
            return
        }

        // Attach to our context.
        ctx := context.WithValue(r.Context(), "boxes", boxes)
        ctx = context.WithValue(ctx, "pagesCount", pagesCount)
        ctx = context.WithValue(ctx, "totalRecords", totalRecords)
        next.ServeHTTP(w, r.WithContext(ctx))
        return
	})
}


func ListBoxesFunc(w http.ResponseWriter, r *http.Request) {
    // Extract from the context our model data and URL parameters.
    pageIndex := r.Context().Value("pageIndex").(uint64)
    pagesCount := r.Context().Value("pagesCount").(uint64)
    totalRecords := r.Context().Value("totalRecords").(uint64)
    boxes := r.Context().Value("boxes").([]model.Box)

    // Render the controller's JSON output.
    if err := render.Render(w, r, box_s.NewBoxListResponse(boxes, pageIndex, pagesCount, totalRecords)); err != nil {
		render.Render(w, r, common_s.ErrRender(err))
		return
	}
}


//----------------------------------------------------------------------------//
//                                 RETRIEVE                                   //
//----------------------------------------------------------------------------//


// Middleware will extract the `boxID` parameter from the URL and
// attempt to lookup the Box model data object in the database. If
// the object was found then attach it to the context, else return an error.
func BoxCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if boxIDString := chi.URLParam(r, "boxID"); boxIDString != "" {
			boxID, _ := strconv.ParseUint(boxIDString, 10, 64)
			box, count := model_manager.BoxManagerInstance().GetByID(boxID)
            if count == 1 {
                // Confirm the authenticated user.
                user := r.Context().Value("user").(*model.User)
                hasPermission := model_manager.BoxManagerInstance().HasPermission(user, boxID)
                if hasPermission == false {
                    render.Render(w, r, common_s.ErrForbidden)
                    return
                }

                // Attach the box to the context.
                ctx := context.WithValue(r.Context(), "box", box)
        		next.ServeHTTP(w, r.WithContext(ctx))
                return
            }
		}
        render.Render(w, r, common_s.ErrNotFound)
        return
	})
}


func RetrieveBoxFunc(w http.ResponseWriter, r *http.Request) {
    // Extract from the context our model data.
    box := r.Context().Value("box").(*model.Box)

    // Render the controller's JSON output.
	if err := render.Render(w, r, box_s.NewBoxDetailResponse(box)); err != nil {
		render.Render(w, r, common_s.ErrRender(err))
		return
	}
}
