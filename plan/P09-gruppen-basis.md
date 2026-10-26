# P09 · Gruppen · Basis

**Status:** ausgeplant · **Geschätzte Prompts:** 8–12

## Zulassung
_Ausgang von P08, unverändert übernommen_

- [ ] Alle Anzeigen-Tests laufen grün.
- [ ] Die Homepage zeigt die 3 neuesten aktiven Anzeigen.
- [ ] (du) Der Durchlauf auf dem Handy ist ohne Fehler bestanden.
- [ ] Commit „P08 Anzeigen-Tests“.
- [ ] (du) Freigabe: P08 fertig.

## Ziel
Gruppen erstellen, finden, beitreten.

## Schritte (konkrete Prompts)
1. „schema.sql: groups (id, name unique, slug unique, beschreibung, owner_id, created_at) und memberships (group_id, user_id, rolle owner/member, joined_at).“
2. „groups.go: GET /gruppen, GET /gruppen/{slug}, GET/POST /gruppen/neu (nur eingeloggt), Ersteller wird automatisch Owner.“
3. „POST /gruppen/{slug}/beitreten und /verlassen; Owner kann nicht verlassen, solange er Owner ist.“
4. „Templates groups_list/groups_show/groups_form, Mitgliederliste mit Avataren.“
5. Fix-Runde nach dem Beitritt mit dem zweiten Account.

## Skills
go-data-structures

## Hilfe (machst du selbst)
- [ ] (du) Eine echte Gruppe anlegen, mit dem zweiten Account beitreten.

## Ausgang
_Wird 1:1 die Zulassung von P10._

- [ ] Gruppen lassen sich erstellen, auflisten und ansehen.
- [ ] Beitreten und Verlassen funktionieren, der Ersteller ist Owner.
- [ ] Commit „P09 Gruppen Basis“.
- [ ] (du) Freigabe: P09 fertig.
