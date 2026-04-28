# Documents for ${{values.app_name}}

This is a **Go** Snake-game app, scaffolded by Backstage.

The whole binary is ~7 MB on a `scratch` image. No runtime, no interpreter, just one statically-linked Go executable that embeds its own HTML/JS via `//go:embed`.

# How to access the app?

You can play the game at: `${{values.app_name}}-${{values.app_env}}.gabrieleweka.dev`

# API endpoints

| Endpoint | Purpose |
|---|---|
| `GET /` | Snake game (HTML/CSS/JS, embedded into the binary) |
| `GET /api/v1/info` | JSON: `time`, `hostname`, `env`, `app_name`, `language`, `deployed_on` |
| `GET /api/v1/healthz` | `{"status":"up"}` — used by k8s liveness/readiness probes |

# Game controls

- Arrow keys or **WASD** on desktop
- **Swipe** on mobile/touch
- Best score saves to browser localStorage

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
| **Docker image** | `eweka/${{values.app_name}}:<commit-sha>` (public on Docker Hub, ~7 MB on scratch) |

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
