Mit „Mail“ meinst du **Main**, oder? Ich gehe davon aus. Dann ist das der richtige Ansatz: Wir starten **flach in einem einzigen `package main`** und lagern ein Paket erst aus, wenn ein konkreter Grund es erzwingt. Ordner auf Vorrat legen wir nicht an.

## Wann ein eigenes Paket nötig wird

Ein Paket lagern wir erst aus, wenn mindestens einer dieser Fälle eintritt:
1. **Zweiter Nutzer:** Ein zweites Programm (`cmd/...`) braucht denselben Code.
2. **Eigener Lebenszyklus:** Etwas läuft im Hintergrund, zum Beispiel ein Mail-Worker.
3. **Sicherheitsgrenze:** Code wie Zahlungen oder Passwörter soll vom Rest abgeschottet sein.
4. **Unübersichtlich:** Eine Datei hat mehr als etwa 500 Zeilen, oder `package main` hat mehr als etwa 15 Dateien.

---

# Zyklus 1: Main, alles flach (Woche 1–7)

| It. | Inhalt | Was dazukommt |
|---|---|---|
| 1 Fundament | Server, Routen, Postgres-Verbindung, Grundlayout, Brand-CSS | `main.go` · `db.go` · `schema.sql` · `templates/base.html` · `static/` |
| 2 Homepage | Hero, Manifest, Vorschau auf Anzeigen und Gruppen, Impressum, Datenschutz, 404 | `home.go` · `templates/home.html` … |
| 3 Accounts | Registrieren, Login, Sessions in der Datenbank, CSRF-Schutz, Profil | `auth.go` · `profile.go` · Tabellen in `schema.sql` |
| 4 Anzeigen | Biete/Suche, 5 Kategorien, Preis, Status, nur der Ersteller darf bearbeiten, Filter | `ads.go` · `templates/ads_*.html` |
| 5 Gruppen | Erstellen, beitreten, Rollen, Pinnwand | `groups.go` · `templates/groups_*.html` |
| 6 Härten & Live | Validierung, Spam-Schutz, Tests, Backups, Deployment | `*_test.go` · `deploy/` (kein Go-Code) |

**Erstes Auslagern, am Ende von Iteration 6:** Sobald echte Daten in der Datenbank liegen, reicht `schema.sql` nicht mehr, weil man bestehende Tabellen damit nicht sicher ändern kann. Dann kommen `migrations/` und `migrate.go` dazu. Das ist noch kein Paket, nur ein Ordner.

**Stand nach Main:**
```
LABv3/
├── go.mod
├── main.go  db.go  migrate.go  auth.go  profile.go
├── home.go  ads.go  groups.go  validate.go  + Tests
├── migrations/
├── templates/
├── static/
└── deploy/
```
Das sind etwa 10 Go-Dateien und rund 3.000 Zeilen, die flach noch gut lesbar sind.

# Zyklus 2: Beta (Woche 8–11), weiter flach
Einladungscodes, Onboarding, ein Feedback-Button, ein Besucherzähler ohne Cookies und ein Mini-Admin-Bereich.
`+ invite.go · feedback.go · stats.go · admin.go`
**Kein Grund zum Auslagern.**

# Zyklus 3: Interaktion (Woche 12–18), **die große Aufteilung**
Direktnachrichten, Portfolio mit Audio-Uploads, Mail-Benachrichtigungen und Volltextsuche.
**Hier greifen zwei Auslöser:** Der Mail-Versand läuft im Hintergrund (eigener Lebenszyklus), und `package main` hätte über 15 Dateien. Deshalb wird jetzt einmal sauber aufgeteilt:
```
cmd/lab/main.go          ← nur noch Start und Verdrahtung
internal/db/             ← zuerst, weil alle Pakete es brauchen
internal/auth/           ← danach, weil jedes Feature es nutzt
internal/mail/           ← der Worker, das neue Paket
internal/ads/  groups/  messages/  media/  search/ ...
```

# Zyklus 4: Vertrauen & Wachstum (Woche 19–24)
Bewertungen (nur nach einem echten Austausch), Moderation, SEO und Auto-Posting.
**Auslöser:** `cmd/poster` ist ein zweites Programm, das Anzeigen liest. Das funktioniert nur, weil `internal/ads` jetzt ein eigenes Paket ist.
`+ cmd/poster/ · internal/reviews/ · internal/moderation/`

# Zyklus 5: Lyrics-Editor (Woche 25–32)
Ab hier bekommt jedes neue Feature direkt **ein eigenes Paket und einen eigenen Branch**, weil die Struktur jetzt steht.
`+ internal/lyrics/ · static/js/editor.js` (das erste nennenswerte JavaScript)

