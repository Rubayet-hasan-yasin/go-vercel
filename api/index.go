package handler

import (
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>Go Vercel Static Endpoint</title>
    <style>
      :root {
        color-scheme: light;
      }
      body {
        margin: 0;
        font-family: "Georgia", "Times New Roman", serif;
        background: radial-gradient(circle at 20% 20%, #f6efe7, #f1f5f8 55%, #e9eef2 100%);
        color: #1b1b1b;
      }
      main {
        min-height: 100vh;
        display: grid;
        place-items: center;
        padding: 48px 24px;
      }
      article {
        max-width: 760px;
        background: #fff;
        border: 1px solid #dedede;
        border-radius: 20px;
        box-shadow: 0 24px 60px rgba(0, 0, 0, 0.08);
        padding: 48px 40px;
      }
      h1 {
        font-size: clamp(2.2rem, 3.2vw, 3.4rem);
        letter-spacing: -0.02em;
        margin: 0 0 16px;
      }
      p {
        font-size: 1.1rem;
        line-height: 1.7;
        margin: 0 0 24px;
      }
      .pill-row {
        display: flex;
        flex-wrap: wrap;
        gap: 12px;
      }
      .pill {
        border: 1px solid #1b1b1b;
        border-radius: 999px;
        padding: 8px 16px;
        font-size: 0.95rem;
        background: #f8f6f1;
      }
    </style>
  </head>
  <body>
    <main>
      <article>
        <h1>Static HTML from Go on Vercel</h1>
        <p>
          This endpoint returns a fully static HTML document with inline styles.
          You can replace the content here without touching any client-side code.
        </p>
        <div class="pill-row">
          <span class="pill">Go HTTP</span>
          <span class="pill">Vercel</span>
          <span class="pill">No JS Required</span>
        </div>
      </article>
    </main>
  </body>
</html>
`))
}
