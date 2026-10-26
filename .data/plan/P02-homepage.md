# P02 · Homepage

**Status:** ausgeplant · **Geschätzte Prompts:** 12–20

## Zulassung
_Ausgang von P01, unverändert übernommen_

- [ ] `go build` läuft fehlerfrei.
- [ ] `GET /health` antwortet 200 mit erreichbarer DB und 503 ohne.
- [ ] Die Startseite rendert `base.html` mit `brand.css`, Fonts und Logo.
- [ ] `schema.sql` wird beim Start ausgeführt.
- [ ] `CLAUDE.md` existiert, erster Commit ist gemacht.
- [ ] (du) Freigabe: P01 fertig.

## Ziel
Eine knappe Startseite nach Brand.md plus die Pflichtseiten.

## Schritte (konkrete Prompts)
1. „Homepage laut Brand.md §11: Hero mit Emblem (Blend-Mode), Name, ein Satz Positionierung, ein CTA. Mobile zuerst.“
2. „Manifest-Block, danach Vorschau-Blöcke für Anzeigen und Gruppen mit default-anzeige-svg als Platzhalter.“
3. „Impressum, Datenschutz und 404-Seite im selben Layout. Footer verlinkt Impressum und Datenschutz.“
4. „Navigation: transparent über dem Hero, ab Scrollen mit Blur. Mobiles Vollbild-Menü.“
5. `/run`: Startseite bei 375 px und 1280 px prüfen, Screenshots zeigen.
6. Design-Runde: alle Wünsche gesammelt in einem Prompt, nicht einzeln.

## Skills
web-design-engineer (Brand.md in echtes HTML/CSS) · `/run` (selbst sehen statt beschreiben).

## Hilfe (machst du selbst)
- [ ] (du) Echte Impressumsdaten liefern (Name, Anschrift, Kontakt-Mail) – das kann ich nicht erfinden.
- [ ] (du) Auf dem eigenen Handy anschauen und Design-Feedback gesammelt geben.

## Ausgang
_Wird 1:1 die Zulassung von P03._

- [ ] Die Startseite zeigt Hero, Manifest, Vorschau Anzeigen und Vorschau Gruppen.
- [ ] Impressum, Datenschutz und 404 sind erreichbar und im Footer verlinkt.
- [ ] Bei 375 px Breite gibt es kein horizontales Scrollen.
- [ ] (du) Design auf dem eigenen Handy abgenommen.
- [ ] Commit „P02 Homepage“.
- [ ] (du) Freigabe: P02 fertig.
