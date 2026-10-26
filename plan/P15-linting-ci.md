# P15 · Linting & CI

**Status:** ausgeplant · **Geschätzte Prompts:** 3–5

## Zulassung
_Ausgang von P14, unverändert übernommen_

- [ ] Die Seite läuft unter der Domain mit HTTPS.
- [ ] systemd startet den Dienst nach einem Absturz neu.
- [ ] Das tägliche Backup läuft, ein Restore in eine leere DB ist erfolgreich getestet.
- [ ] (du) Eine Test-Fehler-Mail ist angekommen.
- [ ] Commit „P14 Deployment“.
- [ ] (du) Freigabe: P14 fertig.

## Ziel
Jeder Push wird automatisch geprüft.

## Schritte (konkrete Prompts)
1. „.golangci.yml mit sinnvollem Standard-Set, Findings fixen.“
2. „.github/workflows/ci.yml: build, test mit Postgres-Service, lint bei jedem Push und PR.“
3. „CI einmal grün laufen lassen, Link zeigen.“

## Skills
golang-linting · golang-continuous-integration

## Hilfe (machst du selbst)
- [ ] (du) GitHub-Repo anlegen und pushen.

## Ausgang
_Wird 1:1 die Zulassung von P16._

- [ ] `golangci-lint` läuft ohne Fehler.
- [ ] Die CI läuft bei jedem Push grün durch.
- [ ] Commit „P15 CI“.
- [ ] (du) Freigabe: P15 fertig.
