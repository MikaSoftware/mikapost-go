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
    tsd_s "github.com/mikasoftware/mikapost-go/serializer/time_series_datum_serializer"
)


//----------------------------------------------------------------------------//
//                                  CREATE                                    //
//----------------------------------------------------------------------------//


func CreateTimeSeriesDatumFunc(w http.ResponseWriter, r *http.Request) {
    // Take the user POST data and serialize it.
    data := &tsd_s.TimeSeriesDatumCreateRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, common_s.ErrInvalidRequest(err))
		return
	}

    // Create our user model in our database from the serialized data.
    thing, _ := data.Save(r.Context())

    // Take newly created TimeSeriesDatum model data object and serialize it
    // to be returned as the result for this API endpoint.
    render.Status(r, http.StatusCreated)
	render.Render(w, r, tsd_s.NewTimeSeriesDatumDetailResponse(thing))
}


//----------------------------------------------------------------------------//
//                                  LIST                                      //
//----------------------------------------------------------------------------//


func ThingDataListCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if thingIDString := chi.URLParam(r, "thingID"); thingIDString != "" {
            // Extract URL Param.
			thingID, err := strconv.ParseUint(thingIDString, 10, 64)

            // If the user inputted a URL param which cannot be converted
            // to an integer then we generate a 404 error.
            if err != nil {
                render.Render(w, r, common_s.ErrNotFound)
                return
            }

            // Extract variables from our context.
            pageIndex := r.Context().Value("pageIndex").(uint64)
            paginateBy := r.Context().Value("paginateBy").(uint64)
            user := r.Context().Value("user").(*model.User)

            // Enforce permission handling for the time series data.
            hasPermission := model_manager.TimeSeriesDatumManagerInstance().HasPermission(user.ID, thingID)
            if hasPermission == false {
                render.Render(w, r, common_s.ErrForbidden)
                return
            }

            // Fetch all the model objects in our data based on user context.
            timeSeriesData, pagesCount, totalRecords := model_manager.TimeSeriesDatumManagerInstance().FilterByThing(thingID, pageIndex, paginateBy)

            // Validate our URL parameters.
            if pageIndex >= pagesCount {
                render.Render(w, r, common_s.ErrNotFound)
                return
            }

            // Attach our model data & URL parameters to the context.
            ctx := context.WithValue(r.Context(), "timeSeriesData", timeSeriesData)
            ctx = context.WithValue(ctx, "pagesCount", pagesCount)
            ctx = context.WithValue(ctx, "totalRecords", totalRecords)
            next.ServeHTTP(w, r.WithContext(ctx))
            return
        }
        render.Render(w, r, common_s.ErrNotFound)
        return
	})
}


func ListThingTimeSeriesDataFunc(w http.ResponseWriter, r *http.Request) {
    // Extract from the context our model data and URL parameters.
    pageIndex := r.Context().Value("pageIndex").(uint64)
    pagesCount := r.Context().Value("pagesCount").(uint64)
    totalRecords := r.Context().Value("totalRecords").(uint64)
    timeSeriesData := r.Context().Value("timeSeriesData").([]model.TimeSeriesDatum)

    // Render the controller's JSON output.
    if err := render.Render(w, r, tsd_s.NewTimeSeriesDataListResponse(timeSeriesData, pageIndex, pagesCount, totalRecords)); err != nil {
		render.Render(w, r, common_s.ErrRender(err))
		return
	}
}


//----------------------------------------------------------------------------//
//                                 RETRIEVE                                   //
//----------------------------------------------------------------------------//


// Middleware will extract the `tsdID` parameter from the URL and
// attempt to lookup the TimeSeriesDatum model data object in the database. If
// the object was found then attach it to the context, else return an error.
func TimeSeriesDatumCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if tsdIDString := chi.URLParam(r, "tsdID"); tsdIDString != "" {
			tsdID, _ := strconv.ParseUint(tsdIDString, 10, 64)
			timeSeriesDatum, count := model_manager.TimeSeriesDatumManagerInstance().GetByID(tsdID)
            if count == 1 {

                user := r.Context().Value("user").(*model.User)
                hasPermission := model_manager.TimeSeriesDatumManagerInstance().HasPermission(user.ID, timeSeriesDatum.ThingID)
                if hasPermission == false {
                    render.Render(w, r, common_s.ErrForbidden)
                    return
                }

                // Attach the timeSeriesDatum to the context.
                ctx := context.WithValue(r.Context(), "timeSeriesDatum", timeSeriesDatum)
        		next.ServeHTTP(w, r.WithContext(ctx))
                return
            }
		}
        render.Render(w, r, common_s.ErrNotFound)
        return
	})
}


func RetrieveTimeSeriesDatumFunc(w http.ResponseWriter, r *http.Request) {
    timeSeriesDatum := r.Context().Value("timeSeriesDatum").(*model.TimeSeriesDatum)

	if err := render.Render(w, r, tsd_s.NewTimeSeriesDatumDetailResponse(timeSeriesDatum)); err != nil {
		render.Render(w, r, common_s.ErrRender(err))
		return
	}
}
