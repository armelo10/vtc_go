package httpapi
import("encoding/json";"net/http";"time";"github.com/armelo10/vtc_go/backend/internal/infrastructure/config")
type Server struct{cfg config.Config;mux *http.ServeMux}
func NewServer(cfg config.Config)*Server{s:=&Server{cfg:cfg,mux:http.NewServeMux()};s.mux.HandleFunc("GET /healthz",func(w http.ResponseWriter,_ *http.Request){writeJSON(w,200,map[string]any{"status":"ok","service":"vtc-api"})});s.mux.HandleFunc("GET /readyz",func(w http.ResponseWriter,_ *http.Request){writeJSON(w,200,map[string]any{"status":"ready","environment":cfg.Environment})});s.mux.HandleFunc("GET /api/v1",func(w http.ResponseWriter,_ *http.Request){writeJSON(w,200,map[string]string{"version":"v1"})});return s}
func(s *Server)Handler()http.Handler{return requestID(s.mux)}
func writeJSON(w http.ResponseWriter,status int,v any){w.Header().Set("Content-Type","application/json");w.WriteHeader(status);_=json.NewEncoder(w).Encode(v)}
func requestID(next http.Handler)http.Handler{return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){id:=r.Header.Get("X-Request-ID");if id==""{id=time.Now().UTC().Format("20060102T150405.000000000Z")};w.Header().Set("X-Request-ID",id);next.ServeHTTP(w,r)})}
