package fallbackroutes

//encore:service
// type Service struct {
//     oldRounter *
// }

// encore:api public path=/!fallback
// func (s *Service) Fallback(w http.ResponseWriter, req *http.Request) *rrt.Response {
// 	query := req.URL.Query().Get("name")
// 	ctx := context.Background()
// 	rsp, err := ping.Ping(ctx, &ping.PingParams{Name: query})
// 	if err != nil {
// 		http.Error(w, "Failed to rount request: "+err.Error(), http.StatusInternalServerError)
// 		return &rrt.Response{Status: "Fail"}
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(rsp)
// 	return &rrt.Response{Status: "S"}
// }
