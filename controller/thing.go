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
    thing_s "github.com/mikasoftware/mikapost-go/serializer/thing_serializer"
)


//----------------------------------------------------------------------------//
//                                  CREATE                                    //
//----------------------------------------------------------------------------//


func CreateThingFunc(w http.ResponseWriter, r *http.Request) {
    // Take the user POST data and serialize it.
    data := &thing_s.ThingCreateRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, common_s.ErrInvalidRequest(err))
		return
	}

    // Create our user model in our database from the serialized data.
    thing, _ := data.Save(r.Context())

    // Take newly created Thing model data object and serialize it
    // to be returned as the result for this API endpoint.
    render.Status(r, http.StatusCreated)
	render.Render(w, r, thing_s.NewThingDetailResponse(thing))
}


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//


func PaginatedThingListCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract variables from our context.
        pageIndex := r.Context().Value("pageIndex").(uint64)
        paginateBy := r.Context().Value("paginateBy").(uint64)
        user := r.Context().Value("user").(*model.User)

        // Fetch all the model objects in our data based on user context.
        things, pagesCount, totalRecords := model_manager.ThingManagerInstance().FilterByUser(user, pageIndex, paginateBy)

        // Validate our URL parameters.
        if pageIndex >= pagesCount {
            render.Render(w, r, common_s.ErrNotFound)
            return
        }

        ctx := context.WithValue(r.Context(), "things", things)
        ctx = context.WithValue(ctx, "pagesCount", pagesCount)
        ctx = context.WithValue(ctx, "totalRecords", totalRecords)
        next.ServeHTTP(w, r.WithContext(ctx))
        return
	})
}


func ListThingsFunc(w http.ResponseWriter, r *http.Request) {
    // Extract from the context our model data and URL parameters.
    pageIndex := r.Context().Value("pageIndex").(uint64)
    pagesCount := r.Context().Value("pagesCount").(uint64)
    totalRecords := r.Context().Value("totalRecords").(uint64)
    things := r.Context().Value("things").([]model.Thing)

    // Render the controller's JSON output.
    if err := render.Render(w, r, thing_s.NewThingListResponse(things, pageIndex, pagesCount, totalRecords)); err != nil {
		render.Render(w, r, common_s.ErrRender(err))
		return
	}
}


//----------------------------------------------------------------------------//
//                                 RETRIEVE                                   //
//----------------------------------------------------------------------------//


// Middleware will extract the `thingID` parameter from the URL and
// attempt to lookup the Thing model data object in the database. If
// the object was found then attach it to the context, else return an error.
func ThingCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if thingIDString := chi.URLParam(r, "thingID"); thingIDString != "" {
			thingID, _ := strconv.ParseUint(thingIDString, 10, 64)
			thing, count := model_manager.ThingManagerInstance().GetByID(thingID)
            if count == 1 {

                user := r.Context().Value("user").(*model.User)
                hasPermission := model_manager.ThingManagerInstance().HasThingPermission(user, thingID)
                if hasPermission == false {
                    render.Render(w, r, common_s.ErrForbidden)
                    return
                }

                // Attach the thing to the context.
                ctx := context.WithValue(r.Context(), "thing", thing)
        		next.ServeHTTP(w, r.WithContext(ctx))
                return
            }
		}
        render.Render(w, r, common_s.ErrNotFound)
        return
	})
}


func RetrieveThingFunc(w http.ResponseWriter, r *http.Request) {
    thing := r.Context().Value("thing").(*model.Thing)

	if err := render.Render(w, r, thing_s.NewThingDetailResponse(thing)); err != nil {
		render.Render(w, r, common_s.ErrRender(err))
		return
	}
}
