# Documents for ${{values.app_name}}

This is a **Go** web app, scaffolded by Backstage.

It serves a Bootstrap-based landing page (Tooplate "Barista" theme) plus two JSON API endpoints. The whole binary is ~17 MB on a `scratch` image. No runtime, no interpreter — just one statically-linked Go executable that embeds its own HTML/CSS/JS/images via `//go:embed`.

# How to access the app

Open: `https://${{values.app_name}}-${{values.app_env}}.gabrieleweka.dev`

# API endpoints

| Endpoint | Purpose |
|---|---|
| `GET /` | Bootstrap landing page, rebranded with this app's name |
| `GET /api/v1/info` | JSON: `time`, `hostname`, `env`, `app_name`, `language`, `deployed_on` |
| `GET /api/v1/healthz` | `{"status":"up"}` — used by k8s liveness/readiness probes |

# Tweaking the page

Static assets live in `src/web/`:

- `index.html` — the landing page
- `reservation.html` — secondary form page
- `css/`, `js/`, `images/`, `videos/`, `fonts/`

Edit any of these and push — CI rebuilds the image with the embedded assets. No separate frontend build step.

---

# Links for THIS app

| What | URL |
|---|---|
| **Live app**         | https://${{values.app_name}}-${{values.app_env}}.gabrieleweka.dev |
| **Source code**      | https://github.com/octal-engine/${{values.app_name}} |
| **CI / CD pipeline** | https://github.com/octal-engine/${{values.app_name}}/actions |
| **ArgoCD app**       | https://argocd.gabrieleweka.dev/applications/${{values.app_name}} |
| **Grafana dashboard** (per-app metrics) | https://grafana.gabrieleweka.dev/d/app-${{values.app_name}}-${{values.app_env}} |
| **SonarQube project** (code quality) | https://sonar.gabrieleweka.dev/dashboard?id=${{values.app_name}} |
| **Docker image** | `eweka/${{values.app_name}}:<commit-sha>` (public on Docker Hub, ~17 MB on scratch) |

---

# Platform-wide URLs & Credentials

| Service | URL | Auth |
|---|---|---|
| **Backstage** (this portal) | https://backstage.gabrieleweka.dev | Sign in with GitHub or Google |
| **ArgoCD** | https://argocd.gabrieleweka.dev | `viewer` / `viewer123` (read-only) |
| **Grafana** | http://grafana.gabrieleweka.dev | No login needed (anonymous Viewer) |
| **Prometheus** | http://prometheus.gabrieleweka.dev | No auth |
| **SonarQube** | https://sonar.gabrieleweka.dev | `viewer` / `viewer123` (read-only) |

For admin access, ping **@Eweka01**.
