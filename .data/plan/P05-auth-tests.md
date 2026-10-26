# P05 · Auth-Tests · MESSPUNKT 1

**Status:** ausgeplant · **Geschätzte Prompts:** 4–6

## Zulassung
_Ausgang von P04, unverändert übernommen_

- [ ] `/u/{name}` ist öffentlich sichtbar.
- [ ] Bearbeiten ist nur für das eigene Profil möglich.
- [ ] Alle `/me`-Routen leiten ohne Login auf `/login` um.
- [ ] (du) Mit zwei Accounts die Profile gegenseitig geprüft.
- [ ] Commit „P04 Profil“.
- [ ] (du) Freigabe: P04 fertig.

## Ziel
Beweisen, dass Fremde und Nicht-Eingeloggte draußen bleiben.

## Schritte (konkrete Prompts)
1. „Test-Setup: eigene DB lab_test über TEST_DATABASE_URL, jede Test-Funktion startet mit leeren Tabellen.“
2. „Tests: falsches Passwort, unbekannte E-Mail (gleiche Meldung), ohne Login auf /me → Redirect, abgelaufene Session, fremde Session-ID.“
3. „Tests: Formular ohne CSRF → 403, fremdes Profil bearbeiten → 403, Login-Rate-Limit → 429.“
4. `/security-review`
5. „fix die Findings, verworfene mit Begründung in diese Datei.“

## Skills
go-testing · `/security-review`

## Hilfe (machst du selbst)
- [ ] (du) Test-Datenbank `lab_test` anlegen.

## Ausgang
_Wird 1:1 die Zulassung von P06._

- [ ] Alle Auth-Tests laufen grün (`go test ./...`).
- [ ] `/security-review` hat keine offenen High- oder Medium-Findings.
- [ ] Commit „P05 Auth-Tests“.
- [ ] (du) Freigabe: P05 fertig.
