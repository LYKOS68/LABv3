# P08 · Anzeigen-Tests · MESSPUNKT 2

**Status:** ausgeplant · **Geschätzte Prompts:** 4–6

## Zulassung
_Ausgang von P07, unverändert übernommen_

- [ ] Nur der Ersteller kann bearbeiten, löschen und den Status ändern, alle anderen bekommen 403.
- [ ] Der Filter nach Kategorie und Typ funktioniert, die Liste ist in Seiten à 20 aufgeteilt.
- [ ] (du) Mit dem zweiten Account ist eine fremde Anzeige nicht änderbar.
- [ ] Commit „P07 Anzeigen verwalten“.
- [ ] (du) Freigabe: P07 fertig.

## Ziel
Anzeigen-Rechte absichern, Homepage mit echten Daten.

## Schritte (konkrete Prompts)
1. „Tests: fremde Anzeige bearbeiten/löschen/Status → 403, ohne Login erstellen → Redirect, pausierte Anzeige nicht in öffentlicher Liste.“
2. „Tests: Filter liefert nur passende Kategorie/Typ, Seite 2 beginnt bei Anzeige 21.“
3. „Homepage-Vorschau: 3 neueste aktive Anzeigen aus der DB, Platzhalter wenn keine da.“
4. `/run`: kompletter Anzeigen-Weg bei 375 px mit Screenshots.

## Skills
go-testing · `/run`

## Hilfe (machst du selbst)
- [ ] (du) Den ganzen Anzeigen-Weg einmal auf dem eigenen Handy durchklicken.

## Ausgang
_Wird 1:1 die Zulassung von P09._

- [ ] Alle Anzeigen-Tests laufen grün.
- [ ] Die Homepage zeigt die 3 neuesten aktiven Anzeigen.
- [ ] (du) Der Durchlauf auf dem Handy ist ohne Fehler bestanden.
- [ ] Commit „P08 Anzeigen-Tests“.
- [ ] (du) Freigabe: P08 fertig.
