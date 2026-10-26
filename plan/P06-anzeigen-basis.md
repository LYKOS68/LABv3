# P06 · Anzeigen · Basis

**Status:** ausgeplant · **Geschätzte Prompts:** 8–12

## Zulassung
_Ausgang von P05, unverändert übernommen_

- [ ] Alle Auth-Tests laufen grün (`go test ./...`).
- [ ] `/security-review` hat keine offenen High- oder Medium-Findings.
- [ ] Commit „P05 Auth-Tests“.
- [ ] (du) Freigabe: P05 fertig.

## Ziel
Anzeigen erstellen und ansehen.

## Schritte (konkrete Prompts)
1. „schema.sql: ads (id, user_id, typ biete/suche, kategorie mix/master/beats/ghostwriting/feature, titel max. 80, text max. 2000, preis_cent nullable = auf Anfrage, status, created_at).“
2. „ads.go: GET /anzeigen (Liste, neueste zuerst), GET /anzeigen/{id} (Detail), GET/POST /anzeigen/neu (nur eingeloggt).“
3. „Templates ads_list/ads_show/ads_form: Karten laut Brand.md §7, Platzhalterbild default-anzeige-svg, Kategorie als Pill.“
4. „Detail: Kontakt-Button führt zum Profil des Erstellers, Preis als ‚€ 120‘ oder ‚auf Anfrage‘.“
5. Fix-Runde nach dem Einstellen der Beispielanzeigen.

## Skills
go-data-structures · web-design-engineer (Karten laut Brand.md §7)

## Hilfe (machst du selbst)
- [ ] (du) 3–5 realistische Beispielanzeigen selbst einstellen.

## Ausgang
_Wird 1:1 die Zulassung von P07._

- [ ] Eingeloggte Nutzer können eine Anzeige mit Typ, Kategorie und Preis erstellen.
- [ ] Liste und Detailseite sind ohne Login sichtbar.
- [ ] Commit „P06 Anzeigen Basis“.
- [ ] (du) Freigabe: P06 fertig.
