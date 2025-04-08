package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hallo Wolrd")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}

	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		return
	}

	rewriteCookie := http.Cookie{
		Name:    "flash",
		MaxAge:  -1,
		Expires: time.Unix(1, 0),
	}

	http.SetCookie(w, &rewriteCookie)
	val, _ := base64.URLEncoding.DecodeString(c.Value)
	fmt.Fprintln(w, string(val))
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "impostor_engineer",
		Value:    "am i ?",
		HttpOnly: true,
	}

	// w.Header().Set("Set-Cookie", c1.String())
	// w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("first_cookie")
	if err != nil {
		return
	}

	cookies := r.Cookies()
	see := fmt.Sprintf("All Cookies: %s, Cookie Value: %s", cookies, c)

	fmt.Fprintln(w, see)
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	http.Handle("/cookie", http.HandlerFunc(setCookie))
	http.HandleFunc("/get_cookie", getCookie)
	http.HandleFunc("/flash_message", setMessage)
	http.HandleFunc("/show_message", showMessage)

	server.ListenAndServe()
}
