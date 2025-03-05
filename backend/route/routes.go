package route

import (
	"net/http"

	"github.com/Githaiga22/forum/backend/authentication"
	"github.com/Githaiga22/forum/backend/handlers"
	"github.com/Githaiga22/forum/backend/middleware"
)

func InitRoutes() *http.ServeMux {
	r := http.NewServeMux()

	fs := http.FileServer(http.Dir("./frontend"))
	r.Handle("/frontend/", http.StripPrefix("/frontend/", fs))

	uploadFs := http.FileServer(http.Dir("./uploads"))
	r.Handle("/uploads/", http.StripPrefix("/uploads/", uploadFs))

	// App routes
	r.HandleFunc("/home", middleware.Authenticate(handlers.IndexHandler))
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/sign-in", handlers.LoginHandler)
	r.HandleFunc("/sign-up", handlers.SignupHandler)
	r.HandleFunc("/upload", middleware.Authenticate(handlers.CreatePost))
	r.HandleFunc("/logout", middleware.Authenticate(handlers.LogoutHandler))
	r.HandleFunc("/comments", middleware.Authenticate(handlers.CommentHandler))
	r.HandleFunc("/reaction", middleware.Authenticate(handlers.ReactionHandler))
	r.HandleFunc("/likes", middleware.Authenticate(handlers.ReactionHandler))
	r.HandleFunc("/dilikes", middleware.Authenticate(handlers.ReactionHandler))
	r.HandleFunc("/filter", handlers.FilterPosts)

	r.HandleFunc("/validate", handlers.ValidateInputHandler)

	r.HandleFunc("/auth/google", authentication.GoogleSignUp)
	r.HandleFunc("/auth/google/callback", authentication.GoogleCallback)
	r.HandleFunc("/auth/google/signin", authentication.GoogleSignIn)
	r.HandleFunc("/auth/google/signin/callback", authentication.GoogleSignInCallback)

	r.HandleFunc("/auth/github", authentication.GitHubSignUp)
	r.HandleFunc("/auth/github/signin", authentication.GitHubSignIn)
	r.HandleFunc("/auth/github/callback", authentication.GitHubCallback)
	return r
}