# Zyklus 6: Monetarisierung (ab Monat 8–9)
**Auslöser:** die Sicherheitsgrenze. Zahlungen werden komplett isoliert.
`+ internal/payments/` (Stripe und Webhooks) · `internal/orders/`

---

Bis zum Ende von Zyklus 2 schreiben wir also nur Inhalt und keine Architektur. Die einzige echte Umbau-Aktion kommt in Zyklus 3, und dann wissen wir genau, wo die Grenzen zwischen den Paketen liegen.

Soll ich mit Iteration 1 anfangen, also `git init`, `go.mod` und `main.go`?

---

Einmal ganz am Anfang tippst du `/golang-how-to configure`. Damit landen die Go-Skills in der `CLAUDE.md` und laden ab dann automatisch, sobald ich Go-Code schreibe. Alles mit `/` davor tippst du selbst am Ende einer Iteration ein. Die übrigen Skills ziehe ich selbst, du kannst sie aber auch mit „nimm go-testing“ erzwingen.

```
Zyklus 1 Main · 6 Iterationen · ~7 Wochen
├── It1 Fundament   → /init + /golang-how-to configure (jede Session kennt den Plan)
│                     golang-project-layout (flach in main festlegen)
│                     golang-popular-libraries (pgx, argon2 einmal sauber wählen)
├── It2 Homepage    → web-design-engineer (Brand.md in echtes HTML/CSS umsetzen)
├── It3 Accounts    → go-error-handling, go-defensive (Login darf nichts verraten)
│                     /security-review am Ende (Passwörter, Sessions, CSRF)
├── It4 Anzeigen    → go-data-structures, go-testing (Rechte per Test absichern)
├── It5 Gruppen     → go-testing, /code-review (ab hier lohnt der Blick von außen)
└── It6 Live        → golang-linting, golang-continuous-integration (Prüfung bei jedem Push)
                      /security-review + /run (alles einmal real durchklicken)

Zyklus 2 Beta · 3 Iterationen · ~4 Wochen
├── It1 Einladung + Onboarding → go-testing
├── It2 Feedback + Statistik   → dataviz (Zahlen im Admin lesbar machen)
└── It3 Fixes                  → /simplify (Wildwuchs aus der Main-Phase aufräumen)

Zyklus 3 Interaktion · 5 Iterationen · ~7 Wochen
├── It0 Umbau in Pakete → golang-project-layout, go-packages, go-naming (sauber schneiden)
├── It1 Nachrichten     → go-concurrency (Worker, Ungelesen-Zähler)
├── It2 Portfolio       → go-defensive + /security-review (Uploads sind das Einfallstor)
├── It3 Mail            → go-concurrency, go-context (Worker sauber stoppen), go-logging
└── It4 Suche           → go-performance (Volltextsuche bleibt schnell)

Zyklus 4 Vertrauen & Wachstum · 4 Iterationen · ~6 Wochen
├── It1 Bewertungen → go-testing (nur nach echtem Austausch, sonst Fake-Bewertungen)
├── It2 Moderation  → /security-review (Admin-Rechte)
├── It3 SEO         → website-bauen (dessen SEO- und GEO-Phase)
└── It4 Auto-Post   → golang-dependency-management (API-Clients prüfen)

Zyklus 5 Lyrics-Editor · 4 Iterationen · ~8 Wochen
├── It1 Editor + Autosave → web-design-engineer (Editor-Oberfläche)
├── It2 Versionen         → go-data-structures
├── It3 Teilen in Gruppen → /code-review
└── It4 Reimhilfe         → go-performance

Zyklus 6 Monetarisierung · 4 Iterationen · ~6 Wochen
├── It1 Recht        → website-bauen (dessen Prüfung von AGB und Cookie-Einwilligung)
├── It2 Stripe       → golang-pkg-go-dev (offizielles SDK und CVEs prüfen)
├── It3 Bestellungen → go-error-handling, go-testing (Geldflüsse ohne Lücke)
└── It4 Abschluss    → /security-review + /code-review
```

`rick-compact-c` und `rick-rewrite` brauchst du hier nicht, die sind für C bzw. fürs Portieren in eine andere Sprache. Die Carta-Skills haben mit diesem Projekt nichts zu tun. `website-bauen` stellt viele Fragen. Wenn du es nutzt, sag ihm am besten gleich, dass es aus `Brand.md` und `draft.md` ableiten soll.
