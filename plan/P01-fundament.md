# P01 · Fundament

**Status:** ausgeplant · **Geschätzte Prompts:** 8–12

## Zulassung
_Ausgangslage_

- [ ] Go ≥ 1.22 ist installiert (`go version`).
- [ ] PostgreSQL läuft lokal, eine leere Datenbank `lab` mit eigenem User existiert.
- [ ] `Brand.md`, `draft.md` und die Logo-/SVG-Assets liegen in `LABv3/`.

## Ziel
Ein laufender Go-Server mit Postgres, Grundlayout und Brand-CSS – alles flach in package main.

## Schritte (konkrete Prompts)
1. „Lies draft.md und Brand.md. Iteration 1: git init, .gitignore, go.mod (Modul `lab`), main.go mit net/http-Server auf :8080 und GET /health. Baue und starte.“
2. „db.go: pgx-Pool aus DATABASE_URL. /health pingt die Datenbank und antwortet 503, wenn sie nicht erreichbar ist.“
3. „schema.sql wird beim Start ausgeführt (nur CREATE … IF NOT EXISTS). Vorerst leer bis auf einen Kommentar.“
4. „templates/base.html mit nav und footer, Render-Helper, static/ ausliefern. Logos und SVGs nach static/img.“
5. „static/css/brand.css: alle Tokens aus Brand.md §4–8, Fonts, Grundtypografie, Primary- und Secondary-Button.“
6. `/golang-how-to configure` und danach `/init`
7. Fix-Runde: alles, was beim ersten Öffnen im Browser auffällt, gesammelt in einem Prompt.

## Skills
golang-project-layout (flache Struktur bewusst festlegen) · golang-popular-libraries (pgx einmal sauber wählen) · `/golang-how-to configure` + `/init` (jede künftige Session kennt Plan und Regeln).

## Hilfe (machst du selbst)
- [ ] (du) Postgres installieren, Datenbank und User anlegen, `DATABASE_URL` setzen.
- [ ] (du) Seite einmal im Browser öffnen und sagen, ob Fonts und Farben stimmen.

## Ausgang
_Wird 1:1 die Zulassung von P02._

- [ ] `go build` läuft fehlerfrei.
- [ ] `GET /health` antwortet 200 mit erreichbarer DB und 503 ohne.
- [ ] Die Startseite rendert `base.html` mit `brand.css`, Fonts und Logo.
- [ ] `schema.sql` wird beim Start ausgeführt.
- [ ] `CLAUDE.md` existiert, erster Commit ist gemacht.
- [ ] (du) Freigabe: P01 fertig.
