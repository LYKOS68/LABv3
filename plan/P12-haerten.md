# P12 · Härten & Mobile

**Status:** ausgeplant · **Geschätzte Prompts:** 12–20

## Zulassung
_Ausgang von P11, unverändert übernommen_

- [ ] `/code-review` ist gelaufen.
- [ ] Jedes Finding ist gefixt oder begründet verworfen.
- [ ] (du) Verworfene Findings bestätigt.
- [ ] Commit „P11 Review-Fixes“.
- [ ] (du) Freigabe: P11 fertig.

## Ziel
Robust gegen Unsinn, Spam und kleine Bildschirme.

## Schritte (konkrete Prompts)
1. „validate.go: Hilfsfunktionen für Pflichtfeld, Länge, Enum, E-Mail; alle Handler nutzen sie, Fehler am Feld.“
2. „Honeypot-Feld in Registrieren, Anzeige, Gruppe, Post; ausgefüllt → still verwerfen.“
3. „Rate-Limit: max. 5 Anzeigen/Posts pro Minute und Nutzer → 429 mit freundlicher Seite.“
4. `/run`: alle Seiten bei 375 px, Liste aller Probleme zeigen.
5. „Mobile-Fixes aus der Liste, alles in einem Durchgang.“
6. Design-Runde: dein gesammeltes Feedback in einem Prompt.

## Skills
go-defensive · web-design-engineer · `/run`

## Hilfe (machst du selbst)
- [ ] (du) Gesammeltes Mobile-Feedback in einem Prompt geben.

## Ausgang
_Wird 1:1 die Zulassung von P13._

- [ ] Alle Formulare validieren serverseitig und zeigen den Fehler am Feld.
- [ ] Honeypot und Post-Limit sind aktiv.
- [ ] Keine Seite scrollt bei 375 px horizontal.
- [ ] (du) Mobile-Design abgenommen.
- [ ] Commit „P12 Härten“.
- [ ] (du) Freigabe: P12 fertig.
