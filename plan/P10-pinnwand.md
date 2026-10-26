# P10 · Pinnwand

**Status:** ausgeplant · **Geschätzte Prompts:** 8–12

## Zulassung
_Ausgang von P09, unverändert übernommen_

- [ ] Gruppen lassen sich erstellen, auflisten und ansehen.
- [ ] Beitreten und Verlassen funktionieren, der Ersteller ist Owner.
- [ ] Commit „P09 Gruppen Basis“.
- [ ] (du) Freigabe: P09 fertig.

## Ziel
In Gruppen posten, mit klaren Rechten.

## Schritte (konkrete Prompts)
1. „schema.sql: posts (id, group_id, user_id, text max. 1000, created_at).“
2. „POST /gruppen/{slug}/posts nur für Mitglieder, sonst 403. Pinnwand neueste zuerst auf der Gruppenseite.“
3. „POST /posts/{id}/loeschen: Owner der Gruppe alle, Mitglied nur eigene, sonst 403.“
4. „Tests: Nicht-Mitglied postet → 403, Mitglied löscht fremden Post → 403, Owner löscht fremden Post → ok.“
5. „Homepage-Vorschau: 3 Gruppen mit den meisten Posts der letzten 14 Tage.“

## Skills
go-testing

## Hilfe (machst du selbst)
- [ ] (du) Als Nicht-Mitglied versuchen zu posten.

## Ausgang
_Wird 1:1 die Zulassung von P11._

- [ ] Nur Mitglieder können posten, Nicht-Mitglieder bekommen 403.
- [ ] Der Owner kann alle Beiträge löschen, ein Mitglied nur seine eigenen.
- [ ] Alle Tests laufen grün, die Homepage zeigt aktive Gruppen.
- [ ] Commit „P10 Pinnwand“.
- [ ] (du) Freigabe: P10 fertig.
