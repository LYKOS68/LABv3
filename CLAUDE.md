# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Projekt

RAP-LAB: Kleinanzeigen- und Community-Plattform für die Musikbranche (Anzeigen für Beats, Mix/Master, Ghostwriting usw., Profile, Gruppen, Pinnwand). Sprache von UI, Code-Kommentaren, Plan und Kommunikation: **Deutsch**.

## Arbeiten nach Plan

Die Arbeit folgt den Plan-Dateien in `.data/plan/` (P01–P16, Zyklus 1 „Main“). Die Langfrist-Roadmap steht in `.data/draft.md`, das Design-System in `.data/Brand.md` (Farben §4, Typografie §5, UI-System §7, Tailwind-Token §10, Landing-Copy §11).

- Pro Schritt gilt: Der **Ausgang** eines P ist wörtlich die **Zulassung** des nächsten. Vor Arbeit an einem P dessen Zulassung prüfen.
- Ein P = ein Commit (Nachricht z. B. „P02 Homepage“).
- Häkchen ohne **(du)** setzt Claude nur mit Nachweis (Build-Ausgabe, Testergebnis, HTTP-Antwort). Häkchen mit **(du)** setzt nur der Nutzer.
- Der nächste P beginnt erst nach ausdrücklicher Freigabe durch den Nutzer.
- Neue Ideen kommen in `draft.md`, nicht in den laufenden Schritt.
- Die Spalte „Skills“ in jeder P-Datei nennt die Skills, die für den Schritt geladen werden sollen.

## Architektur-Regel: flach in `package main`

Alles bleibt in **einem `package main`** und wächst auf `main.go` plus Geschwister-Dateien (`db.go`, `auth.go`, `ads.go`, `groups.go` …). Keine Ordner oder Pakete auf Vorrat. Ein eigenes Paket wird erst ausgelagert, wenn einer dieser Auslöser aus `draft.md` greift:

1. Ein zweites Programm (`cmd/...`) braucht denselben Code.
2. Etwas hat einen eigenen Lebenszyklus (Hintergrund-Worker).
3. Sicherheitsgrenze (z. B. Zahlungen).
4. Eine Datei > ~500 Zeilen oder `package main` > ~15 Dateien.

Die große Aufteilung in `cmd/` + `internal/` ist erst für Zyklus 3 geplant. `migrations/` + `migrate.go` kommen in P13 (Ordner, kein Paket). Die DB-Struktur steht bis dahin in `schema.sql` (nur `CREATE … IF NOT EXISTS`, wird beim Start ausgeführt).

Geplanter Stack laut Plan: `net/http` (Go ≥ 1.22 Routing-Patterns), Postgres über pgx-Pool aus `DATABASE_URL`, Sessions in der DB, argon2id, CSRF in allen Formularen. Konfiguration über Umgebungsvariablen; die Variablen je Phase sind in `.env` dokumentiert (`DATABASE_URL`, `PORT`, `APP_ENV`, `SESSION_SECRET`, `ARGON_*` …). `.env` ist gitignored.

## Code-Stand und bewusste Abweichungen vom Plan

Go-Modul `raplab` im Repo-Root, Server auf Port 8080 (`PORT` überschreibt).

**Entschieden (weicht von P01 ab):** Statt `templates/*.html` mit `html/template` nutzen wir **templ + Tailwind**. Wo der Plan `templates/xyz.html` nennt, entsteht stattdessen `views/xyz.templ`. `views/` ist das einzige Unterpaket und ersetzt den Ordner `templates/`; die Go-Logik bleibt flach in `package main`.

- `main.go`: `net/http`-Server, `static/` per `go:embed` ins Binary eingebettet.
- `views/*.templ`: Seiten mit [templ](https://templ.guide). `views/*_templ.go` ist **generiert**, nie von Hand bearbeiten.
- `static/input.css` → `static/app.css`: Tailwind **3.4.17** (fest gepinnt; kein `npm audit fix --force`, das zieht v4 und bricht den Build). Brand-Farben in `tailwind.config.js`; Brand.md §10 liefert die vollständigen Token.

## Befehle (im Repo-Root)

```sh
npm install                 # einmalig: Tailwind
go tool templ generate      # nach Änderungen an views/*.templ
npm run css                 # nach Änderungen an Styles / neuen Klassen
go run .                    # http://localhost:8080
go vet ./... && go build ./...
go test ./...               # einzelner Test: go test -run TestName ./...
```

Live-Entwicklung: `npm run css:watch` und `go tool templ generate --watch --proxy=http://localhost:8080 --cmd="go run ."` in zwei Terminals.

Tests mit Datenbank (ab P05) laufen gegen eine eigene DB über `TEST_DATABASE_URL`; jede Test-Funktion startet mit leeren Tabellen.

## Go development

Before any Go coding, review, debugging, troubleshooting, or setup task, load the `samber/cc-skills-golang@golang-how-to` skill first — it routes to whichever other Go skills the task needs.

## Required Go skills

The following Go skills from `samber/cc-skills-golang` MUST always be applied when working on this project. Load them at the start of every Go-related task, regardless of whether the user explicitly mentions them.

- `samber/cc-skills-golang@golang-code-style`
- `samber/cc-skills-golang@golang-data-structures`
- `samber/cc-skills-golang@golang-database`
- `samber/cc-skills-golang@golang-design-patterns`
- `samber/cc-skills-golang@golang-documentation`
- `samber/cc-skills-golang@golang-error-handling`
- `samber/cc-skills-golang@golang-modernize`
- `samber/cc-skills-golang@golang-naming`
- `samber/cc-skills-golang@golang-safety`
- `samber/cc-skills-golang@golang-security`
- `samber/cc-skills-golang@golang-testing`
- `samber/cc-skills-golang@golang-troubleshooting`
