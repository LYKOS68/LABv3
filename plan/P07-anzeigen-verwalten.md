# P07 · Anzeigen · Verwalten

**Status:** ausgeplant · **Geschätzte Prompts:** 8–12

## Zulassung
_Ausgang von P06, unverändert übernommen_

- [ ] Eingeloggte Nutzer können eine Anzeige mit Typ, Kategorie und Preis erstellen.
- [ ] Liste und Detailseite sind ohne Login sichtbar.
- [ ] Commit „P06 Anzeigen Basis“.
- [ ] (du) Freigabe: P06 fertig.

## Ziel
Bearbeiten, löschen, Status, Filter.

## Schritte (konkrete Prompts)
1. „GET/POST /anzeigen/{id}/bearbeiten und POST /anzeigen/{id}/loeschen, nur der Ersteller, sonst 403.“
2. „Status aktiv/pausiert/archiviert umschaltbar; öffentliche Liste zeigt nur aktive.“
3. „/me/anzeigen: eigene Anzeigen mit allen Status.“
4. „Filter ?kategorie=&typ= als Pills über der Liste, Seitenaufteilung à 20 mit ?seite=.“
5. Fix-Runde nach dem Test mit zweitem Account.

## Skills
go-defensive

## Hilfe (machst du selbst)
- [ ] (du) Mit dem zweiten Account versuchen, eine fremde Anzeige zu ändern.

## Ausgang
_Wird 1:1 die Zulassung von P08._

- [ ] Nur der Ersteller kann bearbeiten, löschen und den Status ändern, alle anderen bekommen 403.
- [ ] Der Filter nach Kategorie und Typ funktioniert, die Liste ist in Seiten à 20 aufgeteilt.
- [ ] (du) Mit dem zweiten Account ist eine fremde Anzeige nicht änderbar.
- [ ] Commit „P07 Anzeigen verwalten“.
- [ ] (du) Freigabe: P07 fertig.
